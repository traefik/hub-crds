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

// APIPortalAuth defines the authentication configuration for an APIPortal.
type APIPortalAuth struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APIPortalAuth.
	Spec APIPortalAuthSpec `json:"spec,omitempty"`

	// The current status of this APIPortalAuth.
	// +optional
	Status APIPortalAuthStatus `json:"status,omitempty"`
}

// APIPortalAuthSpec configures the authentication for an APIPortal.
type APIPortalAuthSpec struct {
	// IssuerURL is the OIDC provider issuer URL.
	// +kubebuilder:validation:XValidation:message="must be a valid URL",rule="isURL(self)"
	IssuerURL string `json:"issuerUrl"`

	// SecretName is the name of the Kubernetes Secret containing clientId and clientSecret keys.
	// +kubebuilder:validation:MaxLength=253
	SecretName string `json:"secretName"`

	// Scopes is a list of OAuth2 scopes.
	// +optional
	Scopes []string `json:"scopes,omitempty"`

	// Claims configures JWT claim mappings for user attributes.
	// +optional
	Claims *ClaimsSpec `json:"claims,omitempty"`

	// SyncedAttributes is a list of additional attributes to sync from the OIDC provider.
	// Each attribute must correspond to a configured claim field.
	// +optional
	// +kubebuilder:validation:XValidation:message="syncedAttributes must only contain: groups, userId, firstname, lastname, email, company",rule="self.all(attr, attr in ['groups', 'userId', 'firstname', 'lastname', 'email', 'company'])"
	SyncedAttributes []string `json:"syncedAttributes,omitempty"`
}

// ClaimsSpec configures JWT claim mappings for user attributes.
type ClaimsSpec struct {
	// Groups is the JWT claim for user groups. This field is required for authorization.
	Groups string `json:"groups"`

	// UserID is the JWT claim for user ID mapping.
	// +optional
	UserID string `json:"userId,omitempty"`

	// Firstname is the JWT claim for user first name.
	// +optional
	Firstname string `json:"firstname,omitempty"`

	// Lastname is the JWT claim for user last name.
	// +optional
	Lastname string `json:"lastname,omitempty"`

	// Email is the JWT claim for user email.
	// +optional
	Email string `json:"email,omitempty"`

	// Company is the JWT claim for user company.
	// +optional
	Company string `json:"company,omitempty"`
}

// APIPortalAuthStatus is the status of an APIPortalAuth.
type APIPortalAuthStatus struct {
	Version  string       `json:"version,omitempty"`
	SyncedAt *metav1.Time `json:"syncedAt,omitempty"`

	// Hash is a hash representing the APIPortalAuth.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIPortalAuthList defines a list of APIPortalAuth.
type APIPortalAuthList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APIPortalAuth `json:"items"`
}
