//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apisv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Breakout) DeepCopyInto(out *Breakout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Breakout.
func (in *Breakout) DeepCopy() *Breakout {
	if in == nil {
		return nil
	}
	out := new(Breakout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Breakout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreakoutHTTPRoutingRuleMatchesSpec) DeepCopyInto(out *BreakoutHTTPRoutingRuleMatchesSpec) {
	*out = *in
	if in.Matches != nil {
		in, out := &in.Matches, &out.Matches
		*out = make([]apisv1beta1.HTTPRouteMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreakoutHTTPRoutingRuleMatchesSpec.
func (in *BreakoutHTTPRoutingRuleMatchesSpec) DeepCopy() *BreakoutHTTPRoutingRuleMatchesSpec {
	if in == nil {
		return nil
	}
	out := new(BreakoutHTTPRoutingRuleMatchesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreakoutHTTPRoutingRuleSpec) DeepCopyInto(out *BreakoutHTTPRoutingRuleSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]BreakoutHTTPRoutingRuleMatchesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreakoutHTTPRoutingRuleSpec.
func (in *BreakoutHTTPRoutingRuleSpec) DeepCopy() *BreakoutHTTPRoutingRuleSpec {
	if in == nil {
		return nil
	}
	out := new(BreakoutHTTPRoutingRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreakoutList) DeepCopyInto(out *BreakoutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Breakout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreakoutList.
func (in *BreakoutList) DeepCopy() *BreakoutList {
	if in == nil {
		return nil
	}
	out := new(BreakoutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BreakoutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreakoutResourceInstance) DeepCopyInto(out *BreakoutResourceInstance) {
	*out = *in
	if in.OwnerReferences != nil {
		in, out := &in.OwnerReferences, &out.OwnerReferences
		*out = make([]v1.OwnerReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreakoutResourceInstance.
func (in *BreakoutResourceInstance) DeepCopy() *BreakoutResourceInstance {
	if in == nil {
		return nil
	}
	out := new(BreakoutResourceInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreakoutResourceMap) DeepCopyInto(out *BreakoutResourceMap) {
	*out = *in
	if in.ResourceMap != nil {
		in, out := &in.ResourceMap, &out.ResourceMap
		*out = make(map[string][]BreakoutResourceInstance, len(*in))
		for key, val := range *in {
			var outVal []BreakoutResourceInstance
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]BreakoutResourceInstance, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreakoutResourceMap.
func (in *BreakoutResourceMap) DeepCopy() *BreakoutResourceMap {
	if in == nil {
		return nil
	}
	out := new(BreakoutResourceMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreakoutSpec) DeepCopyInto(out *BreakoutSpec) {
	*out = *in
	if in.ModuleRefs != nil {
		in, out := &in.ModuleRefs, &out.ModuleRefs
		*out = make([]ModuleRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.RoutingRules.DeepCopyInto(&out.RoutingRules)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreakoutSpec.
func (in *BreakoutSpec) DeepCopy() *BreakoutSpec {
	if in == nil {
		return nil
	}
	out := new(BreakoutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreakoutStatus) DeepCopyInto(out *BreakoutStatus) {
	*out = *in
	if in.ResourceTypeMap != nil {
		in, out := &in.ResourceTypeMap, &out.ResourceTypeMap
		*out = make(map[string]map[string]BreakoutResourceMap, len(*in))
		for key, val := range *in {
			var outVal map[string]BreakoutResourceMap
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]BreakoutResourceMap, len(*in))
				for key, val := range *in {
					(*out)[key] = *val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreakoutStatus.
func (in *BreakoutStatus) DeepCopy() *BreakoutStatus {
	if in == nil {
		return nil
	}
	out := new(BreakoutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleRef) DeepCopyInto(out *ModuleRef) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleRef.
func (in *ModuleRef) DeepCopy() *ModuleRef {
	if in == nil {
		return nil
	}
	out := new(ModuleRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingRuleSpec) DeepCopyInto(out *RoutingRuleSpec) {
	*out = *in
	in.HTTPRoute.DeepCopyInto(&out.HTTPRoute)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingRuleSpec.
func (in *RoutingRuleSpec) DeepCopy() *RoutingRuleSpec {
	if in == nil {
		return nil
	}
	out := new(RoutingRuleSpec)
	in.DeepCopyInto(out)
	return out
}
