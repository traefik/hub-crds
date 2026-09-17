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

// HTTPClientConfig configures HTTP clients.
type HTTPClientConfig struct {
	// TLS configures TLS for the HTTP client.
	TLS *HTTPClientConfigTLS `json:"tls,omitempty"`
	// TimeoutSeconds configures the maximum amount of seconds to wait before giving up on requests.
	// +kubebuilder:default:=5
	TimeoutSeconds int `json:"timeoutSeconds,omitempty"`
	// MaxRetries defines the maximum number of retry attempts for failed requests.
	// +kubebuilder:default:=3
	MaxRetries int `json:"maxRetries,omitempty"`
}

// HTTPClientConfigTLS configures TLS for HTTP clients.
// +kubebuilder:validation:XValidation:message="ca and caSecretName are mutually exclusive",rule="[has(self.ca), has(self.caSecretName)].filter(x, x).size() <= 1"
type HTTPClientConfigTLS struct {
	// CA sets the CA bundle used to verify the server certificate.
	// Mutually exclusive with CASecretName.
	CA string `json:"ca,omitempty"`
	// CASecretName is the name of the Kubernetes Secret containing the CA bundle used to verify the server certificate.
	// The secret must contain a key named 'tls.ca'.
	// Mutually exclusive with CA.
	// +optional
	// +kubebuilder:validation:MaxLength=253
	CASecretName string `json:"caSecretName,omitempty"`
	// InsecureSkipVerify skips the server certificate validation.
	// For testing purposes only, do not use in production.
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}
