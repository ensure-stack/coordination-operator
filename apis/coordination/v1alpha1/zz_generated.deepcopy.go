// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Adoption) DeepCopyInto(out *Adoption) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Adoption.
func (in *Adoption) DeepCopy() *Adoption {
	if in == nil {
		return nil
	}
	out := new(Adoption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Adoption) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdoptionList) DeepCopyInto(out *AdoptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Adoption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdoptionList.
func (in *AdoptionList) DeepCopy() *AdoptionList {
	if in == nil {
		return nil
	}
	out := new(AdoptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdoptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdoptionSpec) DeepCopyInto(out *AdoptionSpec) {
	*out = *in
	in.Strategy.DeepCopyInto(&out.Strategy)
	out.TargetAPI = in.TargetAPI
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdoptionSpec.
func (in *AdoptionSpec) DeepCopy() *AdoptionSpec {
	if in == nil {
		return nil
	}
	out := new(AdoptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdoptionStatus) DeepCopyInto(out *AdoptionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdoptionStatus.
func (in *AdoptionStatus) DeepCopy() *AdoptionStatus {
	if in == nil {
		return nil
	}
	out := new(AdoptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdoptionStrategy) DeepCopyInto(out *AdoptionStrategy) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(AdoptionStrategyStaticSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdoptionStrategy.
func (in *AdoptionStrategy) DeepCopy() *AdoptionStrategy {
	if in == nil {
		return nil
	}
	out := new(AdoptionStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdoptionStrategyStaticSpec) DeepCopyInto(out *AdoptionStrategyStaticSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdoptionStrategyStaticSpec.
func (in *AdoptionStrategyStaticSpec) DeepCopy() *AdoptionStrategyStaticSpec {
	if in == nil {
		return nil
	}
	out := new(AdoptionStrategyStaticSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdoption) DeepCopyInto(out *ClusterAdoption) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdoption.
func (in *ClusterAdoption) DeepCopy() *ClusterAdoption {
	if in == nil {
		return nil
	}
	out := new(ClusterAdoption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAdoption) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdoptionList) DeepCopyInto(out *ClusterAdoptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAdoption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdoptionList.
func (in *ClusterAdoptionList) DeepCopy() *ClusterAdoptionList {
	if in == nil {
		return nil
	}
	out := new(ClusterAdoptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAdoptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdoptionSpec) DeepCopyInto(out *ClusterAdoptionSpec) {
	*out = *in
	in.Strategy.DeepCopyInto(&out.Strategy)
	out.TargetAPI = in.TargetAPI
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdoptionSpec.
func (in *ClusterAdoptionSpec) DeepCopy() *ClusterAdoptionSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAdoptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdoptionStatus) DeepCopyInto(out *ClusterAdoptionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdoptionStatus.
func (in *ClusterAdoptionStatus) DeepCopy() *ClusterAdoptionStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterAdoptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdoptionStrategy) DeepCopyInto(out *ClusterAdoptionStrategy) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(ClusterAdoptionStrategyStaticSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdoptionStrategy.
func (in *ClusterAdoptionStrategy) DeepCopy() *ClusterAdoptionStrategy {
	if in == nil {
		return nil
	}
	out := new(ClusterAdoptionStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdoptionStrategyStaticSpec) DeepCopyInto(out *ClusterAdoptionStrategyStaticSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdoptionStrategyStaticSpec.
func (in *ClusterAdoptionStrategyStaticSpec) DeepCopy() *ClusterAdoptionStrategyStaticSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAdoptionStrategyStaticSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHandover) DeepCopyInto(out *ClusterHandover) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHandover.
func (in *ClusterHandover) DeepCopy() *ClusterHandover {
	if in == nil {
		return nil
	}
	out := new(ClusterHandover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterHandover) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHandoverList) DeepCopyInto(out *ClusterHandoverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterHandover, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHandoverList.
func (in *ClusterHandoverList) DeepCopy() *ClusterHandoverList {
	if in == nil {
		return nil
	}
	out := new(ClusterHandoverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterHandoverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHandoverSpec) DeepCopyInto(out *ClusterHandoverSpec) {
	*out = *in
	in.Strategy.DeepCopyInto(&out.Strategy)
	out.TargetAPI = in.TargetAPI
	if in.Probes != nil {
		in, out := &in.Probes, &out.Probes
		*out = make([]Probe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHandoverSpec.
func (in *ClusterHandoverSpec) DeepCopy() *ClusterHandoverSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterHandoverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHandoverStatus) DeepCopyInto(out *ClusterHandoverStatus) {
	*out = *in
	if in.Processing != nil {
		in, out := &in.Processing, &out.Processing
		*out = make([]HandoverRef, len(*in))
		copy(*out, *in)
	}
	out.Stats = in.Stats
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHandoverStatus.
func (in *ClusterHandoverStatus) DeepCopy() *ClusterHandoverStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterHandoverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Handover) DeepCopyInto(out *Handover) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Handover.
func (in *Handover) DeepCopy() *Handover {
	if in == nil {
		return nil
	}
	out := new(Handover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Handover) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandoverList) DeepCopyInto(out *HandoverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Handover, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandoverList.
func (in *HandoverList) DeepCopy() *HandoverList {
	if in == nil {
		return nil
	}
	out := new(HandoverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HandoverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandoverRef) DeepCopyInto(out *HandoverRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandoverRef.
func (in *HandoverRef) DeepCopy() *HandoverRef {
	if in == nil {
		return nil
	}
	out := new(HandoverRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandoverSpec) DeepCopyInto(out *HandoverSpec) {
	*out = *in
	in.Strategy.DeepCopyInto(&out.Strategy)
	out.TargetAPI = in.TargetAPI
	if in.Probes != nil {
		in, out := &in.Probes, &out.Probes
		*out = make([]Probe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandoverSpec.
func (in *HandoverSpec) DeepCopy() *HandoverSpec {
	if in == nil {
		return nil
	}
	out := new(HandoverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandoverStatus) DeepCopyInto(out *HandoverStatus) {
	*out = *in
	if in.Processing != nil {
		in, out := &in.Processing, &out.Processing
		*out = make([]HandoverRef, len(*in))
		copy(*out, *in)
	}
	out.Stats = in.Stats
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandoverStatus.
func (in *HandoverStatus) DeepCopy() *HandoverStatus {
	if in == nil {
		return nil
	}
	out := new(HandoverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandoverStatusStats) DeepCopyInto(out *HandoverStatusStats) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandoverStatusStats.
func (in *HandoverStatusStats) DeepCopy() *HandoverStatusStats {
	if in == nil {
		return nil
	}
	out := new(HandoverStatusStats)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandoverStrategy) DeepCopyInto(out *HandoverStrategy) {
	*out = *in
	if in.Relabel != nil {
		in, out := &in.Relabel, &out.Relabel
		*out = new(HandoverStrategyRelabelSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandoverStrategy.
func (in *HandoverStrategy) DeepCopy() *HandoverStrategy {
	if in == nil {
		return nil
	}
	out := new(HandoverStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandoverStrategyRelabelSpec) DeepCopyInto(out *HandoverStrategyRelabelSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandoverStrategyRelabelSpec.
func (in *HandoverStrategyRelabelSpec) DeepCopy() *HandoverStrategyRelabelSpec {
	if in == nil {
		return nil
	}
	out := new(HandoverStrategyRelabelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probe) DeepCopyInto(out *Probe) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ProbeConditionSpec)
		**out = **in
	}
	if in.FieldsEqual != nil {
		in, out := &in.FieldsEqual, &out.FieldsEqual
		*out = new(ProbeFieldsEqualSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probe.
func (in *Probe) DeepCopy() *Probe {
	if in == nil {
		return nil
	}
	out := new(Probe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeConditionSpec) DeepCopyInto(out *ProbeConditionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeConditionSpec.
func (in *ProbeConditionSpec) DeepCopy() *ProbeConditionSpec {
	if in == nil {
		return nil
	}
	out := new(ProbeConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeFieldsEqualSpec) DeepCopyInto(out *ProbeFieldsEqualSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeFieldsEqualSpec.
func (in *ProbeFieldsEqualSpec) DeepCopy() *ProbeFieldsEqualSpec {
	if in == nil {
		return nil
	}
	out := new(ProbeFieldsEqualSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetAPI) DeepCopyInto(out *TargetAPI) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetAPI.
func (in *TargetAPI) DeepCopy() *TargetAPI {
	if in == nil {
		return nil
	}
	out := new(TargetAPI)
	in.DeepCopyInto(out)
	return out
}