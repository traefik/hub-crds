/*
Copyright (C) 2022-2023 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package validation

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/traefik/hub-crds/hub/v1alpha1/crd"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema/cel"
	apiservervalidation "k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"k8s.io/apimachinery/pkg/api/meta"
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/kube-openapi/pkg/validation/validate"
)

// Validator validates the Kubernetes resources against their OpenAPI specification.
// It runs spec, metadata and CEL validations.
type Validator struct {
	structuralSchemas map[string]*schema.Structural
	namespaced        map[string]bool
	schemaValidators  map[string]validate.SchemaValidator
	celValidators     map[string]*cel.Validator
}

// NewValidator creates a new Validator.
func NewValidator() *Validator {
	return &Validator{
		structuralSchemas: make(map[string]*schema.Structural),
		namespaced:        make(map[string]bool),
		schemaValidators:  make(map[string]validate.SchemaValidator),
		celValidators:     make(map[string]*cel.Validator),
	}
}

// Register registers a CRD to be validated later on.
func (v *Validator) Register(crd *apiextensions.CustomResourceDefinition) error {
	for _, version := range crd.Spec.Versions {
		validationSchema, err := apiextensions.GetSchemaForVersion(crd, version.Name)
		if err != nil {
			return fmt.Errorf("obtaining validation schema version %q: %w", version.Name, err)
		}

		structuralSchema, err := schema.NewStructural(validationSchema.OpenAPIV3Schema)
		if err != nil {
			return fmt.Errorf("building structural schema: %w", err)
		}

		schemaValidator, _, err := apiservervalidation.NewSchemaValidator(validationSchema)
		if err != nil {
			return fmt.Errorf("creating schema validator: %w", err)
		}

		celValidator := cel.NewValidator(structuralSchema, true, cel.PerCallLimit)

		key := runtimeschema.GroupVersionKind{
			Group:   crd.Spec.Group,
			Version: version.Name,
			Kind:    crd.Spec.Names.Kind,
		}.String()

		v.schemaValidators[key] = *schemaValidator
		v.celValidators[key] = celValidator
		v.structuralSchemas[key] = structuralSchema
		v.namespaced[key] = crd.Spec.Scope == apiextensions.NamespaceScoped
	}

	return nil
}

// Validate validates the given object and report potential issues.
// Unknown objects are skipped without returning any error.
func (v *Validator) Validate(obj *unstructured.Unstructured) field.ErrorList {
	key := obj.GetObjectKind().GroupVersionKind().String()

	structuralSchema, ok := v.structuralSchemas[key]
	if !ok {
		// Skip unknown resource.
		return nil
	}

	unstructuredContent := obj.UnstructuredContent()

	namespaced := v.namespaced[key]

	var fieldErrs field.ErrorList

	// Validate object metadata.
	accessor, err := meta.Accessor(obj)
	if err != nil {
		fieldErrs = append(fieldErrs, field.Invalid(field.NewPath("metadata"), nil, err.Error()))
	} else {
		// TODO: replace NameIsDNSLabel by NameIsDNSSubdomain once the backend will have relaxed this constrain.
		fieldErrs = append(fieldErrs, apivalidation.ValidateObjectMetaAccessor(accessor, namespaced, apivalidation.NameIsDNSLabel, field.NewPath("metadata"))...)
	}

	// Validate object schema.
	if validator, ok := v.schemaValidators[key]; ok {
		fieldErrs = append(fieldErrs, apiservervalidation.ValidateCustomResource(nil, unstructuredContent, &validator)...)
	}

	// Validate CEL rules.
	if validator, ok := v.celValidators[key]; ok {
		celErrs, _ := validator.Validate(context.Background(), nil, structuralSchema, unstructuredContent, nil, cel.RuntimeCELCostBudget)
		fieldErrs = append(fieldErrs, celErrs...)
	}

	return fieldErrs
}

// BuildHubValidator returns a new Validator.
func BuildHubValidator() (*Validator, error) {
	decoder, err := NewCRDDecoder()
	if err != nil {
		return nil, fmt.Errorf("creating CRD decoder: %w", err)
	}

	manifests, err := loadManifests(crd.CRDs)
	if err != nil {
		return nil, fmt.Errorf("loading CRD documents: %w", err)
	}

	validator := NewValidator()

	for _, m := range manifests {
		customResource, err := decoder.Decode(m.Data)
		if err != nil {
			return nil, fmt.Errorf("decoding manifest %s: %w", m.Path, err)
		}

		if err = validator.Register(customResource); err != nil {
			return nil, fmt.Errorf("registering CRD %q: %w", customResource.Name, err)
		}
	}

	return validator, nil
}

type manifest struct {
	Path string
	Data []byte
}

func loadManifests(filesystem fs.FS) ([]manifest, error) {
	var manifests []manifest

	err := fs.WalkDir(filesystem, ".", func(path string, entry fs.DirEntry, fileErr error) error {
		if fileErr != nil {
			return fileErr
		}

		if filepath.Base(path) == ".git" {
			return filepath.SkipDir
		}

		if !entry.Type().IsRegular() || !isYAMLOrJSON(path) {
			return nil
		}

		reader, err := filesystem.Open(path)
		if err != nil {
			return fmt.Errorf("opening file: %w", err)
		}
		defer func() { _ = reader.Close() }()

		r := yaml.NewYAMLReader(bufio.NewReader(reader))

		for {
			data, err := r.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}

				return fmt.Errorf("reading file content: %w", err)
			}

			manifests = append(manifests, manifest{
				Data: data,
				Path: path,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return manifests, nil
}

func isYAMLOrJSON(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))

	return ext == ".yaml" || ext == ".yml" || ext == ".json"
}
