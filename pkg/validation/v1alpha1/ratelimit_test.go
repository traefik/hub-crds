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

func TestAPIRateLimit_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: 1`),
		},
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: 1
  period: 2s
  strategy: distributed
  groups:
    - my-group
  apis:
    - name: my-api`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: .non-dns-compliant-ratelimit
  namespace: default
spec:
  limit: 10`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-ratelimit", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: ""
  namespace: default
spec:
  limit: 10`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: ratelimit-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default
spec:
  limit: 10`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "ratelimit-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "everyone and groups both set",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: 10
  everyone: true
  groups:
    - my-group`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: "object", Detail: "groups and everyone are mutually exclusive"}},
		},
		{
			desc: "limit must be a positive integer",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: -10
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.limit", BadValue: "integer", Detail: "must be a positive number"}},
		},
		{
			desc: "period must be less than 1 hour",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: 1
  period: 2h
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.period", BadValue: "string", Detail: "must be between 1s and 1h"}},
		},
		{
			desc: "period must be more than 1 second",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: 1
  period: 0s
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.period", BadValue: "string", Detail: "must be between 1s and 1h"}},
		},
		{
			desc: "strategy must be local or distributed",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: 1
  strategy: yolo
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeNotSupported, Field: "spec.strategy", BadValue: "yolo", Detail: "supported values: \"local\", \"distributed\""}},
		},
		{
			desc: "duplicated APIs",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: 1
  apis:
    - name: my-api
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: "array", Detail: "duplicated apis"}},
		},
		{
			desc: "duplicated API: implicit default",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: 1
  apis:
    - name: my-api
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: "array", Detail: "duplicated apis"}},
		},
		{
			desc: "invalid API selector",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIRateLimit
metadata:
  name: my-ratelimit
  namespace: default
spec:
  limit: 1
  apiSelector:
    matchExpressions:
      - key: value`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.apiSelector.matchExpressions[0].operator", BadValue: ""}},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
