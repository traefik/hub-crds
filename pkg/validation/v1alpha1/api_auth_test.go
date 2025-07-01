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

func TestAPIAuth_Validation(t *testing.T) {
	t.Parallel()

	tooLongSecretName := strings.Repeat("x", 254)

	tests := []validationTestCase{
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: "my-auth"
spec:
  isDefault: true
  apiKey: {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "valid: API key authentication",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  apiKey: {}`),
		},
		{
			desc: "valid: JWT with signing secret",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    signingSecretName: "jwt-secret"`),
		},
		{
			desc: "valid: JWT with public key",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    publicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----"`),
		},
		{
			desc: "valid: JWT with JWKS file",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    jwksFile: '{"keys":[{"kty":"RSA","use":"sig","kid":"abc","n":"...","e":"AQAB"}]}'`),
		},
		{
			desc: "valid: JWT with JWKS URL",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    jwksUrl: "https://example.com/.well-known/jwks.json"`),
		},
		{
			desc: "valid: JWT with all options",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    stripAuthorizationHeader: true
    tokenQueryKey: "token"
    forwardHeaders:
      X-User-ID: "sub"
      X-User-Email: "email"
    signingSecretName: "jwt-secret"`),
		},
		{
			desc: "valid: not default",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: false
  apiKey: {}`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: .non-dns-compliant-auth
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-auth", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing authentication method",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: "object", Detail: "exactly one of apiKey or jwt must be specified"}},
		},
		{
			desc: "both authentication methods specified",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  apiKey: {}
  jwt:
    signingSecretName: "jwt-secret"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: "object", Detail: "exactly one of apiKey or jwt must be specified"}},
		},
		{
			desc: "JWT signingSecretName too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    signingSecretName: "` + tooLongSecretName + `"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.jwt.signingSecretName", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "invalid JWKS URL format",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    jwksUrl: "not-a-url"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt.jwksUrl", BadValue: "string", Detail: "must be a valid URL"}},
		},
		{
			desc: "JWT missing verification method",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    stripAuthorizationHeader: true`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt", BadValue: "object", Detail: "exactly one of signingSecretName, publicKey, jwksFile, or jwksUrl must be specified"}},
		},
		{
			desc: "JWT multiple verification methods",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    signingSecretName: "jwt-secret"
    publicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt", BadValue: "object", Detail: "exactly one of signingSecretName, publicKey, jwksFile, or jwksUrl must be specified"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
