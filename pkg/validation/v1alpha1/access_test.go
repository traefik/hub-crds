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

package v1alpha1_test

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestAPIAccess_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  groups:
    - my-group
  apis:
    - name: my-api
      namespace: my-ns
  apiSelector:
    labelSelector:
      key: value
  apiCollections:
    - name: my-api-collection
  apiCollectionSelector:
    labelSelector:
      key: value
  operationFilter:
    include:
      - my-filter`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: .non-dns-compliant-access`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-access", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: ""`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: access-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "access-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "duplicated APIs",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  apis:
    - name: my-api
      namespace: my-ns
    - name: my-api
      namespace: my-ns`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: "array", Detail: "duplicated apis"}},
		},
		{
			desc: "duplicated API: implicit default",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  apis:
    - name: my-api
      namespace: default
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: "array", Detail: "duplicated apis"}},
		},
		{
			desc: "valid: apis with same name but different namespaces",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  apis:
    - name: my-api
      namespace: my-ns-1
    - name: my-api
      namespace: my-ns-2`),
		},
		{
			desc: "valid: apis with same name but only one with namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  apis:
    - name: my-api
      namespace: my-ns
    - name: my-api`),
		},
		{
			desc: "duplicated collections",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  apiCollections:
    - name: my-collection
    - name: my-collection`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apiCollections", BadValue: "array", Detail: "duplicated collections"}},
		},
		{
			desc: "valid: different collections",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  apiCollections:
    - name: my-collection-1
    - name: my-collection-2`),
		},
		{
			desc: "invalid API selector",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  apiSelector:
    matchExpressions:
      - key: value`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.apiSelector.matchExpressions[0].operator", BadValue: ""}},
		},
		{
			desc: "invalid collection selector",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  apiCollectionSelector:
    matchExpressions:
      - key: value`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.apiCollectionSelector.matchExpressions[0].operator", BadValue: ""}},
		},
		{
			desc: "anyGroups and groups both set",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAccess
metadata:
  name: my-access
spec:
  anyGroups: true
  groups:
    - my-group`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: "object", Detail: "groups and anyGroups are mutually exclusive"}},
		},
	}

	checkValidationTestCases(t, tests)
}
