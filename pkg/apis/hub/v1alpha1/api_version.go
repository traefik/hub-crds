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

// APIVersion defines an APIVersion.
// +kubebuilder:printcolumn:name="APIName",type=string,JSONPath=`.spec.apiName`
// +kubebuilder:printcolumn:name="Title",type=string,JSONPath=`.spec.title`
// +kubebuilder:printcolumn:name="Release",type=string,JSONPath=`.spec.release`
// +kubebuilder:printcolumn:name="ServiceName",type=string,JSONPath=`.spec.service.name`
// +kubebuilder:printcolumn:name="ServicePort",type=string,JSONPath=`.spec.service.port.number`
type APIVersion struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APIVersion.
	Spec APIVersionSpec `json:"spec,omitempty"`

	// The current status of this APIVersion.
	// +optional
	Status APIVersionStatus `json:"status,omitempty"`
}

// APIVersionSpec configures an APIVersion.
type APIVersionSpec struct {
	// APIName is the name of the API this version belongs to.
	APIName string `json:"apiName"`

	// Title is the public facing name of the APIVersion.
	// +optional
	Title string `json:"title,omitempty"`

	// Release is the version number of the API.
	// This value must follow the SemVer format: https://semver.org/
	// +optional
	// +kubebuilder:validation:MaxLength=100
	// +kubebuilder:validation:XValidation:message="must be a valid semver version",rule="self.matches(r\"\"\"^v?(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$\"\"\")"
	Release string `json:"release,omitempty"`

	// StripPathPrefix strips the PathPrefix defined in the Routes when forwarding requests to the backend service.
	// +optional
	StripPathPrefix bool `json:"stripPathPrefix"`

	// Routes defines the different ways of accessing this APIVersion.
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Routes []Route `json:"routes,omitempty"`

	// Service defines the backend handling the incoming traffic.
	Service APIService `json:"service"`

	// Headers manipulates HTTP request and response headers.
	// +optional
	Headers *Headers `json:"headers,omitempty"`

	// CORS configures Cross-origin resource sharing headers.
	// +optional
	CORS *CORS `json:"cors,omitempty"`
}

// Route determines how to match the version.
type Route struct {
	// QueryParams defines the URL query parameters that must be present in the request to be routed on this APIVersion.
	// +optional
	QueryParams map[string]string `json:"queryParams,omitempty"`

	// Headers defines the HTTP headers that must be present in the request to be routed on this APIVersion.
	// +optional
	Headers map[string]string `json:"headers,omitempty"`

	// PathPrefix defines the path prefix to be routed to this APIVersion.
	// This PathPrefix is appended to the PathPrefix of the APICollection and API.
	// +optional
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:XValidation:message="must start with a '/'",rule="self.startsWith('/')"
	// +kubebuilder:validation:XValidation:message="cannot contains '../'",rule="!self.matches(r\"\"\"(\\/\\.\\.\\/)|(\\/\\.\\.$)\"\"\")"
	PathPrefix string `json:"pathPrefix,omitempty"`
}

// APIVersionStatus is the status of an APIVersion.
type APIVersionStatus struct {
	Version  string      `json:"version,omitempty"`
	SyncedAt metav1.Time `json:"syncedAt,omitempty"`
	// Hash is a hash representing the APIVersion.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIVersionList defines a list of APIVersionList.
type APIVersionList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APIVersion `json:"items"`
}
