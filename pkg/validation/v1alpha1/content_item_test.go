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

func TestContentItem_Validation(t *testing.T) {
	t.Parallel()

	tooLongName := strings.Repeat("x", 254)

	tests := []validationTestCase{
		{
			desc: "valid: with content",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
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
			desc: "valid: with link",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
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
			desc: "valid: parentRef kind APIBundle",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: APIBundle
    name: my-bundle
  content: "# Hello World"`),
		},
		{
			desc: "valid: parentRef kind APIPortal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: APIPortal
    name: my-portal
  content: "# Hello World"`),
		},
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
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
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
metadata:
  name: .non-dns-compliant-content
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-content", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
metadata:
  name: ""
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
metadata:
  name: content-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "content-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "title is empty",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
metadata:
  name: my-content
  namespace: default
spec:
  title: ""
  order: 0
  parentRef:
    kind: API
    name: my-api
  content: "# Hello World"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.title", BadValue: "", Detail: "spec.title in body should be at least 1 chars long"}},
		},
		{
			desc: "title is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
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
			desc: "invalid parentRef kind",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: 0
  parentRef:
    kind: APIVersion
    name: my-api
  content: "# Hello World"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeNotSupported, Field: "spec.parentRef.kind", BadValue: "APIVersion", Detail: "supported values: \"APIPortal\", \"API\", \"APIBundle\""}},
		},
		{
			desc: "parentRef name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
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
kind: ContentItem
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
kind: ContentItem
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
kind: ContentItem
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
		{
			desc: "invalid order",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: ContentItem
metadata:
  name: my-content
  namespace: default
spec:
  title: My Content
  order: -1
  parentRef:
    kind: API
    name: my-api
  content: "# Hello World"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.order", BadValue: int64(-1), Detail: "spec.order in body should be greater than or equal to 0"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
