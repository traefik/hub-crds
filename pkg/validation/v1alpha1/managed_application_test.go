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

package v1alpha1_test

import (
	"fmt"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestManagedApplication_Validation(t *testing.T) {
	t.Parallel()

	tooLongAPIKey := strings.Repeat("x", 4097)

	tests := []validationTestCase{
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: "my-application"
spec:
  appId: "123"
  owner: "456"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: my-application
  namespace: default
spec:
  appId: "123"
  owner: "456"`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: my-application
  namespace: default
spec:
  appId: "123"
  owner: "456"
  notes: blablabla
  apiKeys:
    - secretName: secret
      title: My APIKey
      suspended: true`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: .non-dns-compliant-application
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-application", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: ""
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: application-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "application-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "appId is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: my-application
  namespace: default
spec:
  appId: Way Toooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo Long AppId
  owner: "456"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.appId", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "owner is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: my-application
  namespace: default
spec:
  appId: "123"
  owner: Way Tooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo Long Owner`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.owner", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "apiKey secretName is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: my-application
  namespace: default
spec:
  appId: "123"
  owner: "456"
  apiKeys:
    - secretName: Way Tooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo Long secret`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.apiKeys[0].secretName", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "apiKey value is too long",
			manifest: []byte(fmt.Sprintf(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: my-application
  namespace: default
spec:
  appId: "123"
  owner: "456"
  apiKeys:
    - value: %s`, tooLongAPIKey)),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.apiKeys[0].value", BadValue: "<value omitted>", Detail: "may not be more than 4096 bytes"}},
		},
		{
			desc: "apiKey secretName and value both set",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ManagedApplication
metadata:
  name: my-application
  namespace: default
spec:
  appId: "123"
  owner: "456"
  apiKeys:
    - secretName: secret
      value: value`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apiKeys[0]", BadValue: "object", Detail: "secretName and value are mutually exclusive"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
