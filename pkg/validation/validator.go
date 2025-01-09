/*
Copyright (C) 2022-2025 Traefik Labs

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
	"context"
	"fmt"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema/cel"
	apiservervalidation "k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"k8s.io/apimachinery/pkg/api/meta"
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apiservercel "k8s.io/apiserver/pkg/apis/cel"
)

// Validator validates the Kubernetes resources against their OpenAPI specification.
// It runs spec, metadata and CEL validations.
type Validator struct {
	structuralSchemas map[string]*schema.Structural
	namespaced        map[string]bool
	schemaValidators  map[string]apiservervalidation.SchemaValidator
	celValidators     map[string]*cel.Validator
}

// NewValidator creates a new Validator.
func NewValidator() *Validator {
	return &Validator{
		structuralSchemas: make(map[string]*schema.Structural),
		namespaced:        make(map[string]bool),
		schemaValidators:  make(map[string]apiservervalidation.SchemaValidator),
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

		schemaValidator, _, err := apiservervalidation.NewSchemaValidator(validationSchema.OpenAPIV3Schema)
		if err != nil {
			return fmt.Errorf("creating schema validator: %w", err)
		}

		celValidator := cel.NewValidator(structuralSchema, true, apiservercel.PerCallLimit)

		key := runtimeschema.GroupVersionKind{
			Group:   crd.Spec.Group,
			Version: version.Name,
			Kind:    crd.Spec.Names.Kind,
		}.String()

		v.schemaValidators[key] = schemaValidator
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
		fieldErrs = append(fieldErrs, apiservervalidation.ValidateCustomResource(nil, unstructuredContent, validator)...)
	}

	// Validate CEL rules.
	if validator, ok := v.celValidators[key]; ok {
		celErrs, _ := validator.Validate(context.Background(), nil, structuralSchema, unstructuredContent, nil, apiservercel.RuntimeCELCostBudget)
		fieldErrs = append(fieldErrs, celErrs...)
	}

	return fieldErrs
}
