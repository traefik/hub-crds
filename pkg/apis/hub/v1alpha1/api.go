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

// API defines an HTTP interface that is exposed to external clients. It specifies the supported versions
// and provides instructions for accessing its documentation. Once instantiated, an API object is associated
// with an Ingress, IngressRoute, or HTTPRoute resource, enabling the exposure of the described API to the outside world.
type API struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec APISpec `json:"spec,omitempty"`

	// The current status of this API.
	// +optional
	Status APIStatus `json:"status,omitempty"`
}

// APISpec describes the API.
type APISpec struct {
	// OpenAPISpec defines the API contract as an OpenAPI specification.
	// +optional
	// +kubebuilder:validation:XValidation:message="path or url must be defined",rule="has(self.path) || has(self.url)"
	OpenAPISpec *OpenAPISpec `json:"openApiSpec,omitempty"`

	// Versions are the different APIVersions available.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:MinItems=1
	Versions []APIVersionRef `json:"versions,omitempty"`
}

// APIVersionRef references an APIVersion.
type APIVersionRef struct {
	// Name of the APIVersion.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
}

// OpenAPISpec defines the API contract as an OpenAPI specification.
type OpenAPISpec struct {
	// URL is a Traefik Hub agent accessible URL for obtaining the OpenAPI specification.
	// The URL must be accessible via a GET request method and should serve a YAML or JSON document containing the OpenAPI specification.
	// +optional
	// +kubebuilder:validation:XValidation:message="must be a valid URL",rule="isURL(self)"
	URL string `json:"url,omitempty"`

	// Path specifies the endpoint path within the Kubernetes Service where the OpenAPI specification can be obtained.
	// The Service queried is determined by the associated Ingress, IngressRoute, or HTTPRoute resource to which the API is attached.
	// It's important to note that this option is incompatible if the Ingress or IngressRoute specifies multiple backend services.
	// The Path must be accessible via a GET request method and should serve a YAML or JSON document containing the OpenAPI specification.
	// +optional
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:XValidation:message="must start with a '/'",rule="self.startsWith('/')"
	// +kubebuilder:validation:XValidation:message="cannot contains '../'",rule="!self.matches(r\"\"\"(\\/\\.\\.\\/)|(\\/\\.\\.$)\"\"\")"
	Path string `json:"path,omitempty"`

	// OperationSets defines the sets of operations to be referenced for granular filtering in APIAccesses.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	OperationSets []OperationSet `json:"operationSets,omitempty"`
}

// OperationSet gives a name to a set of matching OpenAPI operations.
// This set of operations can then be referenced for granular filtering in APIAccesses.
type OperationSet struct {
	// Name is the name of the OperationSet to reference in APIAccesses.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`

	// Matchers defines a list of alternative rules for matching OpenAPI operations.
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:MinItems=1
	Matchers []OperationMatcher `json:"matchers"`
}

// OperationMatcher defines criteria for matching an OpenAPI operation.
// +kubebuilder:validation:MinProperties=1
// +kubebuilder:validation:XValidation:message="path, pathPrefix and pathRegex are mutually exclusive",rule="[has(self.path), has(self.pathPrefix), has(self.pathRegex)].filter(x, x).size() <= 1"
type OperationMatcher struct {
	// Path specifies the exact path of the operations to select.
	// +optional
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:XValidation:message="must start with a '/'",rule="self.startsWith('/')"
	// +kubebuilder:validation:XValidation:message="cannot contains '../'",rule="!self.matches(r\"\"\"(\\/\\.\\.\\/)|(\\/\\.\\.$)\"\"\")"
	Path string `json:"path,omitempty"`

	// PathPrefix specifies the path prefix of the operations to select.
	// +optional
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:XValidation:message="must start with a '/'",rule="self.startsWith('/')"
	// +kubebuilder:validation:XValidation:message="cannot contains '../'",rule="!self.matches(r\"\"\"(\\/\\.\\.\\/)|(\\/\\.\\.$)\"\"\")"
	PathPrefix string `json:"pathPrefix,omitempty"`

	// PathRegex specifies a regular expression pattern for matching operations based on their paths.
	// +optional
	PathRegex string `json:"pathRegex,omitempty"`

	// Methods specifies the HTTP methods to be included for selection.
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Methods *[]string `json:"methods,omitempty"`
}

// APIStatus is the status of the API.
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
