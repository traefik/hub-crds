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
	// ExposeName is the name of the service to expose.
	// By default it uses <namespace>-<name>.
	// +optional
	ExposeName string `json:"exposeName,omitempty"`

	// EntryPoints references uplinkEntryPoints. When omitted, uses default uplinkEntrypoints.
	// +optional
	EntryPoints []string `json:"entryPoints,omitempty"`

	// Weight for WRR on the parent.
	// +optional
	// +kubebuilder:validation:XValidation:message="must be a positive number",rule="self >= 0"
	Weight *int `json:"weight,omitempty"`

	// HealthCheck configures the active health check on the parent cluster for this uplink's load balancer.
	// +optional
	HealthCheck *UplinkHealthCheck `json:"healthcheck,omitempty"`

	// PassiveHealthCheck configures the passive health check on the parent cluster for this uplink's load balancer.
	// +optional
	PassiveHealthCheck *UplinkPassiveHealthCheck `json:"passiveHealthCheck,omitempty"`

	// Sticky configures cookie-based session affinity on the parent cluster for this uplink's services.
	// +optional
	Sticky *UplinkSticky `json:"sticky,omitempty"`
}

// UplinkHealthCheck mirrors Traefik's ServerHealthCheck.
// Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go
type UplinkHealthCheck struct {
	// +optional
	Scheme string `json:"scheme,omitempty"`
	// +optional
	Mode string `json:"mode,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	Method string `json:"method,omitempty"`
	// +optional
	Status int `json:"status,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	Interval *Period `json:"interval,omitempty"`
	// +optional
	UnhealthyInterval *Period `json:"unhealthyInterval,omitempty"`
	// +optional
	Timeout *Period `json:"timeout,omitempty"`
	// +optional
	Hostname string `json:"hostname,omitempty"`
	// +optional
	FollowRedirects *bool `json:"followRedirects,omitempty"`
	// +optional
	Headers map[string]string `json:"headers,omitempty"`
}

// UplinkPassiveHealthCheck mirrors Traefik's PassiveServerHealthCheck.
// Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go
type UplinkPassiveHealthCheck struct {
	// +optional
	FailureWindow *Period `json:"failureWindow,omitempty"`
	// +optional
	MaxFailedAttempts int `json:"maxFailedAttempts,omitempty"`
}

// UplinkSticky mirrors Traefik's Sticky.
// Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go
type UplinkSticky struct {
	// +optional
	Cookie *UplinkCookie `json:"cookie,omitempty"`
}

// UplinkCookie mirrors Traefik's Cookie.
// Based-On: https://github.com/traefik/traefik/blob/master/pkg/config/dynamic/http_config.go
// Same as Traefik type apart from Expires which is `json:"-"`.
type UplinkCookie struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Secure bool `json:"secure,omitempty"`
	// +optional
	HTTPOnly bool `json:"httpOnly,omitempty"`
	// +kubebuilder:validation:Enum=none;lax;strict
	// +optional
	SameSite string `json:"sameSite,omitempty"`
	// +optional
	MaxAge int `json:"maxAge,omitempty"`
	// +optional
	Path *string `json:"path,omitempty"`
	// +optional
	Domain string `json:"domain,omitempty"`
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
