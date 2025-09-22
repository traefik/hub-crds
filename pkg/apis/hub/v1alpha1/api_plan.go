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

// APIPlan defines API Plan policy.
type APIPlan struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APIPlan.
	Spec APIPlanSpec `json:"spec,omitempty"`

	// The current status of this APIPlan.
	// +optional
	Status APIPlanStatus `json:"status,omitempty"`
}

// APIPlanSpec configures an APIPlan.
type APIPlanSpec struct {
	// Title is the human-readable name of the plan.
	Title string `json:"title"`

	// Description describes the plan.
	// +optional
	Description string `json:"description,omitempty"`

	// RateLimit defines the rate limit policy.
	// +optional
	RateLimit *RateLimit `json:"rateLimit,omitempty"`

	// Quota defines the quota policy.
	// +optional
	Quota *Quota `json:"quota,omitempty"`
}

// APIPlanStatus is the status of an APIPlan.
type APIPlanStatus struct {
	Version  string       `json:"version,omitempty"`
	SyncedAt *metav1.Time `json:"syncedAt,omitempty"`
	// Hash is a hash representing the APIPlan.
	Hash string `json:"hash,omitempty"`
}

type RateLimit struct {
	// Limit is the maximum number of token in the bucket.
	// +kubebuilder:validation:XValidation:message="must be a positive number",rule="self >= 0"
	Limit int `json:"limit"`

	// Period is the unit of time for the Limit.
	// +optional
	// +kubebuilder:validation:XValidation:message="must be between 1s and 1h",rule="self >= duration('1s') && self <= duration('1h')"
	Period *Period `json:"period,omitempty"`

	// Bucket defines the bucket strategy for the rate limit.
	// +optional
	// +kubebuilder:default="subscription"
	// +kubebuilder:validation:Enum=subscription;application-api;application
	Bucket Bucket `json:"bucket,omitempty"`
}

type Quota struct {
	// Limit is the maximum number of token in the bucket.
	// +kubebuilder:validation:XValidation:message="must be a positive number",rule="self >= 0"
	Limit int `json:"limit"`

	// Period is the unit of time for the Limit.
	// +optional
	// +kubebuilder:validation:XValidation:message="must be between 1s and 9999h",rule="self >= duration('1s') && self <= duration('9999h')"
	Period *Period `json:"period,omitempty"`

	// Bucket defines the bucket strategy for the quota.
	// +optional
	// +kubebuilder:default="subscription"
	// +kubebuilder:validation:Enum=subscription;application-api;application
	Bucket Bucket `json:"bucket,omitempty"`
}

// Bucket is a bucket strategy.
type Bucket string

const (
	// BucketSubscription shares the rate limit or quota across all APIs and applications
	// within the same subscription, providing a global limit for the entire subscription.
	BucketSubscription Bucket = "subscription"
	// BucketApplicationAPI creates separate rate limit or quota buckets for each unique
	// combination of application and API, allowing fine-grained control per app-API pair.
	BucketApplicationAPI Bucket = "application-api"
	// BucketApplication creates a single rate limit or quota bucket per application,
	// shared across all APIs that the application calls.
	BucketApplication Bucket = "application"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIPlanList defines a list of APIPlans.
type APIPlanList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APIPlan `json:"items"`
}
