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

// APIGateway defines a gateway that exposes APIs.
// +kubebuilder:printcolumn:name="URLs",type=string,JSONPath=`.status.urls`
// +kubebuilder:resource:scope=Cluster
type APIGateway struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APIGateway.
	Spec APIGatewaySpec `json:"spec,omitempty"`

	// The current status of this APIGateway.
	// +optional
	Status APIGatewayStatus `json:"status,omitempty"`
}

// APIGatewaySpec configures an APIGateway.
type APIGatewaySpec struct {
	// +optional
	// APIAccesses holds references to the APIAccess resources, each granting access to APIs that will be exposed
	// thought the gateway.
	APIAccesses []string `json:"apiAccesses,omitempty"`

	// CustomDomains are the custom domains under which the gateway will be exposed.
	// +optional
	// +kubebuilder:validation:MaxItems=20
	// +kubebuilder:validation:XValidation:message="duplicate domains",rule="self.all(x, self.exists_one(y, y == x))"
	CustomDomains []Domain `json:"customDomains,omitempty"`
}

// Domain is the domain name.
// +kubebuilder:validation:MaxLength=253
// +kubebuilder:validation:XValidation:message="custom domain must be a valid domain name",rule="self.matches(r\"\"\"([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]\"\"\")"
type Domain string

// APIGatewayStatus is the status of an APIGateway.
type APIGatewayStatus struct {
	Version  string      `json:"version,omitempty"`
	SyncedAt metav1.Time `json:"syncedAt,omitempty"`

	// URLs are the URLs for accessing the APIGateway.
	URLs string `json:"urls"`

	// HubDomain is the hub generated domain of the APIGateway.
	// +optional
	HubDomain string `json:"hubDomain"`

	// CustomDomains are the custom domains for accessing the exposed APIGateway.
	// +optional
	CustomDomains []string `json:"customDomains,omitempty"`

	// Hash is a hash representing the APIPortal.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIGatewayList defines a list of APIGateway.
type APIGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APIGateway `json:"items"`
}
