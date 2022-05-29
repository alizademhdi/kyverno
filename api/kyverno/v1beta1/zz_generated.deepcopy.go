//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/kyverno/kyverno/api/kyverno/v1"
	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionRequestInfoObject) DeepCopyInto(out *AdmissionRequestInfoObject) {
	*out = *in
	if in.AdmissionRequest != nil {
		in, out := &in.AdmissionRequest, &out.AdmissionRequest
		*out = new(admissionv1.AdmissionRequest)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionRequestInfoObject.
func (in *AdmissionRequestInfoObject) DeepCopy() *AdmissionRequestInfoObject {
	if in == nil {
		return nil
	}
	out := new(AdmissionRequestInfoObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestInfo) DeepCopyInto(out *RequestInfo) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterRoles != nil {
		in, out := &in.ClusterRoles, &out.ClusterRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.AdmissionUserInfo.DeepCopyInto(&out.AdmissionUserInfo)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestInfo.
func (in *RequestInfo) DeepCopy() *RequestInfo {
	if in == nil {
		return nil
	}
	out := new(RequestInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequest) DeepCopyInto(out *UpdateRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequest.
func (in *UpdateRequest) DeepCopy() *UpdateRequest {
	if in == nil {
		return nil
	}
	out := new(UpdateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpdateRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequestList) DeepCopyInto(out *UpdateRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UpdateRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequestList.
func (in *UpdateRequestList) DeepCopy() *UpdateRequestList {
	if in == nil {
		return nil
	}
	out := new(UpdateRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpdateRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequestSpec) DeepCopyInto(out *UpdateRequestSpec) {
	*out = *in
	out.Resource = in.Resource
	in.Context.DeepCopyInto(&out.Context)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequestSpec.
func (in *UpdateRequestSpec) DeepCopy() *UpdateRequestSpec {
	if in == nil {
		return nil
	}
	out := new(UpdateRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequestSpecContext) DeepCopyInto(out *UpdateRequestSpecContext) {
	*out = *in
	in.UserRequestInfo.DeepCopyInto(&out.UserRequestInfo)
	in.AdmissionRequestInfo.DeepCopyInto(&out.AdmissionRequestInfo)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequestSpecContext.
func (in *UpdateRequestSpecContext) DeepCopy() *UpdateRequestSpecContext {
	if in == nil {
		return nil
	}
	out := new(UpdateRequestSpecContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequestStatus) DeepCopyInto(out *UpdateRequestStatus) {
	*out = *in
	if in.GeneratedResources != nil {
		in, out := &in.GeneratedResources, &out.GeneratedResources
		*out = make([]v1.ResourceSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequestStatus.
func (in *UpdateRequestStatus) DeepCopy() *UpdateRequestStatus {
	if in == nil {
		return nil
	}
	out := new(UpdateRequestStatus)
	in.DeepCopyInto(out)
	return out
}
