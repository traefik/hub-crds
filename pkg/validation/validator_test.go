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

package validation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/hub-crds/pkg/validation"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	validator := validation.NewValidator()

	err := validator.Register(&apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{Name: "MyResource"},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Names: apiextensions.CustomResourceDefinitionNames{
				Kind: "MyResource",
			},
			Group: "test",
			Scope: apiextensions.NamespaceScoped,
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name: "v1alpha1",
					Schema: &apiextensions.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
							Type: "object",
							Properties: map[string]apiextensions.JSONSchemaProps{
								"spec": {
									Type: "object",
									Properties: map[string]apiextensions.JSONSchemaProps{
										"foo": {Type: "string"},
										"bar": {Type: "string"},
										"baz": {
											Type: "string",
											XValidations: apiextensions.ValidationRules{
												{
													Rule:    "self.startsWith('baz')",
													Message: "must start with 'baz'",
												},
											},
										},
									},
									Required: []string{"foo"},
								},
							},
							Required: []string{"spec"},
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)

	tests := []struct {
		desc     string
		data     map[string]interface{}
		wantErrs field.ErrorList
	}{
		{
			desc: "valid manifest",
			data: map[string]interface{}{
				"kind":       "MyResource",
				"apiVersion": "test/v1alpha1",
				"metadata": map[string]interface{}{
					"name":      "test",
					"namespace": "default",
				},
				"spec": map[string]interface{}{
					"foo": "foo",
					"bar": "bar",
					"baz": "bazinga",
				},
			},
		},
		{
			desc: "missing required field",
			data: map[string]interface{}{
				"kind":       "MyResource",
				"apiVersion": "test/v1alpha1",
				"metadata": map[string]interface{}{
					"name":      "test",
					"namespace": "default",
				},
				"spec": map[string]interface{}{
					"bar": "bar",
					"baz": "bazinga",
				},
			},
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.foo", BadValue: ""}},
		},
		{
			desc: "invalid CEL rule",
			data: map[string]interface{}{
				"kind":       "MyResource",
				"apiVersion": "test/v1alpha1",
				"metadata": map[string]interface{}{
					"name":      "test",
					"namespace": "default",
				},
				"spec": map[string]interface{}{
					"foo": "foo",
					"bar": "bar",
					"baz": "foobar",
				},
			},
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.baz", BadValue: "string", Detail: "must start with 'baz'"}},
		},
		{
			desc: "invalid field type",
			data: map[string]interface{}{
				"kind":       "MyResource",
				"apiVersion": "test/v1alpha1",
				"metadata": map[string]interface{}{
					"name":      "test",
					"namespace": "default",
				},
				"spec": map[string]interface{}{
					"foo": "foo",
					"bar": 2,
					"baz": "bazinga",
				},
			},
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTypeInvalid, Field: "spec.bar", BadValue: "integer", Detail: "spec.bar in body must be of type string: \"integer\""}},
		},
		{
			desc: "metadata validation",
			data: map[string]interface{}{
				"kind":       "MyResource",
				"apiVersion": "test/v1alpha1",
				"metadata": map[string]interface{}{
					"namespace": "default",
				},
				"spec": map[string]interface{}{
					"foo": "foo",
					"bar": "bar",
					"baz": "bazinga",
				},
			},
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "unknown resource",
			data: map[string]interface{}{
				"kind": "Something",
			},
		},
		{
			desc: "not a K8S resource",
			data: map[string]interface{}{
				"hello": "world",
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			gotErrs := validator.Validate(&unstructured.Unstructured{Object: test.data})
			assert.Equal(t, test.wantErrs, gotErrs)
		})
	}
}
