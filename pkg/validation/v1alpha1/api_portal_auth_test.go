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
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestAPIPortalAuth_Validation(t *testing.T) {
	t.Parallel()

	tooLongSecretName := strings.Repeat("x", 254)

	tests := []validationTestCase{
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: "my-auth"
spec:
  oidc:
    issuerUrl: "https://auth.example.com"
    secretName: "oidc-secret"
    claims:
      groups: "groups"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "valid: minimal configuration",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    issuerUrl: "https://auth.example.com"
    secretName: "oidc-secret"
    claims:
      groups: "groups"`),
		},
		{
			desc: "valid: full configuration",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    issuerUrl: "https://auth.example.com"
    secretName: "oidc-secret"
    scopes:
      - "openid"
      - "profile"
      - "email"
    claims:
      userId: "sub"
      firstname: "given_name"
      lastname: "family_name"
      email: "email"
      groups: "groups"
      company: "organization"
    syncedAttributes:
      - "userId"
      - "firstname"
      - "company"`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: .non-dns-compliant-auth
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-auth", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing required issuerUrl",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    secretName: "oidc-secret"
    claims:
      groups: "groups"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.oidc.issuerUrl", BadValue: ""}},
		},
		{
			desc: "missing required secretName",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    issuerUrl: "https://auth.example.com"
    claims:
      groups: "groups"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.oidc.secretName", BadValue: ""}},
		},
		{
			desc: "missing required groups claim",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    issuerUrl: "https://auth.example.com"
    secretName: "oidc-secret"
    claims: {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.oidc.claims.groups", BadValue: ""}},
		},
		{
			desc: "invalid issuer URL format",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    issuerUrl: "not-a-url"
    secretName: "oidc-secret"
    claims:
      groups: "groups"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.oidc.issuerUrl", BadValue: "string", Detail: "must be a valid URL"}},
		},
		{
			desc: "secretName too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    issuerUrl: "https://auth.example.com"
    secretName: "` + tooLongSecretName + `"
    claims:
      groups: "groups"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.oidc.secretName", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "valid: syncedAttributes with configured claims",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    issuerUrl: "https://auth.example.com"
    secretName: "oidc-secret"
    claims:
      userId: "sub"
      firstname: "given_name"
      groups: "groups"
    syncedAttributes:
      - "userId"
      - "firstname"`),
		},
		{
			desc: "valid: syncedAttributes contains valid fields",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    issuerUrl: "https://auth.example.com"
    secretName: "oidc-secret"
    claims:
      groups: "groups"
    syncedAttributes:
      - "userId"
      - "company"`),
		},
		{
			desc: "invalid: syncedAttributes contains unknown field",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortalAuth
metadata:
  name: my-auth
  namespace: default
spec:
  oidc:
    issuerUrl: "https://auth.example.com"
    secretName: "oidc-secret"
    claims:
      groups: "groups"
    syncedAttributes:
      - "unknownField"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.oidc.syncedAttributes", BadValue: "array", Detail: "syncedAttributes must only contain: groups, userId, firstname, lastname, email, company"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
