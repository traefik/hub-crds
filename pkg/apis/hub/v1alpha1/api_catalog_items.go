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

// APICatalogItems defines APIs that will be part of the API catalog on the portal.
type APICatalogItems struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APICatalogItems.
	// +kubebuilder:validation:XValidation:message="groups and everyone are mutually exclusive",rule="(has(self.everyone) && has(self.groups)) ? !(self.everyone && self.groups.size() > 0) : true"
	Spec APICatalogItemsSpec `json:"spec,omitempty"`

	// The current status of this APICatalogItems.
	// +optional
	Status APICatalogItemsStatus `json:"status,omitempty"`
}

// APICatalogItemsSpec configures an APICatalogItems.
type APICatalogItemsSpec struct {
	// Groups are the consumer groups that will see the APIs.
	// +optional
	Groups []string `json:"groups,omitempty"`

	// Everyone indicates that all users will see these APIs.
	// +optional
	Everyone bool `json:"everyone,omitempty"`

	// APIBundles defines a set of APIBundle that will be visible to the configured audience.
	// Multiple APICatalogItems can select the same APIBundles.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated apiBundles",rule="self.all(x, self.exists_one(y, x.name == y.name))"
	APIBundles []APIBundleReference `json:"apiBundles,omitempty"`

	// APISelector selects the APIs that will be visible to the configured audience.
	// Multiple APICatalogItems can select the same set of APIs.
	// This field is optional and follows standard label selector semantics.
	// An empty APISelector matches any API.
	// +optional
	APISelector *metav1.LabelSelector `json:"apiSelector,omitempty"`

	// APIs defines a set of APIs that will be visible to the configured audience.
	// Multiple APICatalogItems can select the same APIs.
	// When combined with APISelector, this set of APIs is appended to the matching APIs.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated apis",rule="self.all(x, self.exists_one(y, x.name == y.name))"
	APIs []APIReference `json:"apis,omitempty"`

	// OperationFilter specifies the visible operations on APIs and APIVersions.
	// If not set, all operations are available.
	// An empty OperationFilter prohibits all operations.
	// +optional
	OperationFilter *OperationFilter `json:"operationFilter,omitempty"`

	// APIPlan defines which APIPlan will be available.
	// If multiple APICatalogItems specify the same API with different APIPlan, the API consumer will be able to pick
	// a plan from this list.
	// +optional
	APIPlan *APIPlanReference `json:"apiPlan,omitempty"`
}

// APICatalogItemsStatus is the status of an APICatalogItems.
type APICatalogItemsStatus struct {
	Version  string       `json:"version,omitempty"`
	SyncedAt *metav1.Time `json:"syncedAt,omitempty"`

	// Hash is a hash representing the APICatalogItems.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APICatalogItemsList defines a list of APICatalogItems.
type APICatalogItemsList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APICatalogItems `json:"items"`
}
