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
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIRateLimit defines how group of consumers are rate limited on a set of APIs.
// +kubebuilder:resource:scope=Cluster
type APIRateLimit struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APIRateLimit.
	// +kubebuilder:validation:XValidation:message="groups and anyGroups are mutually exclusive",rule="(has(self.anyGroups) && has(self.groups)) ? !(self.anyGroups && self.groups.size() > 0) : true"
	Spec APIRateLimitSpec `json:"spec,omitempty"`

	// The current status of this APIRateLimit.
	// +optional
	Status APIRateLimitStatus `json:"status,omitempty"`
}

// APIRateLimitSpec configures an APIRateLimit.
// The rate limiter is implemented using the Token Bucket algorithm, meaning, a virtual bucket is refilled at
// a constant rate defined by Limit/Period: https://en.wikipedia.org/wiki/Token_bucket
type APIRateLimitSpec struct {
	// Limit is the maximum number of token in the bucket.
	// +kubebuilder:validation:XValidation:message="must be a positive number",rule="self >= 0"
	Limit int `json:"limit"`

	// Period is frequency of bucket full refill.
	// +optional
	// +kubebuilder:validation:XValidation:message="must be between 1s and 1h",rule="self >= duration('1s') && self <= duration('1h')"
	Period *Period `json:"period,omitempty"`

	// Strategy defines how the bucket state will be synchronized between the different Traefik Hub agent instances.
	// It can be, either "local" or "distributed".
	// +optional
	// +kubebuilder:validation:Enum=local;distributed
	Strategy Strategy `json:"strategy,omitempty"`

	// Groups are the user groups that will be rate limited.
	// Multiple APIRateLimits can target the same set of user groups, the most restrictive one applies.
	// When a user belongs to multiple groups, the least restrictive APIRateLimit applies.
	// +optional
	Groups []string `json:"groups"`

	// AnyGroups states that all user groups will by default be rate limited with this configuration.
	// If an APIRateLimit explicitly target a group, the default rate limit will be ignored.
	// +optional
	AnyGroups bool `json:"anyGroups"`

	// APISelector selects the APIs that will be rate limited.
	// Multiple APIRateLimits can select the same set of APIs.
	// This field is optional and follows standard label selector semantics.
	// An empty APISelector matches any API.
	// +optional
	APISelector *metav1.LabelSelector `json:"apiSelector,omitempty"`

	// APIs defines a set of APIs that will be rate limited.
	// Multiple APIRateLimits can select the same APIs.
	// When combined with APISelector, this set of APIs is appended to the matching APIs.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:message="duplicated apis",rule="self.all(x, self.exists_one(y, x.name == y.name && (has(x.__namespace__) && x.__namespace__ != '' ? x.__namespace__ : 'default') == (has(y.__namespace__) && y.__namespace__ != '' ? y.__namespace__ : 'default')))"
	APIs []APIReference `json:"apis,omitempty"`
}

// Strategy defines how the rate limit buckets will be stored.
type Strategy string

// Supported rate limiting strategy.
const (
	StrategyLocal       Strategy = "local"
	StrategyDistributed Strategy = "distributed"
)

// APIRateLimitStatus is the status of an APIRateLimit.
type APIRateLimitStatus struct {
	Version  string      `json:"version,omitempty"`
	SyncedAt metav1.Time `json:"syncedAt,omitempty"`
	// Hash is a hash representing the APIRateLimit.
	Hash string `json:"hash,omitempty"`
}

// Period describes the time window on which a limit applies.
// +kubebuilder:validation:Type=string
// +kubebuilder:validation:Format=duration
type Period time.Duration

// NewPeriod creates a new Period.
func NewPeriod(d time.Duration) *Period {
	p := Period(d)

	return &p
}

// Seconds returns the period in seconds.
func (p *Period) Seconds() float64 {
	if p == nil {
		return 0
	}

	return time.Duration(*p).Seconds()
}

// IsZero checks whether the period is a zero-value Period.
func (p *Period) IsZero() bool {
	return p == nil || time.Duration(*p) == 0
}

// MarshalJSON marshals the Period.
func (p *Period) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte("null"), nil
	}

	return json.Marshal(toStringShortDuration(time.Duration(*p)))
}

// UnmarshalJSON unmarshals the buffer into a Period.
func (p *Period) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	if value, ok := v.(string); ok {
		duration, err := time.ParseDuration(value)
		if err != nil {
			return fmt.Errorf("parse: %w", err)
		}

		*p = Period(duration)

		return nil
	}

	return errors.New("invalid period")
}

func (p *Period) String() string {
	return toStringShortDuration(time.Duration(*p))
}

// toStringShortDuration stringifies the given duration in it's shorted form.
// Contrary to Go standard stringifier, here time.Minute will become "1m" instead of "1m0s".
func toStringShortDuration(duration time.Duration) string {
	short := duration.String()
	if strings.HasSuffix(short, "m0s") {
		short = short[:len(short)-2]
	}

	if strings.HasSuffix(short, "h0m") {
		short = short[:len(short)-2]
	}

	return short
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIRateLimitList defines a list of APIRateLimits.
type APIRateLimitList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []APIRateLimit `json:"items"`
}
