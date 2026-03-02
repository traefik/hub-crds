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
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Content defines additional documentation for given resource.
// +kubebuilder:subresource:status
type Content struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this Content.
	Spec ContentSpec `json:"spec,omitempty"`

	// The current status of this Content.
	// +optional
	Status ContentStatus `json:"status,omitempty"`
}

// ContentSpec configures a Content.
// +kubebuilder:validation:XValidation:message="exactly one of content or link must be specified",rule="[has(self.content), has(self.link)].filter(x, x).size() == 1"
type ContentSpec struct {
	// Title is the public-facing name of the Content.
	// +kubebuilder:validation:MaxLength=253
	Title string `json:"title"`
	// Order defines the order of the content in the UI.
	Order uint32 `json:"order"`
	// ParentRef is the reference to the resource that this content belongs to.
	ParentRef ContentParentRef `json:"parentRef"`
	// Link is the link to the content.
	// +optional
	Link *LinkDetails `json:"link,omitempty"`
	// Content is the valid markdown content.
	// +optional
	// +kubebuilder:validation:MaxLength=1500000
	Content string `json:"content,omitempty"`
}

type ContentParentRef struct {
	// Kind is the kind of the resource that this content belongs to.
	// +kubebuilder:validation:MaxLength=253
	Kind string `json:"kind"`
	// Name is the name of the resource that this content belongs to.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
}

type LinkDetails struct {
	// Href is the public URL of the content.
	// +kubebuilder:validation:XValidation:message="must be a valid URL",rule="isURL(self)"
	Href string `json:"href"`
}

// ContentStatus is the status of a Content.
type ContentStatus struct {
	Version  string       `json:"version,omitempty"`
	SyncedAt *metav1.Time `json:"syncedAt,omitempty"`

	// Hash is a hash representing the Content.
	Hash string `json:"hash,omitempty"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContentList defines a list of Content.
type ContentList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Content `json:"items"`
}
