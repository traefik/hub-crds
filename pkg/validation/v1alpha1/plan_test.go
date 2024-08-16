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

func TestAPIPlan_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan
  namespace: default
spec:
  title: my-plan`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan
  namespace: default
spec:
  title: my-plan
  description: |-
    # Full Markdown description
    - with content
  rateLimit:
    limit: 1
    period: 2s
  quota:
    limit: 1
    period: 2s`),
		},
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "missing title",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan
  namespace: default
spec:
  rateLimit:
    limit: 1`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.title", BadValue: ""}},
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: .non-dns-compliant-plan
  namespace: default
spec:
  title: my-plan
  rateLimit:
    limit: 1
  quota:
    limit: 10`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-plan", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: ""
  namespace: default
spec:
  title: my-plan
  rateLimit:
    limit: 1
  quota:
    limit: 10`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: plan-with-a-way-toooooooooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default
spec:
  title: my-plan
  rateLimit:
    limit: 1
  quota:
    limit: 10`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "plan-with-a-way-toooooooooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "quota limit must be a positive integer",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan
  namespace: default
spec:
  title: my-plan
  quota:
    limit: -10
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.quota.limit", BadValue: "integer", Detail: "must be a positive number"}},
		},
		{
			desc: "ratelimit limit must be a positive integer",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan
  namespace: default
spec:
  title: my-plan
  rateLimit:
    limit: -1`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.rateLimit.limit", BadValue: "integer", Detail: "must be a positive number"}},
		},
		{
			desc: "quota period must be less than 9999 hour",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan
  namespace: default
spec:
  title: my-plan
  quota:
    limit: 1
    period: 10000h`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.quota.period", BadValue: "string", Detail: "must be between 1s and 9999h"}},
		},
		{
			desc: "rate limit period must be less than 1 hour",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan
  namespace: default
spec:
  title: my-plan
  rateLimit:
    limit: 1
    period: 2h`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.rateLimit.period", BadValue: "string", Detail: "must be between 1s and 1h"}},
		},
		{
			desc: "quota period must be more than 1 second",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan
  namespace: default
spec:
  title: my-plan
  quota:
    limit: 1
    period: 0s`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.quota.period", BadValue: "string", Detail: "must be between 1s and 9999h"}},
		},
		{
			desc: "ratelimit period must be more than 1 second",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPlan
metadata:
  name: my-plan
  namespace: default
spec:
  title: my-plan
  rateLimit:
    limit: 1
    period: 0s`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.rateLimit.period", BadValue: "string", Detail: "must be between 1s and 1h"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
