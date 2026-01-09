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
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestAPIBundle_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: "my-bundle"
spec:
  apis:
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: my-bundle
  namespace: default
spec:
  apis:
    - name: my-api`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: my-bundle
  namespace: default
spec:
  title: My API Bundle Title
  apis:
    - name: my-api
  apiSelector:
    matchLabels:
      key: value`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: .non-dns-compliant-access
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-access", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: ""
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: access-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "access-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "title is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: my-bundle
  namespace: default
spec:
  title: Way Toooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo Long Title`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.title", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "duplicated APIs",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: my-bundle
  namespace: default
spec:
  apis:
    - name: my-api
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: "array", Detail: "duplicated apis"}},
		},
		{
			desc: "duplicated API: implicit default",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: my-bundle
  namespace: default
spec:
  apis:
    - name: my-api
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: "array", Detail: "duplicated apis"}},
		},
		{
			desc: "invalid API selector",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIBundle
metadata:
  name: my-bundle
  namespace: default
spec:
  apiSelector:
    matchExpressions:
      - key: value`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.apiSelector.matchExpressions[0].operator", BadValue: ""}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
