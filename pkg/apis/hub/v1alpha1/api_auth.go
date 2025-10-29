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

// APIAuth defines the authentication configuration for APIs.
type APIAuth struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APIAuth.
	Spec APIAuthSpec `json:"spec,omitempty"`

	// The current status of this APIAuth.
	// +optional
	Status APIAuthStatus `json:"status,omitempty"`
}

// APIAuthSpec configures the authentication for APIs.
// +kubebuilder:validation:XValidation:message="exactly one authentication method must be specified",rule="[has(self.apiKey), has(self.jwt), has(self.ldap)].filter(x, x).size() == 1"
type APIAuthSpec struct {
	// IsDefault specifies if this APIAuth should be used as the default API authentication method for the namespace.
	// Only one APIAuth per namespace should have isDefault set to true.
	IsDefault bool `json:"isDefault"`

	// APIKey configures API key authentication.
	// +optional
	APIKey *APIKeyAuthSpec `json:"apiKey,omitempty"`

	// JWT configures JWT authentication.
	// +optional
	JWT *JWTAuthSpec `json:"jwt,omitempty"`

	// LDAP configures LDAP authentication.
	// +optional
	LDAP *LDAPConnectionConfig `json:"ldap,omitempty"`
}

// APIKeyAuthSpec configures API key authentication.
// +kubebuilder:pruning:PreserveUnknownFields
// PreserveUnknownFields annotation is needed because this is an empty struct,
// which would generate an invalid OpenAPI schema without explicit properties.
type APIKeyAuthSpec struct{}

// JWTAuthSpec configures JWT authentication.
// +kubebuilder:validation:XValidation:message="exactly one of signingSecretName, publicKey, jwksFile, or jwksUrl must be specified",rule="[has(self.signingSecretName), has(self.publicKey), has(self.jwksFile), has(self.jwksUrl)].filter(x, x).size() == 1"
type JWTAuthSpec struct {
	// StripAuthorizationHeader determines whether to strip the Authorization header before forwarding the request.
	// +optional
	StripAuthorizationHeader bool `json:"stripAuthorizationHeader,omitempty"`

	// TokenQueryKey specifies the query parameter name for the JWT token.
	// +optional
	TokenQueryKey string `json:"tokenQueryKey,omitempty"`

	// AppIDClaim is the name of the claim holding the identifier of the application.
	// This field is sometimes named `client_id`.
	AppIDClaim string `json:"appIdClaim"`

	// TokenNameClaim is the name of the claim holding the name of the token.
	// This name, if provided, will be used in the metrics.
	// +optional
	TokenNameClaim string `json:"tokenNameClaim,omitempty"`

	// ForwardHeaders specifies additional headers to forward with the request.
	// +optional
	ForwardHeaders map[string]string `json:"forwardHeaders,omitempty"`

	// SigningSecretName is the name of the Kubernetes Secret containing the signing secret.
	// The secret must be of type Opaque and contain a key named 'value'.
	// +optional
	// +kubebuilder:validation:MaxLength=253
	SigningSecretName string `json:"signingSecretName,omitempty"`

	// PublicKey is the PEM-encoded public key for JWT verification.
	// +optional
	PublicKey string `json:"publicKey,omitempty"`

	// JWKSFile contains the JWKS file content for JWT verification.
	// +optional
	JWKSFile string `json:"jwksFile,omitempty"`

	// JWKSURL is the URL to fetch the JWKS for JWT verification.
	// +optional
	// +kubebuilder:validation:XValidation:message="must be a valid URL",rule="isURL(self)"
	JWKSURL string `json:"jwksUrl,omitempty"`
}

// APIAuthStatus is the status of an APIAuth.
type APIAuthStatus struct {
	Version  string       `json:"version,omitempty"`
	SyncedAt *metav1.Time `json:"syncedAt,omitempty"`

	// Hash is a hash representing the APIAuth.
	Hash string `json:"hash,omitempty"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIAuthList defines a list of APIAuth.
type APIAuthList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APIAuth `json:"items"`
}
