//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
The GNU AFFERO GENERAL PUBLIC LICENSE

Copyright (c) 2020-2024 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *API) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIAccess) DeepCopyInto(out *APIAccess) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIAccess.
func (in *APIAccess) DeepCopy() *APIAccess {
	if in == nil {
		return nil
	}
	out := new(APIAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIAccess) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIAccessList) DeepCopyInto(out *APIAccessList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIAccess, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIAccessList.
func (in *APIAccessList) DeepCopy() *APIAccessList {
	if in == nil {
		return nil
	}
	out := new(APIAccessList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIAccessList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIAccessSpec) DeepCopyInto(out *APIAccessSpec) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APISelector != nil {
		in, out := &in.APISelector, &out.APISelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.APIs != nil {
		in, out := &in.APIs, &out.APIs
		*out = make([]APIReference, len(*in))
		copy(*out, *in)
	}
	if in.OperationFilter != nil {
		in, out := &in.OperationFilter, &out.OperationFilter
		*out = new(OperationFilter)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIAccessSpec.
func (in *APIAccessSpec) DeepCopy() *APIAccessSpec {
	if in == nil {
		return nil
	}
	out := new(APIAccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIAccessStatus) DeepCopyInto(out *APIAccessStatus) {
	*out = *in
	if in.SyncedAt != nil {
		in, out := &in.SyncedAt, &out.SyncedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIAccessStatus.
func (in *APIAccessStatus) DeepCopy() *APIAccessStatus {
	if in == nil {
		return nil
	}
	out := new(APIAccessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIList) DeepCopyInto(out *APIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]API, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIList.
func (in *APIList) DeepCopy() *APIList {
	if in == nil {
		return nil
	}
	out := new(APIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPortal) DeepCopyInto(out *APIPortal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPortal.
func (in *APIPortal) DeepCopy() *APIPortal {
	if in == nil {
		return nil
	}
	out := new(APIPortal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPortal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPortalList) DeepCopyInto(out *APIPortalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIPortal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPortalList.
func (in *APIPortalList) DeepCopy() *APIPortalList {
	if in == nil {
		return nil
	}
	out := new(APIPortalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPortalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPortalSpec) DeepCopyInto(out *APIPortalSpec) {
	*out = *in
	if in.TrustedURLs != nil {
		in, out := &in.TrustedURLs, &out.TrustedURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UI != nil {
		in, out := &in.UI, &out.UI
		*out = new(UISpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPortalSpec.
func (in *APIPortalSpec) DeepCopy() *APIPortalSpec {
	if in == nil {
		return nil
	}
	out := new(APIPortalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPortalStatus) DeepCopyInto(out *APIPortalStatus) {
	*out = *in
	if in.SyncedAt != nil {
		in, out := &in.SyncedAt, &out.SyncedAt
		*out = (*in).DeepCopy()
	}
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = new(OIDCConfigStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPortalStatus.
func (in *APIPortalStatus) DeepCopy() *APIPortalStatus {
	if in == nil {
		return nil
	}
	out := new(APIPortalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRateLimit) DeepCopyInto(out *APIRateLimit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRateLimit.
func (in *APIRateLimit) DeepCopy() *APIRateLimit {
	if in == nil {
		return nil
	}
	out := new(APIRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIRateLimit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRateLimitList) DeepCopyInto(out *APIRateLimitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIRateLimit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRateLimitList.
func (in *APIRateLimitList) DeepCopy() *APIRateLimitList {
	if in == nil {
		return nil
	}
	out := new(APIRateLimitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIRateLimitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRateLimitSpec) DeepCopyInto(out *APIRateLimitSpec) {
	*out = *in
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(Period)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APISelector != nil {
		in, out := &in.APISelector, &out.APISelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.APIs != nil {
		in, out := &in.APIs, &out.APIs
		*out = make([]APIReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRateLimitSpec.
func (in *APIRateLimitSpec) DeepCopy() *APIRateLimitSpec {
	if in == nil {
		return nil
	}
	out := new(APIRateLimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRateLimitStatus) DeepCopyInto(out *APIRateLimitStatus) {
	*out = *in
	if in.SyncedAt != nil {
		in, out := &in.SyncedAt, &out.SyncedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRateLimitStatus.
func (in *APIRateLimitStatus) DeepCopy() *APIRateLimitStatus {
	if in == nil {
		return nil
	}
	out := new(APIRateLimitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIReference) DeepCopyInto(out *APIReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIReference.
func (in *APIReference) DeepCopy() *APIReference {
	if in == nil {
		return nil
	}
	out := new(APIReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	if in.OpenAPISpec != nil {
		in, out := &in.OpenAPISpec, &out.OpenAPISpec
		*out = new(OpenAPISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]APIVersionRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIStatus) DeepCopyInto(out *APIStatus) {
	*out = *in
	if in.SyncedAt != nil {
		in, out := &in.SyncedAt, &out.SyncedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIStatus.
func (in *APIStatus) DeepCopy() *APIStatus {
	if in == nil {
		return nil
	}
	out := new(APIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIVersion) DeepCopyInto(out *APIVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIVersion.
func (in *APIVersion) DeepCopy() *APIVersion {
	if in == nil {
		return nil
	}
	out := new(APIVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIVersionList) DeepCopyInto(out *APIVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIVersionList.
func (in *APIVersionList) DeepCopy() *APIVersionList {
	if in == nil {
		return nil
	}
	out := new(APIVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIVersionRef) DeepCopyInto(out *APIVersionRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIVersionRef.
func (in *APIVersionRef) DeepCopy() *APIVersionRef {
	if in == nil {
		return nil
	}
	out := new(APIVersionRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIVersionSpec) DeepCopyInto(out *APIVersionSpec) {
	*out = *in
	if in.OpenAPISpec != nil {
		in, out := &in.OpenAPISpec, &out.OpenAPISpec
		*out = new(OpenAPISpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIVersionSpec.
func (in *APIVersionSpec) DeepCopy() *APIVersionSpec {
	if in == nil {
		return nil
	}
	out := new(APIVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIVersionStatus) DeepCopyInto(out *APIVersionStatus) {
	*out = *in
	if in.SyncedAt != nil {
		in, out := &in.SyncedAt, &out.SyncedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIVersionStatus.
func (in *APIVersionStatus) DeepCopy() *APIVersionStatus {
	if in == nil {
		return nil
	}
	out := new(APIVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlOAuthIntro) DeepCopyInto(out *AccessControlOAuthIntro) {
	*out = *in
	in.ClientConfig.DeepCopyInto(&out.ClientConfig)
	out.TokenSource = in.TokenSource
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlOAuthIntro.
func (in *AccessControlOAuthIntro) DeepCopy() *AccessControlOAuthIntro {
	if in == nil {
		return nil
	}
	out := new(AccessControlOAuthIntro)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlOAuthIntroClientConfig) DeepCopyInto(out *AccessControlOAuthIntroClientConfig) {
	*out = *in
	in.HTTPClientConfig.DeepCopyInto(&out.HTTPClientConfig)
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlOAuthIntroClientConfig.
func (in *AccessControlOAuthIntroClientConfig) DeepCopy() *AccessControlOAuthIntroClientConfig {
	if in == nil {
		return nil
	}
	out := new(AccessControlOAuthIntroClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicy) DeepCopyInto(out *AccessControlPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicy.
func (in *AccessControlPolicy) DeepCopy() *AccessControlPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessControlPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyAPIKey) DeepCopyInto(out *AccessControlPolicyAPIKey) {
	*out = *in
	out.KeySource = in.KeySource
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]AccessControlPolicyAPIKeyKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyAPIKey.
func (in *AccessControlPolicyAPIKey) DeepCopy() *AccessControlPolicyAPIKey {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyAPIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyAPIKeyKey) DeepCopyInto(out *AccessControlPolicyAPIKeyKey) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyAPIKeyKey.
func (in *AccessControlPolicyAPIKeyKey) DeepCopy() *AccessControlPolicyAPIKeyKey {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyAPIKeyKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyBasicAuth) DeepCopyInto(out *AccessControlPolicyBasicAuth) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyBasicAuth.
func (in *AccessControlPolicyBasicAuth) DeepCopy() *AccessControlPolicyBasicAuth {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyBasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyJWT) DeepCopyInto(out *AccessControlPolicyJWT) {
	*out = *in
	if in.StripAuthorizationHeader != nil {
		in, out := &in.StripAuthorizationHeader, &out.StripAuthorizationHeader
		*out = new(bool)
		**out = **in
	}
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyJWT.
func (in *AccessControlPolicyJWT) DeepCopy() *AccessControlPolicyJWT {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyJWT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyList) DeepCopyInto(out *AccessControlPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessControlPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyList.
func (in *AccessControlPolicyList) DeepCopy() *AccessControlPolicyList {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessControlPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyOIDC) DeepCopyInto(out *AccessControlPolicyOIDC) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(corev1.SecretReference)
		**out = **in
	}
	if in.DisableAuthRedirectionPaths != nil {
		in, out := &in.DisableAuthRedirectionPaths, &out.DisableAuthRedirectionPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AuthParams != nil {
		in, out := &in.AuthParams, &out.AuthParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StateCookie != nil {
		in, out := &in.StateCookie, &out.StateCookie
		*out = new(StateCookie)
		**out = **in
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = new(Session)
		(*in).DeepCopyInto(*out)
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyOIDC.
func (in *AccessControlPolicyOIDC) DeepCopy() *AccessControlPolicyOIDC {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyOIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyOIDCGoogle) DeepCopyInto(out *AccessControlPolicyOIDCGoogle) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(corev1.SecretReference)
		**out = **in
	}
	if in.AuthParams != nil {
		in, out := &in.AuthParams, &out.AuthParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StateCookie != nil {
		in, out := &in.StateCookie, &out.StateCookie
		*out = new(StateCookie)
		**out = **in
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = new(Session)
		(*in).DeepCopyInto(*out)
	}
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Emails != nil {
		in, out := &in.Emails, &out.Emails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyOIDCGoogle.
func (in *AccessControlPolicyOIDCGoogle) DeepCopy() *AccessControlPolicyOIDCGoogle {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyOIDCGoogle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicySpec) DeepCopyInto(out *AccessControlPolicySpec) {
	*out = *in
	if in.JWT != nil {
		in, out := &in.JWT, &out.JWT
		*out = new(AccessControlPolicyJWT)
		(*in).DeepCopyInto(*out)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(AccessControlPolicyBasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(AccessControlPolicyAPIKey)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = new(AccessControlPolicyOIDC)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDCGoogle != nil {
		in, out := &in.OIDCGoogle, &out.OIDCGoogle
		*out = new(AccessControlPolicyOIDCGoogle)
		(*in).DeepCopyInto(*out)
	}
	if in.OAuthIntro != nil {
		in, out := &in.OAuthIntro, &out.OAuthIntro
		*out = new(AccessControlOAuthIntro)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicySpec.
func (in *AccessControlPolicySpec) DeepCopy() *AccessControlPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyStatus) DeepCopyInto(out *AccessControlPolicyStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyStatus.
func (in *AccessControlPolicyStatus) DeepCopy() *AccessControlPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPClientConfig) DeepCopyInto(out *HTTPClientConfig) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(HTTPClientConfigTLS)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPClientConfig.
func (in *HTTPClientConfig) DeepCopy() *HTTPClientConfig {
	if in == nil {
		return nil
	}
	out := new(HTTPClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPClientConfigTLS) DeepCopyInto(out *HTTPClientConfigTLS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPClientConfigTLS.
func (in *HTTPClientConfigTLS) DeepCopy() *HTTPClientConfigTLS {
	if in == nil {
		return nil
	}
	out := new(HTTPClientConfigTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCConfigStatus) DeepCopyInto(out *OIDCConfigStatus) {
	*out = *in
	if in.SyncedAttributes != nil {
		in, out := &in.SyncedAttributes, &out.SyncedAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCConfigStatus.
func (in *OIDCConfigStatus) DeepCopy() *OIDCConfigStatus {
	if in == nil {
		return nil
	}
	out := new(OIDCConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenAPISpec) DeepCopyInto(out *OpenAPISpec) {
	*out = *in
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(Override)
		(*in).DeepCopyInto(*out)
	}
	if in.OperationSets != nil {
		in, out := &in.OperationSets, &out.OperationSets
		*out = make([]OperationSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenAPISpec.
func (in *OpenAPISpec) DeepCopy() *OpenAPISpec {
	if in == nil {
		return nil
	}
	out := new(OpenAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationFilter) DeepCopyInto(out *OperationFilter) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationFilter.
func (in *OperationFilter) DeepCopy() *OperationFilter {
	if in == nil {
		return nil
	}
	out := new(OperationFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationMatcher) DeepCopyInto(out *OperationMatcher) {
	*out = *in
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationMatcher.
func (in *OperationMatcher) DeepCopy() *OperationMatcher {
	if in == nil {
		return nil
	}
	out := new(OperationMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationSet) DeepCopyInto(out *OperationSet) {
	*out = *in
	if in.Matchers != nil {
		in, out := &in.Matchers, &out.Matchers
		*out = make([]OperationMatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationSet.
func (in *OperationSet) DeepCopy() *OperationSet {
	if in == nil {
		return nil
	}
	out := new(OperationSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Override) DeepCopyInto(out *Override) {
	*out = *in
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]Server, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Override.
func (in *Override) DeepCopy() *Override {
	if in == nil {
		return nil
	}
	out := new(Override)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Session) DeepCopyInto(out *Session) {
	*out = *in
	if in.Refresh != nil {
		in, out := &in.Refresh, &out.Refresh
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Session.
func (in *Session) DeepCopy() *Session {
	if in == nil {
		return nil
	}
	out := new(Session)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateCookie) DeepCopyInto(out *StateCookie) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateCookie.
func (in *StateCookie) DeepCopy() *StateCookie {
	if in == nil {
		return nil
	}
	out := new(StateCookie)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSource) DeepCopyInto(out *TokenSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSource.
func (in *TokenSource) DeepCopy() *TokenSource {
	if in == nil {
		return nil
	}
	out := new(TokenSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UISpec) DeepCopyInto(out *UISpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UISpec.
func (in *UISpec) DeepCopy() *UISpec {
	if in == nil {
		return nil
	}
	out := new(UISpec)
	in.DeepCopyInto(out)
	return out
}
