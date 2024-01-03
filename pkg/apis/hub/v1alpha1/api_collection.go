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
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APICollection defines a collection of APIs exposed within an APIPortal.
// +kubebuilder:printcolumn:name="PathPrefix",type=string,JSONPath=`.spec.pathPrefix`
// +kubebuilder:printcolumn:name="APISelector",type=string,JSONPath=`.status.apiSelector`
// +kubebuilder:resource:scope=Cluster
type APICollection struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec APICollectionSpec `json:"spec,omitempty"`

	// The current status of this APICollection.
	// +optional
	Status APICollectionStatus `json:"status,omitempty"`
}

// APICollectionSpec configures an APICollection.
type APICollectionSpec struct {
	// PathPrefix is the path prefix under which all selected APIs will be exposed.
	// +optional
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:XValidation:message="must start with a '/'",rule="self.startsWith('/')"
	// +kubebuilder:validation:XValidation:message="cannot contains '../'",rule="!self.matches(r\"\"\"(\\/\\.\\.\\/)|(\\/\\.\\.$)\"\"\")"
	PathPrefix string `json:"pathPrefix,omitempty"`

	// APISelector selects the APIs which are member of this APICollection object.
	// Multiple APICollections can select the same set of APIs.
	// This field is optional and follows standard label selector semantics.
	// An empty APISelector matches any API.
	// +optional
	APISelector *metav1.LabelSelector `json:"apiSelector"`

	// APIs defines a set of APIs which are member of this APICollection object.
	// Multiple APICollections can select the same APIs.
	// When combined with APISelector, this set of APIs is appended to the matching APIs.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated apis",rule="self.all(x, self.exists_one(y, x.name == y.name && (has(x.__namespace__) && x.__namespace__ != '' ? x.__namespace__ : 'default') == (has(y.__namespace__) && y.__namespace__ != '' ? y.__namespace__ : 'default')))"
	APIs []APIReference `json:"apis,omitempty"`
}

// APICollectionStatus is the status of an APICollection.
type APICollectionStatus struct {
	APISelector string      `json:"apiSelector,omitempty"`
	Version     string      `json:"version,omitempty"`
	SyncedAt    metav1.Time `json:"syncedAt,omitempty"`

	// Hash is a hash representing the APICollection.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APICollectionList defines a list of APICollections.
type APICollectionList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APICollection `json:"items"`
}
