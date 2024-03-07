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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// API defines an API exposed within a portal.
// +kubebuilder:printcolumn:name="PathPrefix",type=string,JSONPath=`.spec.pathPrefix`
// +kubebuilder:printcolumn:name="ServiceName",type=string,JSONPath=`.spec.service.name`
// +kubebuilder:printcolumn:name="ServicePort",type=string,JSONPath=`.spec.service.port.number`
// +kubebuilder:printcolumn:name="Versions",type=string,JSONPath=`.spec.currentVersion`
type API struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec APISpec `json:"spec,omitempty"`

	// The current status of this API.
	// +optional
	Status APIStatus `json:"status,omitempty"`
}

// APISpec configures an API.
type APISpec struct {
	// OpenAPISpec defines where to obtain the OpenAPI specification of the Service.
	// +optional
	// +kubebuilder:validation:XValidation:message="path or url must be defined",rule="has(self.path) || has(self.url)"
	OpenAPISpec *OpenAPISpec `json:"openApiSpec,omitempty"`

	// Versions defines the current APIVersion.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:MinItems=1
	Versions []APIVersionRef `json:"versions,omitempty"`
}

// APIVersionRef holds an APIVersion name.
type APIVersionRef struct {
	Name string `json:"name"`
}

// OpenAPISpec defines the OpenAPI spec of an API.
type OpenAPISpec struct {
	// URL is a Traefik Hub agent accessible URL for obtaining the specification.
	// This URL must be queryable with a GET method and serve a YAML or JSON document.
	// +optional
	// +kubebuilder:validation:XValidation:message="must be a valid URL",rule="isURL(self)"
	URL string `json:"url,omitempty"`

	// Path is the path on the Kubernetes Service for obtaining the specification.
	// This Path must be queryable with a GET method and serve a YAML or JSON document.
	// +optional
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:XValidation:message="must start with a '/'",rule="self.startsWith('/')"
	// +kubebuilder:validation:XValidation:message="cannot contains '../'",rule="!self.matches(r\"\"\"(\\/\\.\\.\\/)|(\\/\\.\\.$)\"\"\")"
	Path string `json:"path,omitempty"`

	// OperationSets defines the sets of operations that can be used for advanced filtering in APIAccesses.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	OperationSets []OperationSet `json:"operationSets,omitempty"`
}

// OperationSet selects a set of OpenAPI operations that can be referenced for advanced filtering on APIAccesses.
type OperationSet struct {
	// Name is the name of the OperationSet.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`

	// Matchers defines a set of OperationMatchers that selects spec operations.
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:MinItems=1
	Matchers []OperationMatcher `json:"matchers"`
}

// OperationMatcher selects the operations that will be part of the OperationSet.
// +kubebuilder:validation:MinProperties=1
// +kubebuilder:validation:XValidation:message="path, pathPrefix and pathRegex are mutually exclusive",rule="[has(self.path), has(self.pathPrefix), has(self.pathRegex)].filter(x, x).size() <= 1"
type OperationMatcher struct {
	// Path defines the exact path of the spec operations to select.
	// +optional
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:XValidation:message="must start with a '/'",rule="self.startsWith('/')"
	// +kubebuilder:validation:XValidation:message="cannot contains '../'",rule="!self.matches(r\"\"\"(\\/\\.\\.\\/)|(\\/\\.\\.$)\"\"\")"
	Path string `json:"path,omitempty"`

	// PathPrefix defines the path prefix of the spec operations to select.
	// +optional
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:XValidation:message="must start with a '/'",rule="self.startsWith('/')"
	// +kubebuilder:validation:XValidation:message="cannot contains '../'",rule="!self.matches(r\"\"\"(\\/\\.\\.\\/)|(\\/\\.\\.$)\"\"\")"
	PathPrefix string `json:"pathPrefix,omitempty"`

	// PathRegex defines the path regex of the matching spec operations to select.
	// +optional
	PathRegex string `json:"pathRegex,omitempty"`

	// Methods defines a set of methods of the specs operation to select.
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Methods *[]string `json:"methods,omitempty"`
}

// Headers configures the requests and responses headers manipulations.
type Headers struct {
	// Request configures the request headers.
	// +optional
	Request *HeadersConfig `json:"request,omitempty"`

	// Response configures the response headers.
	// +optional
	Response *HeadersConfig `json:"response,omitempty"`
}

// HeadersConfig configures headers manipulations.
type HeadersConfig struct {
	// Set sets the value of headers
	// +optional
	Set map[string]string `json:"set,omitempty"`

	// Delete deletes headers.
	// +optional
	Delete []string `json:"delete,omitempty"`
}

// CORS configures the CORS for the API.
type CORS struct {
	// AllowCredentials defines whether the request can include user credentials.
	// +optional
	AllowCredentials bool `json:"allowCredentials"`

	// AllowHeaders defines the Access-Control-Request-Headers values sent in preflight response.
	// +optional
	AllowHeaders []string `json:"allowHeaders,omitempty"`

	// AllowMethods defines the Access-Control-Request-Method values sent in preflight response.
	// +optional
	AllowMethods []string `json:"allowMethods,omitempty"`

	// AllowOriginList is a list of allowable origins. Can also be a wildcard origin "*".
	// +optional
	AllowOriginList []string `json:"allowOriginList,omitempty"`

	// AllowOriginListRegex is a list of allowable origins written following the Regular Expression syntax (https://golang.org/pkg/regexp/).
	// +optional
	AllowOriginListRegex []string `json:"allowOriginListRegex,omitempty"`

	// ExposeHeaders defines the Access-Control-Expose-Headers values sent in preflight response.
	// +optional
	ExposeHeaders []string `json:"exposeHeaders,omitempty"`

	// MaxAge defines the time that a preflight request may be cached.
	// +optional
	MaxAge int64 `json:"maxAge"`
}

// APIStatus is the status of an API.
type APIStatus struct {
	Version  string      `json:"version,omitempty"`
	SyncedAt metav1.Time `json:"syncedAt,omitempty"`
	// Hash is a hash representing the API.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIList defines a list of APIs.
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []API `json:"items"`
}
