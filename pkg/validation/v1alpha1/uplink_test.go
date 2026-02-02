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

func TestUplink_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: my-uplink
  namespace: default`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: my-uplink
  namespace: default
spec:
  exposeName: my-custom-name
  entryPoints:
    - multi-cluster
  weight: 10`),
		},
		{
			desc: "valid: with exposeName",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: my-uplink
  namespace: default
spec:
  exposeName: default-my-uplink`),
		},
		{
			desc: "valid: multiple entrypoints",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: my-uplink
  namespace: default
spec:
  entryPoints:
    - multi-cluster
    - another-entrypoint`),
		},
		{
			desc: "valid: weight zero",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: my-uplink
  namespace: default
spec:
  weight: 0`),
		},
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: my-uplink`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: .non-dns-compliant-uplink
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-uplink", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: ""
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: uplink-with-a-way-tooooooooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "uplink-with-a-way-tooooooooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "negative weight",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Uplink
metadata:
  name: my-uplink
  namespace: default
spec:
  weight: -1`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.weight", BadValue: "integer", Detail: "must be a positive number"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
