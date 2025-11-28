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
// +kubebuilder:validation:XValidation:message="exactly one of oidc or ldap must be specified",rule="[has(self.oidc), has(self.ldap)].filter(x, x).size() == 1"
type APIPortalAuthSpec struct {
	// OIDC configures the OIDC authentication.
	// +optional
	OIDC *OIDCConfig `json:"oidc,omitempty"`

	// LDAP configures the LDAP authentication.
	// +optional
	LDAP *PortalAuthLDAPConfig `json:"ldap,omitempty"`
}

// OIDCConfig configures OIDC authentication for an APIPortal.
type OIDCConfig struct {
	// IssuerURL is the OIDC provider issuer URL.
	// +kubebuilder:validation:XValidation:message="must be a valid URL",rule="isURL(self)"
	IssuerURL string `json:"issuerUrl"`

	// SecretName is the name of the Kubernetes Secret containing clientId and clientSecret keys.
	// +kubebuilder:validation:MaxLength=253
	SecretName string `json:"secretName"`

	// Scopes is a list of OAuth2 scopes.
	// +optional
	Scopes *[]string `json:"scopes,omitempty"`

	// Claims configures JWT claim mappings for user attributes.
	Claims Claims `json:"claims"`

	// SyncedAttributes are the user attributes to synchronize with Hub platform.
	// +optional
	// +kubebuilder:validation:MaxItems=6
	// +kubebuilder:validation:items:Enum=groups;userId;firstname;lastname;email;company
	SyncedAttributes []string `json:"syncedAttributes,omitempty"`
}

// Claims configures JWT claim mappings for user attributes.
type Claims struct {
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

// LDAPConnectionConfig holds LDAP connection configuration.
type LDAPConnectionConfig struct {
	// URL is the URL of the LDAP server, including the protocol (ldap or ldaps) and the port.
	// +kubebuilder:validation:XValidation:message="must be a valid LDAP URL",rule="isURL(self) && (self.startsWith('ldap://') || self.startsWith('ldaps://'))"
	URL string `json:"url"`

	// StartTLS instructs the middleware to issue a StartTLS request when initializing the connection with the LDAP server.
	// +optional
	StartTLS bool `json:"startTls,omitempty"`

	// InsecureSkipVerify controls whether the server's certificate chain and host name is verified.
	// +optional
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`

	// CertificateAuthority is a PEM-encoded certificate to use to establish a connection with the LDAP server if the
	// connection uses TLS but that the certificate was signed by a custom Certificate Authority.
	// +optional
	CertificateAuthority string `json:"certificateAuthority,omitempty"`

	// BindDN is the domain name to bind to in order to authenticate to the LDAP server when running in search mode.
	// If empty, an anonymous bind will be done.
	// +optional
	BindDN string `json:"bindDn,omitempty"`

	// BindPasswordSecretName is the name of the Kubernetes Secret containing the password for the bind DN.
	// The secret must contain a key named 'password'.
	// +optional
	// +kubebuilder:validation:MaxLength=253
	BindPasswordSecretName string `json:"bindPasswordSecretName,omitempty"`

	// BaseDN is the base domain name that should be used for bind and search queries.
	BaseDN string `json:"baseDn"`

	// Attribute is the LDAP object attribute used to form a bind DN when sending bind queries.
	// The bind DN is formed as <Attribute>=<Username>,<BaseDN>.
	// +optional
	// +kubebuilder:default="cn"
	Attribute string `json:"attribute,omitempty"`

	// SearchFilter is used to filter LDAP search queries.
	// Example: (&(objectClass=inetOrgPerson)(gidNumber=500)(uid=%s))
	// %s can be used as a placeholder for the username.
	// +optional
	SearchFilter string `json:"searchFilter,omitempty"`
}

// PortalAuthLDAPConfig holds LDAP configuration for portal authentication.
type PortalAuthLDAPConfig struct {
	LDAPConnectionConfig `json:",inline"`

	// Groups configures group extraction.
	// +optional
	Groups *LDAPGroups `json:"groups,omitempty"`

	// Attributes configures LDAP attribute mappings for user attributes.
	// +optional
	Attributes *Attributes `json:"attributes"`

	// SyncedAttributes are the user attributes to synchronize with Hub platform.
	// +optional
	// +kubebuilder:validation:MaxItems=6
	// +kubebuilder:validation:items:Enum=groups;userId;firstname;lastname;email;company
	SyncedAttributes []string `json:"syncedAttributes,omitempty"`
}

// Attributes configures LDAP attribute mappings for user attributes.
type Attributes struct {
	// UserID is the LDAP attribute for user ID mapping.
	// +optional
	UserID string `json:"userId,omitempty"`

	// Firstname is the LDAP attribute for user first name.
	// +optional
	Firstname string `json:"firstname,omitempty"`

	// Lastname is the LDAP attribute for user last name.
	// +optional
	Lastname string `json:"lastname,omitempty"`

	// Email is the LDAP attribute for user email.
	// +optional
	Email string `json:"email,omitempty"`

	// Company is the LDAP attribute for user company.
	// +optional
	Company string `json:"company,omitempty"`
}

// LDAPGroups configures LDAP group extraction.
type LDAPGroups struct {
	// MemberOfAttribute is the LDAP attribute containing group memberships (e.g., "memberOf").
	// +kubebuilder:default="memberOf"
	MemberOfAttribute string `json:"memberOfAttribute,omitempty"`
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
