/*
Copyright (C) 2022-2023 Traefik Labs

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
// +kubebuilder:printcolumn:name="CurrentVersion",type=string,JSONPath=`.spec.currentVersion`
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
	PathPrefix string `json:"pathPrefix"`
	// +optional
	Service *APIService `json:"service,omitempty"`
	// +optional
	Headers *Headers `json:"headers,omitempty"`
	// +optional
	CORS *CORS `json:"cors,omitempty"`
	// +optional
	CurrentVersion string `json:"currentVersion,omitempty"`
}

// APIService configures the service to exposed on the edge.
type APIService struct {
	Name string `json:"name"`
	// port of the referenced service. A port name or port number
	// is required for an APIServiceBackendPort.
	Port        APIServiceBackendPort `json:"port"`
	OpenAPISpec OpenAPISpec           `json:"openApiSpec,omitempty"`
}

// APIServiceBackendPort is the service port being referenced.
type APIServiceBackendPort struct {
	// name is the name of the port on the Service.
	// This must be an IANA_SVC_NAME (following RFC6335).
	// This is a mutually exclusive setting with "Number".
	// +optional
	Name string `json:"name"`

	// number is the numerical port number (e.g. 80) on the Service.
	// This is a mutually exclusive setting with "Name".
	// +optional
	Number int32 `json:"number"`
}

// OpenAPISpec defines the OpenAPI spec of an API.
type OpenAPISpec struct {
	// +optional
	URL string `json:"url,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	Port *APIServiceBackendPort `json:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty"`
}

// Headers configures the requests and responses headers manipulations.
type Headers struct {
	// +optional
	Request *HeadersConfig `json:"request,omitempty"`
	// +optional
	Response *HeadersConfig `json:"response,omitempty"`
}

// HeadersConfig configures headers manipulations.
type HeadersConfig struct {
	// +optional
	Set map[string]string `json:"set,omitempty"`
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
