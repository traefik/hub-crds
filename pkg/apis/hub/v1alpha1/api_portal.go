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

// APIPortal defines a developer portal for accessing the documentation of APIs.
type APIPortal struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APIPortal.
	Spec APIPortalSpec `json:"spec,omitempty"`

	// The current status of this APIPortal.
	// +optional
	Status APIPortalStatus `json:"status,omitempty"`
}

// APIPortalSpec configures an APIPortal.
type APIPortalSpec struct {
	// Title is the public facing name of the APIPortal.
	// +optional
	Title string `json:"title,omitempty"`

	// Description of the APIPortal.
	// +optional
	Description string `json:"description,omitempty"`

	// TrustedURLs are the urls that are trusted by the OAuth 2.0 authorization server.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:XValidation:message="must be a valid URLs",rule="self.all(x, isURL(x))"
	TrustedURLs []string `json:"trustedUrls"`

	// UI holds the UI customization options.
	// +optional
	UI *UISpec `json:"ui,omitempty"`
}

// UISpec configures the UI customization.
type UISpec struct {
	// LogoURL is the public URL of the logo.
	// +optional
	LogoURL string `json:"logoUrl,omitempty"`
}

// OIDCConfigStatus is the OIDC configuration status.
type OIDCConfigStatus struct {
	// Generic indicates whether or not the APIPortal authentication relies on Generic OIDC.
	// +optional
	Generic bool `json:"generic,omitempty"`

	// Issuer is the OIDC issuer for accessing the exposed APIPortal WebUI.
	// +optional
	Issuer string `json:"issuer,omitempty"`

	// ClientID is the OIDC ClientID for accessing the exposed APIPortal WebUI.
	// +optional
	ClientID string `json:"clientId,omitempty"`

	// SecretName is the name of the secret containing the OIDC ClientSecret for accessing the exposed APIPortal WebUI.
	// +optional
	SecretName string `json:"secretName,omitempty"`

	// Scopes is the OIDC scopes for getting user attributes during the authentication to the exposed APIPortal WebUI.
	// +optional
	Scopes string `json:"scopes,omitempty"`

	// ExternalIDClaim is the name of the JWT claim containing the user external ID.
	// +optional
	ExternalIDClaim string `json:"externalIdClaim,omitempty"`

	// FirstnameClaim is the name of the JWT claim containing the user firstname.
	// +optional
	FirstnameClaim string `json:"firstnameClaim,omitempty"`

	// LastnameClaim is the name of the JWT claim containing the user lastname.
	// +optional
	LastnameClaim string `json:"lastnameClaim,omitempty"`

	// EmailClaim is the name of the JWT claim containing the user email.
	// +optional
	EmailClaim string `json:"emailClaim,omitempty"`

	// GroupsClaim is the name of the JWT claim containing the user groups.
	// +optional
	GroupsClaim string `json:"groupsClaim,omitempty"`

	// CompanyClaim is the name of the JWT claim containing the user groups.
	// +optional
	CompanyClaim string `json:"companyClaim,omitempty"`
}

// APIPortalStatus is the status of an APIPortal.
type APIPortalStatus struct {
	Version  string       `json:"version,omitempty"`
	SyncedAt *metav1.Time `json:"syncedAt,omitempty"`

	// OIDC is the OIDC configuration for accessing the exposed APIPortal WebUI.
	// +optional
	OIDC *OIDCConfigStatus `json:"oidc,omitempty"`

	// Hash is a hash representing the APIPortal.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIPortalList defines a list of APIPortals.
type APIPortalList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APIPortal `json:"items"`
}
