/*
Copyright (C) 2022-2026 Traefik Labs

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

// ManagedSubscription defines a Subscription managed by the API manager as the result of a pre-negotiation with its
// API consumers. This subscription grant consuming access to a set of APIs to a set of Applications.
// +kubebuilder:subresource:status
type ManagedSubscription struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this ManagedSubscription.
	Spec ManagedSubscriptionSpec `json:"spec,omitempty"`

	// The current status of this ManagedSubscription.
	// +optional
	Status ManagedSubscriptionStatus `json:"status,omitempty"`
}

// ManagedSubscriptionSpec configures an ManagedSubscription.
type ManagedSubscriptionSpec struct {
	// Applications references the Applications that will gain access to the specified APIs.
	// Multiple ManagedSubscriptions can select the same AppID.
	//
	// Deprecated: Use ManagedApplications instead.
	//
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Applications []ApplicationReference `json:"applications,omitempty"`

	// ManagedApplications references the ManagedApplications that will gain access to the specified APIs.
	// Multiple ManagedSubscriptions can select the same ManagedApplication.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated managed applications",rule="self.all(x, self.exists_one(y, x.name == y.name))"
	ManagedApplications []ManagedApplicationReference `json:"managedApplications,omitempty"`

	// APIBundles defines a set of APIBundle that will be accessible.
	// Multiple ManagedSubscriptions can select the same APIBundles.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated apiBundles",rule="self.all(x, self.exists_one(y, x.name == y.name))"
	APIBundles []APIBundleReference `json:"apiBundles,omitempty"`

	// APISelector selects the APIs that will be accessible.
	// Multiple ManagedSubscriptions can select the same set of APIs.
	// This field is optional and follows standard label selector semantics.
	// An empty APISelector matches any API.
	// +optional
	APISelector *metav1.LabelSelector `json:"apiSelector,omitempty"`

	// APIs defines a set of APIs that will be accessible.
	// Multiple ManagedSubscriptions can select the same APIs.
	// When combined with APISelector, this set of APIs is appended to the matching APIs.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated apis",rule="self.all(x, self.exists_one(y, x.name == y.name))"
	APIs []APIReference `json:"apis,omitempty"`

	// OperationFilter specifies the allowed operations on APIs and APIVersions.
	// If not set, all operations are available.
	// An empty OperationFilter prohibits all operations.
	// +optional
	OperationFilter *OperationFilter `json:"operationFilter,omitempty"`

	// APIPlan defines which APIPlan will be used.
	APIPlan APIPlanReference `json:"apiPlan"`

	// Weight specifies the evaluation order of the APIPlan.
	// When multiple ManagedSubscriptions targets the same API and Application with different APIPlan,
	// the APIPlan with the highest weight will be enforced. If weights are equal, alphabetical order is used.
	// +kubebuilder:validation:XValidation:message="must be a positive number",rule="self >= 0"
	// +optional
	Weight int `json:"weight,omitempty"`

	// Claims specifies an expression that validate claims in order to authorize the request.
	// +optional
	Claims string `json:"claims,omitempty"`
}

// ManagedSubscriptionStatus is the status of an ManagedSubscription.
type ManagedSubscriptionStatus struct {
	Version  string       `json:"version,omitempty"`
	SyncedAt *metav1.Time `json:"syncedAt,omitempty"`

	// Hash is a hash representing the ManagedSubscription.
	Hash string `json:"hash,omitempty"`

	// Conditions is the list of status conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// ResolvedAPIs is the list of APIs that were successfully resolved.
	// +optional
	ResolvedAPIs []ResolvedAPIReference `json:"resolvedApis,omitempty"`

	// UnresolvedAPIs is the list of APIs that could not be resolved.
	// +optional
	UnresolvedAPIs []ResolvedAPIReference `json:"unresolvedApis,omitempty"`
}

// ApplicationReference references an Application.
type ApplicationReference struct {
	// AppID is the public identifier of the application.
	// In the case of OIDC, it corresponds to the clientId.
	// +kubebuilder:validation:MaxLength=253
	AppID string `json:"appId"`
}

// ManagedApplicationReference references a ManagedApplication.
type ManagedApplicationReference struct {
	// Name is the name of the ManagedApplication.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedSubscriptionList defines a list of ManagedSubscriptions.
type ManagedSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ManagedSubscription `json:"items"`
}
