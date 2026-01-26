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
			desc: "valid: minimal API key authentication - will use default values",
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
			desc: "valid: full API key authentication",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  apiKey:
    keySource:
      header: Authorization
      headerAuthScheme: Auth-Scheme
      query: key`),
		},
		{
			desc: "invalid: API key with empty keySource",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  apiKey:
    keySource: {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apiKey.keySource", BadValue: int64(0), Detail: "spec.apiKey.keySource in body should have at least 1 properties"}},
		},
		{
			desc: "invalid: API key with headerAuthScheme on non-Authorization header",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  apiKey:
    keySource:
      header: X-API-Key
      headerAuthScheme: Bearer`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apiKey.keySource", BadValue: "object", Detail: "headerAuthScheme can only be used when header is 'Authorization'"}},
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
    appIdClaim: "client_id"
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
    appIdClaim: "client_id"
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
    appIdClaim: "client_id"
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
    appIdClaim: "client_id"
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
    appIdClaim: "client_id"
    tokenNameClaim: "token_name"
    forwardHeaders:
      X-User-ID: "sub"
      X-User-Email: "email"
    signingSecretName: "jwt-secret"`),
		},
		{
			desc: "valid: JWT with single trusted issuer with issuer specified",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    trustedIssuers:
      - jwksUrl: "https://tenant-a.example.com/jwks.json"
        issuer: "https://tenant-a.example.com/"`),
		},
		{
			desc: "valid: JWT with multiple trusted issuers with different issuers",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    trustedIssuers:
      - jwksUrl: "https://tenant-a.example.com/jwks.json"
        issuer: "https://tenant-a.example.com/"
      - jwksUrl: "https://tenant-b.example.com/jwks.json"
        issuer: "https://tenant-b.example.com/"`),
		},
		{
			desc: "valid: JWT with multiple trusted issuers with mix of specific and fallback",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    trustedIssuers:
      - jwksUrl: "https://tenant-a.example.com/jwks.json"
        issuer: "https://tenant-a.example.com/"
      - jwksUrl: "https://tenant-b.example.com/jwks.json"
        issuer: "https://tenant-b.example.com/"
      - jwksUrl: "https://fallback.example.com/jwks.json"`),
		},
		{
			desc: "valid: JWT with trusted issuers fallback only (no issuer specified)",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    trustedIssuers:
      - jwksUrl: "https://fallback.example.com/jwks.json"`),
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
			desc: "valid: minimal LDAP configuration",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-ldap-auth
  namespace: default
spec:
  isDefault: true
  ldap:
    url: "ldap://ldap.example.com:389"
    baseDn: "dc=example,dc=com"`),
		},
		{
			desc: "valid: LDAP with LDAPS",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-ldap-auth
  namespace: default
spec:
  isDefault: true
  ldap:
    url: "ldaps://ldap.example.com:636"
    baseDn: "dc=example,dc=com"`),
		},
		{
			desc: "valid: full LDAP configuration",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-ldap-auth
  namespace: default
spec:
  isDefault: true
  ldap:
    url: "ldaps://ldap.example.com:636"
    startTls: true
    insecureSkipVerify: false
    certificateAuthority: |
      -----BEGIN CERTIFICATE-----
      MIIBCzCBsqADAgECAhBaooOsws+BLdvtfqQ1ggx5MAoGCCqGSM49BAMCMBIxEDAO
      BgNVBAoTB0V4YW1wbGUwHhcNMjQwMTAxMDAwMDAwWhcNMjUwMTAxMDAwMDAwWjAS
      MRAwDgYDVQQKEwdFeGFtcGxlMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKNho
      zEli5D+VsLgKJcgT0rp+MnYGJ4PjN8qgXfQx1F5JhtVBVnH8qWmza2XwJvVZAgUg
      WijH8vDvBJU8su1w16MdMBswDgYDVR0PAQH/BAQDAgWgMAkGA1UdEwQCMAAwCgYI
      KoZIzj0EAwIDSAAwRQIgcu4/UZKPaUPCAB2jjqKbW8XqBp8fv1F8D5FO5hL1DqwC
      IQDB3g0Lhx4QbM3Kw6bk0gvvVVLkb/TXe2Nvl4dH8gOCEw==
      -----END CERTIFICATE-----
    bindDn: "cn=admin,dc=example,dc=com"
    bindPasswordSecretName: "ldap-bind-password"
    baseDn: "dc=example,dc=com"
    attribute: "uid"
    searchFilter: "(&(objectClass=inetOrgPerson)(uid=%s))"`),
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
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: "object", Detail: "exactly one authentication method must be specified"}},
		},
		{
			desc: "multiple authentication methods specified",
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
    appIdClaim: "client_id"
    signingSecretName: "jwt-secret"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: "object", Detail: "exactly one authentication method must be specified"}},
		},
		{
			desc: "JWT missing required appIdClaim",
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
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.jwt.appIdClaim", BadValue: ""}},
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
    appIdClaim: "client_id"
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
    appIdClaim: "client_id"
    jwksUrl: "not-a-url"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt.jwksUrl", BadValue: "string", Detail: "must be a valid HTTPS URL"}},
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
    appIdClaim: "client_id"
    stripAuthorizationHeader: true`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt", BadValue: "object", Detail: "exactly one of signingSecretName, publicKey, jwksFile, jwksUrl, or trustedIssuers must be specified"}},
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
    appIdClaim: "client_id"
    signingSecretName: "jwt-secret"
    publicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt", BadValue: "object", Detail: "exactly one of signingSecretName, publicKey, jwksFile, jwksUrl, or trustedIssuers must be specified"}},
		},
		{
			desc: "JWT trustedIssuers with invalid JWKS URL",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    trustedIssuers:
      - jwksUrl: "not-a-valid-url"
        issuer: "https://tenant-a.example.com/"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt.trustedIssuers[0].jwksUrl", BadValue: "string", Detail: "must be a valid HTTPS URL"}},
		},
		{
			desc: "JWT trustedIssuers combined with jwksUrl, mutual exclusivity",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    jwksUrl: "https://example.com/.well-known/jwks.json"
    trustedIssuers:
      - jwksUrl: "https://tenant-a.example.com/jwks.json"
        issuer: "https://tenant-a.example.com/"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt", BadValue: "object", Detail: "exactly one of signingSecretName, publicKey, jwksFile, jwksUrl, or trustedIssuers must be specified"}},
		},
		{
			desc: "JWT trustedIssuers combined with publicKey, mutual exclusivity",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    publicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----"
    trustedIssuers:
      - jwksUrl: "https://tenant-a.example.com/jwks.json"
        issuer: "https://tenant-a.example.com/"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt", BadValue: "object", Detail: "exactly one of signingSecretName, publicKey, jwksFile, jwksUrl, or trustedIssuers must be specified"}},
		},
		{
			desc: "JWT trustedIssuers with empty array",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    trustedIssuers: []`),
			wantErrs: field.ErrorList{
				{Type: field.ErrorTypeInvalid, Field: "spec.jwt.trustedIssuers", BadValue: int64(0), Detail: "spec.jwt.trustedIssuers in body should have at least 1 items"},
				{Type: field.ErrorTypeInvalid, Field: "spec.jwt", BadValue: "object", Detail: "trustedIssuers must not be empty when specified"},
			},
		},
		{
			desc: "JWT trustedIssuers with multiple fallback entries",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    trustedIssuers:
      - jwksUrl: "https://fallback-1.example.com/jwks.json"
      - jwksUrl: "https://fallback-2.example.com/jwks.json"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt", BadValue: "object", Detail: "only one entry in trustedIssuers may omit the issuer field"}},
		},
		{
			desc: "JWT trustedIssuers with empty jwksUrl",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    trustedIssuers:
      - jwksUrl: ""
        issuer: "https://tenant-a.example.com/"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt.trustedIssuers[0].jwksUrl", BadValue: "string", Detail: "must be a valid HTTPS URL"}},
		},
		{
			desc: "JWT trustedIssuers with HTTP URL (not HTTPS)",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    trustedIssuers:
      - jwksUrl: "http://tenant-a.example.com/jwks.json"
        issuer: "https://tenant-a.example.com/"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt.trustedIssuers[0].jwksUrl", BadValue: "string", Detail: "must be a valid HTTPS URL"}},
		},
		{
			desc: "JWT jwksUrl with HTTP URL (not HTTPS)",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-auth
  namespace: default
spec:
  isDefault: true
  jwt:
    appIdClaim: "client_id"
    jwksUrl: "http://example.com/.well-known/jwks.json"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.jwt.jwksUrl", BadValue: "string", Detail: "must be a valid HTTPS URL"}},
		},
		{
			desc: "LDAP missing required URL",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-ldap-auth
  namespace: default
spec:
  isDefault: true
  ldap:
    baseDn: "dc=example,dc=com"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.ldap.url", BadValue: ""}},
		},
		{
			desc: "LDAP missing required baseDn",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-ldap-auth
  namespace: default
spec:
  isDefault: true
  ldap:
    url: "ldap://ldap.example.com:389"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.ldap.baseDn", BadValue: ""}},
		},
		{
			desc: "LDAP invalid URL format",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-ldap-auth
  namespace: default
spec:
  isDefault: true
  ldap:
    url: "https://ldap.example.com"
    baseDn: "dc=example,dc=com"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.ldap.url", BadValue: "string", Detail: "must be a valid LDAP URL"}},
		},
		{
			desc: "LDAP bindPasswordSecretName too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIAuth
metadata:
  name: my-ldap-auth
  namespace: default
spec:
  isDefault: true
  ldap:
    url: "ldap://ldap.example.com:389"
    baseDn: "dc=example,dc=com"
    bindPasswordSecretName: "` + tooLongSecretName + `"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.ldap.bindPasswordSecretName", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
