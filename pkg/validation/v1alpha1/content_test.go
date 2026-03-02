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

func TestContent_Validation(t *testing.T) {
	t.Parallel()

	tooLongName := strings.Repeat("x", 254)

	tests := []validationTestCase{
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: "my-content"
spec:
  title: My Content
  order: 0
  parentRef:
    kind: API
    name: my-api
  content: "# Hello World"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "valid: minimal with content",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: API
    name: my-api
  content: "# Hello World"`),
		},
		{
			desc: "valid: minimal with link",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: API
    name: my-api
  link:
    href: https://example.com`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content Title
  order: 42
  parentRef:
    kind: APIVersion
    name: my-api-version
  content: "# Full Content\n\nThis is the full content."`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: .non-dns-compliant-content
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-content", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: ""
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: content-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "content-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "title is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: my-content
  namespace: default
spec:
  title: Way Toooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo Long Title
  order: 0
  parentRef:
    kind: API
    name: my-api
  content: "# Hello World"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.title", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "parentRef kind is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: "` + tooLongName + `"
    name: my-api
  content: "# Hello World"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.parentRef.kind", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "parentRef name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: API
    name: "` + tooLongName + `"
  content: "# Hello World"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.parentRef.name", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "both content and link specified",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: API
    name: my-api
  content: "# Hello World"
  link:
    href: https://example.com`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: "object", Detail: "exactly one of content or link must be specified"}},
		},
		{
			desc: "neither content nor link specified",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: API
    name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: "object", Detail: "exactly one of content or link must be specified"}},
		},
		{
			desc: "invalid link href",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: Content
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: API
    name: my-api
  link:
    href: not-a-valid-url`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.link.href", BadValue: "string", Detail: "must be a valid URL"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
