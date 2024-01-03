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

func TestAPIGateway_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIGateway
metadata:
  name: my-gateway`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIGateway
metadata:
  name: my-gateway
spec:
  apiAccesses:
    - my-access
  customDomains:
    - test.example.com`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIGateway
metadata:
  name: .non-dns-compliant-gateway`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-gateway", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIGateway
metadata:
  name: ""`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIGateway
metadata:
  name: gateway-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "gateway-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "empty custom domain",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIGateway
metadata:
  name: my-gateway
spec:
  apiAccesses:
    - my-access
  customDomains:
    - ""`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.customDomains[0]", BadValue: "string", Detail: "custom domain must be a valid domain name"}},
		},
		{
			desc: "invalid custom domain",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIGateway
metadata:
  name: my-gateway
spec:
  apiAccesses:
    - my-access
  customDomains:
    - example..com`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.customDomains[0]", BadValue: "string", Detail: "custom domain must be a valid domain name"}},
		},
		{
			desc: "duplicate custom domains",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIGateway
metadata:
  name: my-gateway
spec:
  apiAccesses:
    - my-access
  customDomains:
    - example.com
    - example.com`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.customDomains", BadValue: "array", Detail: "duplicate domains"}},
		},
	}

	checkValidationTestCases(t, tests)
}
