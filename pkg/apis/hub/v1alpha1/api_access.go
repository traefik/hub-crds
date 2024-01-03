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

// APIAccess defines which group of consumers can access APIs and APICollections.
// +kubebuilder:resource:scope=Cluster
type APIAccess struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APIAccess.
	// +kubebuilder:validation:XValidation:message="groups and anyGroups are mutually exclusive",rule="(has(self.anyGroups) && has(self.groups)) ? !(self.anyGroups && self.groups.size() > 0) : true"
	Spec APIAccessSpec `json:"spec,omitempty"`

	// The current status of this APIAccess.
	// +optional
	Status APIAccessStatus `json:"status,omitempty"`
}

// APIAccessSpec configures an APIAccess.
type APIAccessSpec struct {
	// Groups are the user groups that will gain access to the selected APIs.
	// +optional
	Groups []string `json:"groups"`

	// AnyGroups states that everyone will gain access to the selected APIs.
	// +optional
	AnyGroups bool `json:"anyGroups"`

	// APISelector selects the APIs that will be accessible to the configured user groups.
	// Multiple APIAccesses can select the same set of APIs.
	// This field is optional and follows standard label selector semantics.
	// An empty APISelector matches any API.
	// +optional
	APISelector *metav1.LabelSelector `json:"apiSelector,omitempty"`

	// APIs defines a set of APIs that will be accessible to the configured user groups.
	// Multiple APIAccesses can select the same APIs.
	// When combined with APISelector, this set of APIs is appended to the matching APIs.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated apis",rule="self.all(x, self.exists_one(y, x.name == y.name && (has(x.__namespace__) && x.__namespace__ != '' ? x.__namespace__ : 'default') == (has(y.__namespace__) && y.__namespace__ != '' ? y.__namespace__ : 'default')))"
	APIs []APIReference `json:"apis,omitempty"`

	// APICollectionSelector selects the APICollections that will be accessible to the configured user groups.
	// Multiple APIAccesses can select the same set of APICollections.
	// This field is optional and follows standard label selector semantics.
	// An empty APICollectionSelector matches any APICollection.
	// +optional
	APICollectionSelector *metav1.LabelSelector `json:"apiCollectionSelector,omitempty"`

	// APICollections defines a set of APICollections that will be accessible to the configured user groups.
	// Multiple APIAccesses can select the same APICollections.
	// When combined with APICollectionSelector, this set of APICollections is appended to the matching APICollections.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated collections",rule="self.all(x, self.exists_one(y, x.name == y.name))"
	APICollections []APICollectionReference `json:"apiCollections,omitempty"`

	// OperationFilter selects the OperationSets defined on an API or an APIVersion that will be accessible to the configured user groups.
	// If not set, all spec operations will be accessible.
	// An empty OperationFilter matches no OperationSet.
	// +optional
	OperationFilter *OperationFilter `json:"operationFilter,omitempty"`
}

// APIReference contains information to identify an API to add to an APIAccess, an APICollection or an APIRateLimit.
type APIReference struct {
	// Name of the API.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`

	// Namespace of the API.
	// +optional
	// +kubebuilder:validation:MaxLength=63
	Namespace string `json:"namespace,omitempty"`
}

// APICollectionReference contains information to identify an APICollection to add to APIAccess.
type APICollectionReference struct {
	// Name of the APICollection.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
}

// OperationFilter contains information to select OperationSets defined on an API or an APIVersion.
type OperationFilter struct {
	// Include defines the names of OperationSets that will be accessible.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Include []string `json:"include,omitempty"`
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
