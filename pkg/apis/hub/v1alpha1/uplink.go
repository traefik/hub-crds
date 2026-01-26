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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Uplink is an inter-cluster service advertisement: a child cluster declares an Uplink to advertise
// to a parent cluster that it can handle a particular workload.
// +kubebuilder:subresource:status
type Uplink struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec UplinkSpec `json:"spec,omitempty"`

	// The current status of this Uplink.
	// +optional
	Status UplinkStatus `json:"status,omitempty"`
}

// UplinkSpec describes the Uplink.
type UplinkSpec struct {
	// Entrypoints references uplinkEntrypoints. When omitted, uses default uplinkEntrypoints.
	// +optional
	Entrypoints []string `json:"entrypoints,omitempty"`

	// Weight for WRR on the parent.
	// +optional
	// +kubebuilder:validation:XValidation:message="must be a positive number",rule="self >= 0"
	Weight *int `json:"weight,omitempty"`
}

// UplinkStatus is the status of the Uplink.
type UplinkStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UplinkList defines a list of Uplinks.
type UplinkList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Uplink `json:"items"`
}
