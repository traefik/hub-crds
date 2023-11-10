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

	Spec APIVersionSpec `json:"spec,omitempty"`

	// The current status of this APIVersion.
	// +optional
	Status APIVersionStatus `json:"status,omitempty"`
}

// APIVersionSpec configures an APIVersion.
type APIVersionSpec struct {
	APIName string `json:"apiName"`
	// +optional
	Title string `json:"title,omitempty"`
	// +optional
	Release string `json:"release,omitempty"`
	// +optional
	StripPathPrefix bool `json:"stripPathPrefix"`
	// +optional
	Routes  []Route    `json:"routes,omitempty"`
	Service APIService `json:"service"`
	// +optional
	Headers *Headers `json:"headers,omitempty"`
	// +optional
	CORS *CORS `json:"cors,omitempty"`
}

// Route determines how to match the version.
type Route struct {
	// +optional
	QueryParams map[string]string `json:"queryParams,omitempty"`
	// +optional
	Headers map[string]string `json:"headers,omitempty"`
	// +optional
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
