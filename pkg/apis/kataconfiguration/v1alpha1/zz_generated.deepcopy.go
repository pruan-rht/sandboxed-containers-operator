// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailedNode) DeepCopyInto(out *FailedNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailedNode.
func (in *FailedNode) DeepCopy() *FailedNode {
	if in == nil {
		return nil
	}
	out := new(FailedNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KataConfig) DeepCopyInto(out *KataConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KataConfig.
func (in *KataConfig) DeepCopy() *KataConfig {
	if in == nil {
		return nil
	}
	out := new(KataConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KataConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KataConfigList) DeepCopyInto(out *KataConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KataConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KataConfigList.
func (in *KataConfigList) DeepCopy() *KataConfigList {
	if in == nil {
		return nil
	}
	out := new(KataConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KataConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KataConfigSpec) DeepCopyInto(out *KataConfigSpec) {
	*out = *in
	if in.KataConfigPoolSelector != nil {
		in, out := &in.KataConfigPoolSelector, &out.KataConfigPoolSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	out.Config = in.Config
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KataConfigSpec.
func (in *KataConfigSpec) DeepCopy() *KataConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KataConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KataConfigStatus) DeepCopyInto(out *KataConfigStatus) {
	*out = *in
	if in.FailedNodes != nil {
		in, out := &in.FailedNodes, &out.FailedNodes
		*out = make([]FailedNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KataConfigStatus.
func (in *KataConfigStatus) DeepCopy() *KataConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KataConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KataInstallConfig) DeepCopyInto(out *KataInstallConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KataInstallConfig.
func (in *KataInstallConfig) DeepCopy() *KataInstallConfig {
	if in == nil {
		return nil
	}
	out := new(KataInstallConfig)
	in.DeepCopyInto(out)
	return out
}