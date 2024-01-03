/*
Copyright (C) 2022-2024 Traefik Labs

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

func TestAPICollection_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: my-collection`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: my-collection
  pathPrefix: /collection
  apis:
    - name: my-api`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: .non-dns-compliant-collection`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-collection", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: ""`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: collection-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "collection-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "duplicated APIs",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: my-collection
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
kind: APICollection
metadata:
  name: my-collection
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
kind: APICollection
metadata:
  name: my-collection
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
kind: APICollection
metadata:
  name: my-collection
spec:
  apis:
    - name: my-api
      namespace: my-ns
    - name: my-api`),
		},
		{
			desc: "invalid API selector",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: my-collection
spec:
  apiSelector:
    matchExpressions:
      - key: value`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.apiSelector.matchExpressions[0].operator", BadValue: ""}},
		},
		{
			desc: "path prefix must start with a /",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: my-collection
spec:
  pathPrefix: something
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.pathPrefix", BadValue: "string", Detail: "must start with a '/'"}},
		},
		{
			desc: "path prefix cannot contains ../",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: my-collection
spec:
  pathPrefix: /foo/../bar
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.pathPrefix", BadValue: "string", Detail: "cannot contains '../'"}},
		},
		{
			desc: "path prefix cannot ends with /..",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: my-collection
spec:
  pathPrefix: /foo/..
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.pathPrefix", BadValue: "string", Detail: "cannot contains '../'"}},
		},
		{
			desc: "valid: pathPrefix with segment starting with ..",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: my-collection
spec:
  pathPrefix: /foo/..bar
`),
		},
		{
			desc: "valid: pathPrefix with segment starting with .well-known",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICollection
metadata:
  name: my-collection
spec:
  pathPrefix: /foo/.well-known
`),
		},
	}

	checkValidationTestCases(t, tests)
}
