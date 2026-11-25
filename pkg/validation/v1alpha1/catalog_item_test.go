/*
Copyright (C) 2022-2026 Traefik Labs

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

package v1alpha1_test

import (
	"fmt"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestAPICatalogItem_Validation(t *testing.T) {
	t.Parallel()

	parentRefs := make([]string, 101)
	for i := range parentRefs {
		parentRefs[i] = fmt.Sprintf("    - name: my-portal-%d", i)
	}
	tooManyParentRefs := "\n" + strings.Join(parentRefs, "\n")

	tests := []validationTestCase{
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: "my-catalog-items"
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "valid: minimal with everyone true",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true`),
		},
		{
			desc: "valid: minimal with groups",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  groups:
    - my-group`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  apiPlan:
    name: my-plan
  groups:
    - my-group
  apis:
    - name: my-api
  apiSelector:
    matchLabels:
      key: value
  operationFilter:
    include:
      - my-filter
  parentRefs:
    - name: my-portal
      namespace: portal`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: .non-dns-compliant-catalog-items
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-catalog-items", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: ""
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: catalog-items-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "catalog-items-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "duplicated APIs",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  apiPlan:
    name: my-plan
  apis:
    - name: my-api
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: field.OmitValueType{}, Detail: "duplicated apis"}},
		},
		{
			desc: "duplicated API: implicit default",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  apiPlan:
    name: my-plan
  apis:
    - name: my-api
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: field.OmitValueType{}, Detail: "duplicated apis"}},
		},
		{
			desc: "invalid API selector",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  apiPlan:
    name: my-plan
  apiSelector:
    matchExpressions:
      - key: value`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.apiSelector.matchExpressions[0].operator", BadValue: ""}},
		},
		{
			desc: "everyone and groups both set",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  apiPlan:
    name: my-plan
  everyone: true
  groups:
    - my-group`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: field.OmitValueType{}, Detail: "groups and everyone are mutually exclusive"}},
		},
		{
			desc: "everyone is false and no groups",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: false`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: field.OmitValueType{}, Detail: "groups is required when everyone is false"}},
		},
		{
			desc: "missing apiPlan name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  apiPlan: {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.apiPlan.name", BadValue: ""}},
		},
		{
			desc: "valid: parentRefs without namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  parentRefs:
    - name: my-portal`),
		},
		{
			desc: "valid: same parentRef name in different namespaces",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  parentRefs:
    - name: my-portal
      namespace: portal
    - name: my-portal
      namespace: another-portal`),
		},
		{
			desc: "empty parentRefs",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  parentRefs: []`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.parentRefs", BadValue: int64(0), Detail: "spec.parentRefs in body should have at least 1 items"}},
		},
		{
			desc: "missing parentRef name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  parentRefs:
    - namespace: portal`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.parentRefs[0].name", BadValue: ""}},
		},
		{
			desc: "parentRef name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  parentRefs:
    - name: "` + strings.Repeat("x", 254) + `"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.parentRefs[0].name", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "parentRef namespace is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  parentRefs:
    - name: my-portal
      namespace: "` + strings.Repeat("x", 64) + `"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.parentRefs[0].namespace", BadValue: "<value omitted>", Detail: "may not be more than 63 bytes"}},
		},
		{
			desc: "duplicated parentRefs without namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  parentRefs:
    - name: my-portal
    - name: my-portal`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeDuplicate, Field: "spec.parentRefs[1]", BadValue: map[string]any{"name": "my-portal"}}},
		},
		{
			desc: "duplicated parentRefs with namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  parentRefs:
    - name: my-portal
      namespace: portal
    - name: my-portal
      namespace: portal`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeDuplicate, Field: "spec.parentRefs[1]", BadValue: map[string]any{"name": "my-portal", "namespace": "portal"}}},
		},
		{
			desc: "too many parentRefs",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  everyone: true
  parentRefs:` + tooManyParentRefs),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooMany, Field: "spec.parentRefs", BadValue: 101, Detail: "must have at most 100 items"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
