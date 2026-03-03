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

// ContentItem defines additional documentation for given resource.
// +kubebuilder:subresource:status
type ContentItem struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Defines the documentation to attach to the referenced resource.
	Spec ContentItemSpec `json:"spec,omitempty"`

	// The current status of this ContentItem.
	// +optional
	Status ContentItemStatus `json:"status,omitempty"`
}

// ContentItemSpec configures a ContentItem.
// +kubebuilder:validation:XValidation:message="exactly one of content or link must be specified",rule="[has(self.content), has(self.link)].filter(x, x).size() == 1"
type ContentItemSpec struct {
	// Title is the public-facing name of the ContentItem.
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:MinLength=1
	Title string `json:"title"`
	// Order defines the order of the content in the UI.
	// +kubebuilder:validation:Minimum=0
	Order int32 `json:"order"`
	// ParentRef is the reference to the resource that this content belongs to.
	ParentRef ContentItemParentRef `json:"parentRef"`
	// Link is the link to the content.
	// +optional
	Link *LinkDetails `json:"link,omitempty"`
	// Content is the valid markdown content.
	// +optional
	// +kubebuilder:validation:MaxLength=1500000
	Content string `json:"content,omitempty"`
}

// ContentItemParentRef references the resource to which ContentItem belongs.
type ContentItemParentRef struct {
	// Kind is the kind of the resource that this content belongs to.
	// +kubebuilder:validation:Enum=APIPortal;API;APIBundle
	Kind string `json:"kind"`
	// Name is the name of the resource that this content belongs to.
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
}

// LinkDetails accepts URL for configuring Content Data.
type LinkDetails struct {
	// Href is the public URL of the content.
	// +kubebuilder:validation:XValidation:message="must be a valid URL",rule="isURL(self)"
	Href string `json:"href"`
}

// ContentItemStatus is the status of a ContentItem.
type ContentItemStatus struct {
	Version  string       `json:"version,omitempty"`
	SyncedAt *metav1.Time `json:"syncedAt,omitempty"`

	// Hash is a hash representing the ContentItem.
	Hash string `json:"hash,omitempty"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContentItemList defines a list of ContentItem.
type ContentItemList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ContentItem `json:"items"`
}
