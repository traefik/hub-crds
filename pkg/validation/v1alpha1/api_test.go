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
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestAPI_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  title: My API Title
  description: My API Description
  cors:
    maxAge: 3
    allowOriginListRegex:
      - ".*"
    allowCredentials: true
    addVaryHeader: true
    exposeHeadersList:
      - "*"
    allowOriginsList:
      - "http://example.com"
    allowMethodsList:
      - POST
    allowHeadersList:
      - "*"
  openApiSpec:
    path: /openapi.json
    validateRequestMethodAndPath: true
    operationSets:
    - name: my-operation-set
      matchers:
        - path: /foo
          methods:
            - GET
            - OPTION`),
		},
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: "my-api"`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: .non-dns-compliant-api
  namespace: my-ns`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-api", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: ""
  namespace: my-ns
spec:
  versions:
    - name: my-api-v2`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: api-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: my-ns
spec:
  versions:
    - name: my-api-v2`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "api-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "title is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  title: Way Toooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo Long Title`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: "spec.title", BadValue: "<value omitted>", Detail: "may not be more than 253 bytes"}},
		},
		{
			desc: "openApiSpec must have a path or an url",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec: {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec", BadValue: "object", Detail: "path or url must be defined"}},
		},

		{
			desc: "openApiSpec url must be a valid URL",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    url: ../invalid-spec-url.json`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.url", BadValue: "string", Detail: "must be a valid URL"}},
		},
		{
			desc: "openApiSpec path must start with a /",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: something`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.path", BadValue: "string", Detail: "must start with a '/'"}},
		},
		{
			desc: "openApiSpec path cannot contains ../",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo/../bar`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.path", BadValue: "string", Detail: "cannot contains '../'"}},
		},
		{
			desc: "openApiSpec path cannot ends with /..",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo/..`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.path", BadValue: "string", Detail: "cannot contains '../'"}},
		},
		{
			desc: "override server url is broken",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    override:
      servers:
      - url: aaa`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.override.servers[0].url", BadValue: "string", Detail: "must be a valid URL"}},
		},
		{
			desc: "valid: openApiSpec path with segment starting with ..",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo/..bar`),
		},
		{
			desc: "missing operationSet name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - matchers:
        - path: /foo`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.openApiSpec.operationSets[0].name", BadValue: "", Detail: ""}},
		},
		{
			desc: "operationSet with empty matchers",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers: []`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers", BadValue: int64(0), Detail: "spec.openApiSpec.operationSets[0].matchers in body should have at least 1 items"}},
		},
		{
			desc: "operationSet with no matcher",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.openApiSpec.operationSets[0].matchers", BadValue: "", Detail: ""}},
		},
		{
			desc: "operationSet with an empty matcher",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0]", BadValue: int64(0), Detail: "spec.openApiSpec.operationSets[0].matchers[0] in body should have at least 1 properties"}},
		},
		{
			desc: "operationSet matcher path and pathPrefix are mutually exclusive",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - path: /foo
            pathPrefix: /foo`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0]", BadValue: "object", Detail: "path, pathPrefix and pathRegex are mutually exclusive"}},
		},
		{
			desc: "operationSet matcher path and pathRegex are mutually exclusive",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - path: /foo
            pathRegex: /.*/foo`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0]", BadValue: "object", Detail: "path, pathPrefix and pathRegex are mutually exclusive"}},
		},
		{
			desc: "operationSet matcher pathPrefix and pathRegex are mutually exclusive",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - pathPrefix: /foo
            pathRegex: /.*/foo`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0]", BadValue: "object", Detail: "path, pathPrefix and pathRegex are mutually exclusive"}},
		},
		{
			desc: "operationSet matcher path, pathPrefix and pathRegex are mutually exclusive",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - path: /foo
            pathPrefix: /foo
            pathRegex: /.*/foo`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0]", BadValue: "object", Detail: "path, pathPrefix and pathRegex are mutually exclusive"}},
		},
		{
			desc: "operationSet matcher path must start with a /",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - path: something`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0].path", BadValue: "string", Detail: "must start with a '/'"}},
		},
		{
			desc: "operationSet matcher path cannot contains ../",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - path: /foo/../bar`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0].path", BadValue: "string", Detail: "cannot contains '../'"}},
		},
		{
			desc: "operationSet matcher path cannot ends with /..",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - path: /foo/..`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0].path", BadValue: "string", Detail: "cannot contains '../'"}},
		},
		{
			desc: "valid: operationSet matcher path with segment starting with ..",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - path: /foo/..bar`),
		},
		{
			desc: "operationSet matcher pathPrefix must start with a /",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - pathPrefix: something`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0].pathPrefix", BadValue: "string", Detail: "must start with a '/'"}},
		},
		{
			desc: "operationSet matcher pathPrefix cannot contains ../",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - pathPrefix: /foo/../bar`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0].pathPrefix", BadValue: "string", Detail: "cannot contains '../'"}},
		},
		{
			desc: "operationSet matcher pathPrefix cannot ends with /..",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - pathPrefix: /foo/..`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.openApiSpec.operationSets[0].matchers[0].pathPrefix", BadValue: "string", Detail: "cannot contains '../'"}},
		},
		{
			desc: "valid: operationSet matcher pathPrefix with segment starting with ..",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - pathPrefix: /foo/..bar`),
		},
		{
			desc: "valid: operationSet matcher with methods only",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  openApiSpec:
    path: /foo
    operationSets:
      - name: my-operation-set
        matchers:
          - methods:
            - GET`),
		},
		{
			desc: "valid: empty version ..",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: API
metadata:
  name: my-api
  namespace: my-ns
spec:
  versions: []`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.versions", BadValue: int64(0), Detail: "spec.versions in body should have at least 1 items"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
