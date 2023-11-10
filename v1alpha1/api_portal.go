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

// APIPortal defines a portal that exposes APIs.
// +kubebuilder:printcolumn:name="URLs",type=string,JSONPath=`.status.urls`
// +kubebuilder:resource:scope=Cluster
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
	// +optional
	Title string `json:"title,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	APIGateway  string `json:"apiGateway"`
	// CustomDomains are the custom domains under which the portal will be exposed.
	// +optional
	CustomDomains []string `json:"customDomains,omitempty"`
	// +optional
	UI *UISpec `json:"ui,omitempty"`
}

// APIPortalStatus is the status of an APIPortal.
type APIPortalStatus struct {
	Version  string      `json:"version,omitempty"`
	SyncedAt metav1.Time `json:"syncedAt,omitempty"`

	// URLs are the URLs for accessing the APIPortal WebUI.
	URLs string `json:"urls"`

	// HubDomain is the hub generated domain of the APIPortal WebUI.
	// +optional
	HubDomain string `json:"hubDomain"`

	// CustomDomains are the custom domains for accessing the exposed APIPortal WebUI.
	// +optional
	CustomDomains []string `json:"customDomains,omitempty"`

	// Hash is a hash representing the APIPortal.
	Hash string `json:"hash,omitempty"`
}

// UISpec configures the UI customization.
type UISpec struct {
	// +optional
	Service *UIService `json:"service,omitempty"`
	// +optional
	LogoURL string `json:"logoUrl,omitempty"`
}

// UIService configures the service to expose on the edge.
type UIService struct {
	Name string `json:"name"`
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// port of the referenced service. A port name and port number
	// is required for an UIServiceBackendPort.
	Port UIServiceBackendPort `json:"port"`
}

// UIServiceBackendPort is the service port being referenced.
type UIServiceBackendPort struct {
	// name is the name of the port on the Service.
	// This must be an IANA_SVC_NAME (following RFC6335).
	// This is a mutually exclusive setting with "Number".
	// +optional
	Name string `json:"name"`

	// number is the numerical port number (e.g. 80) on the Service.
	// This is a mutually exclusive setting with "Name".
	// +optional
	Number int32 `json:"number"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIPortalList defines a list of APIPortals.
type APIPortalList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APIPortal `json:"items"`
}
