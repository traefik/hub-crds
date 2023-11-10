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
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIAccess defines which group of consumers can access APIs and APICollections.
// +kubebuilder:resource:scope=Cluster
type APIAccess struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec APIAccessSpec `json:"spec,omitempty"`

	// The current status of this APIAccess.
	// +optional
	Status APIAccessStatus `json:"status,omitempty"`
}

// APIAccessSpec configures an APIAccess.
type APIAccessSpec struct {
	// +optional
	Groups []string `json:"groups"`
	// +optional
	AnyGroups bool `json:"anyGroups"`
	// APISelector selects the APIs which are member of this APIAccess object.
	// Multiple APIAccesses can select the same set of APIs.
	// This field is optional and follows standard label selector semantics.
	// An empty APISelector matches any API.
	// +optional
	APISelector *metav1.LabelSelector `json:"apiSelector,omitempty"`
	// APIs defines a set of APIs which are member of this APIAccess object.
	// Multiple APIAccesses can select the same APIs.
	// When combined with APISelector, this set of APIs is appended to the matching APIs.
	// +optional
	APIs []APIReference `json:"apis,omitempty"`
	// APICollectionSelector selects the APICollections which are member of this APIAccess object.
	// Multiple APIAccesses can select the same set of APICollections.
	// This field is optional and follows standard label selector semantics.
	// An empty APICollectionSelector matches any APICollection.
	// +optional
	APICollectionSelector *metav1.LabelSelector `json:"apiCollectionSelector,omitempty"`
	// APICollections defines a set of APICollections which are member of this APIAccess object.
	// Multiple APIAccesses can select the same APICollections.
	// When combined with APICollectionSelector, this set of APICollections is appended to the matching APICollections.
	// +optional
	APICollections []APICollectionReference `json:"apiCollections,omitempty"`
}

// APIReference contains information to identify an API to add to an APIAccess, an APICollection or an APIRateLimit.
type APIReference struct {
	Name string `json:"name"`
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// APICollectionReference contains information to identify an APICollection to add to APIAccess.
type APICollectionReference struct {
	Name string `json:"name"`
}

// APIAccessStatus is the status of an APIAccess.
type APIAccessStatus struct {
	Version  string      `json:"version,omitempty"`
	SyncedAt metav1.Time `json:"syncedAt,omitempty"`
	// Hash is a hash representing the APIAccess.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIAccessList defines a list of APIAccesses.
type APIAccessList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APIAccess `json:"items"`
}
