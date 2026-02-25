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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

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
type UplinkHealthCheck struct {
	// +optional
	// Scheme replaces the server URL scheme for the health check endpoint.
	Scheme string `json:"scheme,omitempty"`

	// +optional
	// Mode defines the health check mode.
	// If defined to grpc, will use the gRPC health check protocol to probe the server.
	// Default: http
	Mode string `json:"mode,omitempty"`

	// +optional
	// Path defines the server URL path for the health check endpoint.
	Path string `json:"path,omitempty"`

	// +optional
	// Method defines the healthcheck method.
	Method string `json:"method,omitempty"`

	// +optional
	// Status defines the expected HTTP status code of the response to the health check request.
	Status int `json:"status,omitempty"`

	// +optional
	// Port defines the server URL port for the health check endpoint.
	Port int `json:"port,omitempty"`

	// +optional
	// Interval defines the frequency of the health check calls for healthy targets.
	// Default: 30s
	Interval *intstr.IntOrString `json:"interval,omitempty"`

	// +optional
	// UnhealthyInterval defines the frequency of the health check calls for unhealthy targets.
	// When UnhealthyInterval is not defined, it defaults to the Interval value.
	// Default: 30s
	UnhealthyInterval *intstr.IntOrString `json:"unhealthyInterval,omitempty"`

	// +optional
	// Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy.
	// Default: 5s
	Timeout *intstr.IntOrString `json:"timeout,omitempty"`

	// +optional
	// Hostname defines the value of hostname in the Host header of the health check request.
	Hostname string `json:"hostname,omitempty"`

	// +optional
	// FollowRedirects defines whether redirects should be followed during the health check calls.
	// Default: true
	FollowRedirects *bool `json:"followRedirects,omitempty"`

	// +optional
	// Headers defines custom headers to be sent to the health check endpoint.
	Headers map[string]string `json:"headers,omitempty"`
}

// UplinkPassiveHealthCheck mirrors Traefik's PassiveServerHealthCheck.
type UplinkPassiveHealthCheck struct {
	// +optional
	// FailureWindow defines the time window during which the failed attempts must occur for the server to be marked as unhealthy. It also defines for how long the server will be considered unhealthy.
	FailureWindow *intstr.IntOrString `json:"failureWindow,omitempty"`

	// +optional
	// MaxFailedAttempts is the number of consecutive failed attempts allowed within the failure window before marking the server as unhealthy.
	MaxFailedAttempts *int `json:"maxFailedAttempts,omitempty"`
}

// UplinkSticky mirrors Traefik's Sticky.
type UplinkSticky struct {
	// +optional
	Cookie *UplinkCookie `json:"cookie,omitempty"`
}

// UplinkCookie mirrors Traefik's Cookie.
// Same as Traefik type apart from Expires which is `json:"-"`.
type UplinkCookie struct {
	// +optional
	// Name defines the Cookie name.
	Name string `json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" export:"true"`

	// +optional
	// Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).
	Secure bool `json:"secure,omitempty" toml:"secure,omitempty" yaml:"secure,omitempty" export:"true"`

	// +optional
	// HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.
	HTTPOnly bool `json:"httpOnly,omitempty" toml:"httpOnly,omitempty" yaml:"httpOnly,omitempty" export:"true"`

	// +optional
	// SameSite defines the same site policy.
	// More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite
	// +kubebuilder:validation:Enum=none;lax;strict
	SameSite string `json:"sameSite,omitempty" toml:"sameSite,omitempty" yaml:"sameSite,omitempty" export:"true"`

	// +optional
	// MaxAge defines the number of seconds until the cookie expires.
	// When set to a negative number, the cookie expires immediately.
	// When set to zero, the cookie never expires.
	MaxAge int `json:"maxAge,omitempty" toml:"maxAge,omitempty" yaml:"maxAge,omitempty" export:"true"`

	// +optional
	// Path defines the path that must exist in the requested URL for the browser to send the Cookie header.
	// When not provided the cookie will be sent on every request to the domain.
	// More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#pathpath-value
	Path *string `json:"path,omitempty" toml:"path,omitempty" yaml:"path,omitempty" export:"true"`

	// +optional
	// Domain defines the host to which the cookie will be sent.
	// More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#domaindomain-value
	Domain string `json:"domain,omitempty" toml:"domain,omitempty" yaml:"domain,omitempty"`
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
