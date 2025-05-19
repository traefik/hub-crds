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

	// TokenSecrets references token secrets.
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated secrets",rule="self.all(x, self.exists_one(y, x.name == y.name))"
	// +optional
	TokenSecrets []TokenSecretRef `json:"tokenSecrets,omitempty"`
}

// TokenSecretRef represents a reference to a kubernetes secret.
type TokenSecretRef struct {
	// Name of the Secret.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
}

// ManagedApplicationStatus is the status of the ManagedApplication.
type ManagedApplicationStatus struct {
	Version             string            `json:"version,omitempty"`
	TokenSecretVersions map[string]string `json:"tokenSecretVersions,omitempty"`
	SyncedAt            *metav1.Time      `json:"syncedAt,omitempty"`
	// Hash is a hash representing the ManagedApplication.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedApplicationList defines a list of ManagedApplication.
type ManagedApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ManagedApplication `json:"items"`
}
