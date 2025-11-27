/*
Copyright (C) 2022-2025 Traefik Labs

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

// ManagedApplication represents a managed application.
// +kubebuilder:subresource:status
type ManagedApplication struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ManagedApplicationSpec `json:"spec,omitempty"`

	// The current status of this ManagedApplication.
	// +optional
	Status ManagedApplicationStatus `json:"status,omitempty"`
}

// ManagedApplicationSpec describes the ManagedApplication.
type ManagedApplicationSpec struct {
	// AppID is the identifier of the ManagedApplication.
	// It should be unique.
	// +kubebuilder:validation:MaxLength=253
	AppID string `json:"appId"`

	// Owner represents the owner of the ManagedApplication.
	// It should be:
	// - `sub` when using OIDC
	// - `externalID` when using external IDP
	// +kubebuilder:validation:MaxLength=253
	Owner string `json:"owner"`

	// Notes contains notes about application.
	// +optional
	Notes string `json:"notes,omitempty"`

	// APIKeys references the API keys used to authenticate the application when calling APIs.
	// +kubebuilder:validation:MaxItems=100
	// +optional
	APIKeys []APIKey `json:"apiKeys,omitempty"`
}

// APIKey describes an API key used to authenticate the application when calling APIs.
// +kubebuilder:validation:XValidation:message="secretName and value are mutually exclusive",rule="[has(self.secretName), has(self.value)].filter(x, x).size() <= 1"
type APIKey struct {
	// SecretName references the name of the secret containing the API key.
	// +kubebuilder:validation:MaxLength=253
	// +optional
	SecretName string `json:"secretName,omitempty"`

	// Value is the API key value.
	// +kubebuilder:validation:MaxLength=4096
	// +optional
	Value string `json:"value,omitempty"`

	// +optional
	Title string `json:"title"`

	// +optional
	Suspended bool `json:"suspended,omitempty"`
}

// ManagedApplicationStatus is the status of the ManagedApplication.
type ManagedApplicationStatus struct {
	Version        string            `json:"version,omitempty"`
	APIKeyVersions map[string]string `json:"apiKeyVersions,omitempty"`
	SyncedAt       *metav1.Time      `json:"syncedAt,omitempty"`

	// Hash is a hash representing the ManagedApplication.
	Hash string `json:"hash,omitempty"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedApplicationList defines a list of ManagedApplication.
type ManagedApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ManagedApplication `json:"items"`
}
