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

func TestAPIPortal_Validation(t *testing.T) {
	t.Parallel()

	tooLongAuthName := strings.Repeat("x", 254)
	tooLongNamespaceName := strings.Repeat("x", 64)

	namespaceNames := make([]string, 101)
	for i := range namespaceNames {
		namespaceNames[i] = fmt.Sprintf("namespace-%d", i)
	}
	tooManyNamespaceNames := strings.Join(namespaceNames, ", ")

	tests := []validationTestCase{
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  title: title
  description: description
  trustedUrls: ["https://example.com"]
  ui:
    logoUrl: https://example.com/logo.png
  allowedAPICatalogItems:
    namespaces:
      from: Selector
      selector:
        matchLabels:
          portal-visibility: my-portal
      names:
        - bank
  `),
		},
		{
			desc: "valid: with auth reference",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  auth:
    name: my-auth
  `),
		},
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
spec:
  trustedUrls: ["https://example.com"]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: .non-dns-compliant-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-portal", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: ""
  namespace: default
spec:
  trustedUrls: ["https://example.com"]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: portal-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default
spec:
  trustedUrls: ["https://example.com"]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "portal-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "missing trustedUrls",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec: {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.trustedUrls", BadValue: ""}},
		},
		{
			desc: "empty trustedUrls",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: []`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.trustedUrls", BadValue: int64(0), Detail: "spec.trustedUrls in body should have at least 1 items"}},
		},
		{
			desc: "too many trustedUrls",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com", https://another.example.com]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooMany, Field: "spec.trustedUrls", BadValue: 2, Detail: "must have at most 1 item"}},
		},
		{
			desc: "auth name too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  auth:
    name: "` + tooLongAuthName + `"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.auth.name", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "missing allowedAPICatalogItems namespaces",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems: {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.allowedAPICatalogItems.namespaces", BadValue: ""}},
		},
		{
			desc: "valid: empty allowedAPICatalogItems namespaces",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces: {}`),
		},
		{
			desc: "valid: allowedAPICatalogItems from Same",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: Same`),
		},
		{
			desc: "valid: allowedAPICatalogItems from All",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: All`),
		},
		{
			desc: "valid: allowedAPICatalogItems from Selector with an empty selector",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: Selector
      selector: {}`),
		},
		{
			desc: "valid: allowedAPICatalogItems from Selector with names only",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: Selector
      names:
        - bank
        - insurance`),
		},
		{
			desc: "valid: allowedAPICatalogItems from Selector without criteria",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: Selector`),
		},
		{
			desc: "unsupported allowedAPICatalogItems from",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: Everything`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeNotSupported, Field: "spec.allowedAPICatalogItems.namespaces.from", BadValue: "Everything", Detail: `supported values: "All", "Selector", "Same"`}},
		},
		{
			desc: "invalid allowedAPICatalogItems selector",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: Selector
      selector:
        matchExpressions:
          - key: value`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.allowedAPICatalogItems.namespaces.selector.matchExpressions[0].operator", BadValue: ""}},
		},
		{
			desc: "allowedAPICatalogItems namespace name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: Selector
      names:
        - "` + tooLongNamespaceName + `"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.allowedAPICatalogItems.namespaces.names[0]", BadValue: "<value omitted>", Detail: "may not be more than 63 bytes"}},
		},
		{
			desc: "duplicated allowedAPICatalogItems namespace names",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: Selector
      names:
        - bank
        - bank`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeDuplicate, Field: "spec.allowedAPICatalogItems.namespaces.names[1]", BadValue: "bank"}},
		},
		{
			desc: "too many allowedAPICatalogItems namespace names",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]
  allowedAPICatalogItems:
    namespaces:
      from: Selector
      names: [` + tooManyNamespaceNames + `]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooMany, Field: "spec.allowedAPICatalogItems.namespaces.names", BadValue: 101, Detail: "must have at most 100 items"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
