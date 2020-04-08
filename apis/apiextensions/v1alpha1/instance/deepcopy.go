package instance

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	corev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// NOTE(muvaf): controller-gen does not generate DeepCopy functions for structs
// with fields that have interface{} type. So, we manually write these functions
// to make it compatible with runtime.Object interface.

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraInstance) DeepCopyInto(out *InfraInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraInstance.
func (in *InfraInstance) DeepCopy() *InfraInstance {
	if in == nil {
		return nil
	}
	out := new(InfraInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfraInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraInstanceSpec) DeepCopyInto(out *InfraInstanceSpec) {
	*out = *in
	in.ResourceInstanceCommonSpec.DeepCopyInto(&out.ResourceInstanceCommonSpec)
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraInstanceSpec.
func (in *InfraInstanceSpec) DeepCopy() *InfraInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InfraInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInstanceCommonSpec) DeepCopyInto(out *ResourceInstanceCommonSpec) {
	*out = *in
	if in.WriteConnectionSecretToReference != nil {
		in, out := &in.WriteConnectionSecretToReference, &out.WriteConnectionSecretToReference
		*out = new(corev1alpha1.SecretReference)
		**out = **in
	}
	if in.CompositionSelector != nil {
		in, out := &in.CompositionSelector, &out.CompositionSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.CompositionReference != nil {
		in, out := &in.CompositionReference, &out.CompositionReference
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.ResourceReferences != nil {
		in, out := &in.ResourceReferences, &out.ResourceReferences
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInstanceCommonSpec.
func (in *ResourceInstanceCommonSpec) DeepCopy() *ResourceInstanceCommonSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceInstanceCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraInstanceList) DeepCopyInto(out *InfraInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InfraInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureDefinitionList.
func (in *InfraInstanceList) DeepCopy() *InfraInstanceList {
	if in == nil {
		return nil
	}
	out := new(InfraInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfraInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
