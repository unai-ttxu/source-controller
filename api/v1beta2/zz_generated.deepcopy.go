//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Flux authors

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

package v1beta2

import (
	"github.com/fluxcd/pkg/apis/acl"
	"github.com/fluxcd/pkg/apis/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Artifact) DeepCopyInto(out *Artifact) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int64)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Artifact.
func (in *Artifact) DeepCopy() *Artifact {
	if in == nil {
		return nil
	}
	out := new(Artifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
	out.Interval = in.Interval
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Ignore != nil {
		in, out := &in.Ignore, &out.Ignore
		*out = new(string)
		**out = **in
	}
	if in.AccessFrom != nil {
		in, out := &in.AccessFrom, &out.AccessFrom
		*out = new(acl.AccessFrom)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Artifact != nil {
		in, out := &in.Artifact, &out.Artifact
		*out = new(Artifact)
		(*in).DeepCopyInto(*out)
	}
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepository) DeepCopyInto(out *GitRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepository.
func (in *GitRepository) DeepCopy() *GitRepository {
	if in == nil {
		return nil
	}
	out := new(GitRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepositoryInclude) DeepCopyInto(out *GitRepositoryInclude) {
	*out = *in
	out.GitRepositoryRef = in.GitRepositoryRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepositoryInclude.
func (in *GitRepositoryInclude) DeepCopy() *GitRepositoryInclude {
	if in == nil {
		return nil
	}
	out := new(GitRepositoryInclude)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepositoryList) DeepCopyInto(out *GitRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepositoryList.
func (in *GitRepositoryList) DeepCopy() *GitRepositoryList {
	if in == nil {
		return nil
	}
	out := new(GitRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepositoryRef) DeepCopyInto(out *GitRepositoryRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepositoryRef.
func (in *GitRepositoryRef) DeepCopy() *GitRepositoryRef {
	if in == nil {
		return nil
	}
	out := new(GitRepositoryRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepositorySpec) DeepCopyInto(out *GitRepositorySpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
	out.Interval = in.Interval
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(GitRepositoryRef)
		**out = **in
	}
	if in.Verification != nil {
		in, out := &in.Verification, &out.Verification
		*out = new(GitRepositoryVerification)
		**out = **in
	}
	if in.Ignore != nil {
		in, out := &in.Ignore, &out.Ignore
		*out = new(string)
		**out = **in
	}
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]GitRepositoryInclude, len(*in))
		copy(*out, *in)
	}
	if in.AccessFrom != nil {
		in, out := &in.AccessFrom, &out.AccessFrom
		*out = new(acl.AccessFrom)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepositorySpec.
func (in *GitRepositorySpec) DeepCopy() *GitRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(GitRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepositoryStatus) DeepCopyInto(out *GitRepositoryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Artifact != nil {
		in, out := &in.Artifact, &out.Artifact
		*out = new(Artifact)
		(*in).DeepCopyInto(*out)
	}
	if in.IncludedArtifacts != nil {
		in, out := &in.IncludedArtifacts, &out.IncludedArtifacts
		*out = make([]*Artifact, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Artifact)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepositoryStatus.
func (in *GitRepositoryStatus) DeepCopy() *GitRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(GitRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepositoryVerification) DeepCopyInto(out *GitRepositoryVerification) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepositoryVerification.
func (in *GitRepositoryVerification) DeepCopy() *GitRepositoryVerification {
	if in == nil {
		return nil
	}
	out := new(GitRepositoryVerification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChart) DeepCopyInto(out *HelmChart) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChart.
func (in *HelmChart) DeepCopy() *HelmChart {
	if in == nil {
		return nil
	}
	out := new(HelmChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmChart) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartList) DeepCopyInto(out *HelmChartList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmChart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartList.
func (in *HelmChartList) DeepCopy() *HelmChartList {
	if in == nil {
		return nil
	}
	out := new(HelmChartList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmChartList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartSpec) DeepCopyInto(out *HelmChartSpec) {
	*out = *in
	out.SourceRef = in.SourceRef
	out.Interval = in.Interval
	if in.ValuesFiles != nil {
		in, out := &in.ValuesFiles, &out.ValuesFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessFrom != nil {
		in, out := &in.AccessFrom, &out.AccessFrom
		*out = new(acl.AccessFrom)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartSpec.
func (in *HelmChartSpec) DeepCopy() *HelmChartSpec {
	if in == nil {
		return nil
	}
	out := new(HelmChartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartStatus) DeepCopyInto(out *HelmChartStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Artifact != nil {
		in, out := &in.Artifact, &out.Artifact
		*out = new(Artifact)
		(*in).DeepCopyInto(*out)
	}
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartStatus.
func (in *HelmChartStatus) DeepCopy() *HelmChartStatus {
	if in == nil {
		return nil
	}
	out := new(HelmChartStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRepository) DeepCopyInto(out *HelmRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRepository.
func (in *HelmRepository) DeepCopy() *HelmRepository {
	if in == nil {
		return nil
	}
	out := new(HelmRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRepositoryList) DeepCopyInto(out *HelmRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRepositoryList.
func (in *HelmRepositoryList) DeepCopy() *HelmRepositoryList {
	if in == nil {
		return nil
	}
	out := new(HelmRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRepositorySpec) DeepCopyInto(out *HelmRepositorySpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
	out.Interval = in.Interval
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.AccessFrom != nil {
		in, out := &in.AccessFrom, &out.AccessFrom
		*out = new(acl.AccessFrom)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRepositorySpec.
func (in *HelmRepositorySpec) DeepCopy() *HelmRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(HelmRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRepositoryStatus) DeepCopyInto(out *HelmRepositoryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Artifact != nil {
		in, out := &in.Artifact, &out.Artifact
		*out = new(Artifact)
		(*in).DeepCopyInto(*out)
	}
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRepositoryStatus.
func (in *HelmRepositoryStatus) DeepCopy() *HelmRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(HelmRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalHelmChartSourceReference) DeepCopyInto(out *LocalHelmChartSourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalHelmChartSourceReference.
func (in *LocalHelmChartSourceReference) DeepCopy() *LocalHelmChartSourceReference {
	if in == nil {
		return nil
	}
	out := new(LocalHelmChartSourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCILayerSelector) DeepCopyInto(out *OCILayerSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCILayerSelector.
func (in *OCILayerSelector) DeepCopy() *OCILayerSelector {
	if in == nil {
		return nil
	}
	out := new(OCILayerSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIRepository) DeepCopyInto(out *OCIRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIRepository.
func (in *OCIRepository) DeepCopy() *OCIRepository {
	if in == nil {
		return nil
	}
	out := new(OCIRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIRepositoryList) DeepCopyInto(out *OCIRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OCIRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIRepositoryList.
func (in *OCIRepositoryList) DeepCopy() *OCIRepositoryList {
	if in == nil {
		return nil
	}
	out := new(OCIRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIRepositoryRef) DeepCopyInto(out *OCIRepositoryRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIRepositoryRef.
func (in *OCIRepositoryRef) DeepCopy() *OCIRepositoryRef {
	if in == nil {
		return nil
	}
	out := new(OCIRepositoryRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIRepositorySpec) DeepCopyInto(out *OCIRepositorySpec) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(OCIRepositoryRef)
		**out = **in
	}
	if in.LayerSelector != nil {
		in, out := &in.LayerSelector, &out.LayerSelector
		*out = new(OCILayerSelector)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
	if in.Verify != nil {
		in, out := &in.Verify, &out.Verify
		*out = new(OCIRepositoryVerification)
		(*in).DeepCopyInto(*out)
	}
	if in.CertSecretRef != nil {
		in, out := &in.CertSecretRef, &out.CertSecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
	out.Interval = in.Interval
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Ignore != nil {
		in, out := &in.Ignore, &out.Ignore
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIRepositorySpec.
func (in *OCIRepositorySpec) DeepCopy() *OCIRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(OCIRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIRepositoryStatus) DeepCopyInto(out *OCIRepositoryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Artifact != nil {
		in, out := &in.Artifact, &out.Artifact
		*out = new(Artifact)
		(*in).DeepCopyInto(*out)
	}
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIRepositoryStatus.
func (in *OCIRepositoryStatus) DeepCopy() *OCIRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(OCIRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIRepositoryVerification) DeepCopyInto(out *OCIRepositoryVerification) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIRepositoryVerification.
func (in *OCIRepositoryVerification) DeepCopy() *OCIRepositoryVerification {
	if in == nil {
		return nil
	}
	out := new(OCIRepositoryVerification)
	in.DeepCopyInto(out)
	return out
}
