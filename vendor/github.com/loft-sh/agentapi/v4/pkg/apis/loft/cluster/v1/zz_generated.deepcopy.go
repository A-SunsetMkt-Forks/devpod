//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bash) DeepCopyInto(out *Bash) {
	*out = *in
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bash.
func (in *Bash) DeepCopy() *Bash {
	if in == nil {
		return nil
	}
	out := new(Bash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chart) DeepCopyInto(out *Chart) {
	*out = *in
	if in.UsernameRef != nil {
		in, out := &in.UsernameRef, &out.UsernameRef
		*out = new(ChartSecretRef)
		(*in).DeepCopyInto(*out)
	}
	if in.PasswordRef != nil {
		in, out := &in.PasswordRef, &out.PasswordRef
		*out = new(ChartSecretRef)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chart.
func (in *Chart) DeepCopy() *Chart {
	if in == nil {
		return nil
	}
	out := new(Chart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartInfo) DeepCopyInto(out *ChartInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartInfo.
func (in *ChartInfo) DeepCopy() *ChartInfo {
	if in == nil {
		return nil
	}
	out := new(ChartInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChartInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartInfoList) DeepCopyInto(out *ChartInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChartInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartInfoList.
func (in *ChartInfoList) DeepCopy() *ChartInfoList {
	if in == nil {
		return nil
	}
	out := new(ChartInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChartInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartInfoSpec) DeepCopyInto(out *ChartInfoSpec) {
	*out = *in
	in.Chart.DeepCopyInto(&out.Chart)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartInfoSpec.
func (in *ChartInfoSpec) DeepCopy() *ChartInfoSpec {
	if in == nil {
		return nil
	}
	out := new(ChartInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartInfoStatus) DeepCopyInto(out *ChartInfoStatus) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(Metadata)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartInfoStatus.
func (in *ChartInfoStatus) DeepCopy() *ChartInfoStatus {
	if in == nil {
		return nil
	}
	out := new(ChartInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSecretRef) DeepCopyInto(out *ChartSecretRef) {
	*out = *in
	if in.ProjectSecretRef != nil {
		in, out := &in.ProjectSecretRef, &out.ProjectSecretRef
		*out = new(ProjectSecretRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSecretRef.
func (in *ChartSecretRef) DeepCopy() *ChartSecretRef {
	if in == nil {
		return nil
	}
	out := new(ChartSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EpochInfo) DeepCopyInto(out *EpochInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EpochInfo.
func (in *EpochInfo) DeepCopy() *EpochInfo {
	if in == nil {
		return nil
	}
	out := new(EpochInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Feature) DeepCopyInto(out *Feature) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Feature.
func (in *Feature) DeepCopy() *Feature {
	if in == nil {
		return nil
	}
	out := new(Feature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Feature) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureList) DeepCopyInto(out *FeatureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Feature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureList.
func (in *FeatureList) DeepCopy() *FeatureList {
	if in == nil {
		return nil
	}
	out := new(FeatureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSpec) DeepCopyInto(out *FeatureSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSpec.
func (in *FeatureSpec) DeepCopy() *FeatureSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureStatus) DeepCopyInto(out *FeatureStatus) {
	*out = *in
	out.Feature = in.Feature
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureStatus.
func (in *FeatureStatus) DeepCopy() *FeatureStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRelease) DeepCopyInto(out *HelmRelease) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRelease.
func (in *HelmRelease) DeepCopy() *HelmRelease {
	if in == nil {
		return nil
	}
	out := new(HelmRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmRelease) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseApp) DeepCopyInto(out *HelmReleaseApp) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseApp.
func (in *HelmReleaseApp) DeepCopy() *HelmReleaseApp {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseConfig) DeepCopyInto(out *HelmReleaseConfig) {
	*out = *in
	in.Chart.DeepCopyInto(&out.Chart)
	if in.Bash != nil {
		in, out := &in.Bash, &out.Bash
		*out = new(Bash)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseConfig.
func (in *HelmReleaseConfig) DeepCopy() *HelmReleaseConfig {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseList) DeepCopyInto(out *HelmReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmRelease, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseList.
func (in *HelmReleaseList) DeepCopy() *HelmReleaseList {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseSpec) DeepCopyInto(out *HelmReleaseSpec) {
	*out = *in
	in.HelmReleaseConfig.DeepCopyInto(&out.HelmReleaseConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseSpec.
func (in *HelmReleaseSpec) DeepCopy() *HelmReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseStatus) DeepCopyInto(out *HelmReleaseStatus) {
	*out = *in
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = new(Info)
		(*in).DeepCopyInto(*out)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(Metadata)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseStatus.
func (in *HelmReleaseStatus) DeepCopy() *HelmReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Info) DeepCopyInto(out *Info) {
	*out = *in
	in.FirstDeployed.DeepCopyInto(&out.FirstDeployed)
	in.LastDeployed.DeepCopyInto(&out.LastDeployed)
	in.Deleted.DeepCopyInto(&out.Deleted)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Info.
func (in *Info) DeepCopy() *Info {
	if in == nil {
		return nil
	}
	out := new(Info)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastActivityInfo) DeepCopyInto(out *LastActivityInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastActivityInfo.
func (in *LastActivityInfo) DeepCopy() *LastActivityInfo {
	if in == nil {
		return nil
	}
	out := new(LastActivityInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Maintainer) DeepCopyInto(out *Maintainer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Maintainer.
func (in *Maintainer) DeepCopy() *Maintainer {
	if in == nil {
		return nil
	}
	out := new(Maintainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Maintainers != nil {
		in, out := &in.Maintainers, &out.Maintainers
		*out = make([]*Maintainer, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Maintainer)
				**out = **in
			}
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Urls != nil {
		in, out := &in.Urls, &out.Urls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSecretRef) DeepCopyInto(out *ProjectSecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSecretRef.
func (in *ProjectSecretRef) DeepCopy() *ProjectSecretRef {
	if in == nil {
		return nil
	}
	out := new(ProjectSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepModeConfig) DeepCopyInto(out *SleepModeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepModeConfig.
func (in *SleepModeConfig) DeepCopy() *SleepModeConfig {
	if in == nil {
		return nil
	}
	out := new(SleepModeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SleepModeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepModeConfigSpec) DeepCopyInto(out *SleepModeConfigSpec) {
	*out = *in
	if in.ForceSleepDuration != nil {
		in, out := &in.ForceSleepDuration, &out.ForceSleepDuration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepModeConfigSpec.
func (in *SleepModeConfigSpec) DeepCopy() *SleepModeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SleepModeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepModeConfigStatus) DeepCopyInto(out *SleepModeConfigStatus) {
	*out = *in
	if in.LastActivityInfo != nil {
		in, out := &in.LastActivityInfo, &out.LastActivityInfo
		*out = new(LastActivityInfo)
		**out = **in
	}
	if in.CurrentEpoch != nil {
		in, out := &in.CurrentEpoch, &out.CurrentEpoch
		*out = new(EpochInfo)
		**out = **in
	}
	if in.LastEpoch != nil {
		in, out := &in.LastEpoch, &out.LastEpoch
		*out = new(EpochInfo)
		**out = **in
	}
	if in.SleptLastThirtyDays != nil {
		in, out := &in.SleptLastThirtyDays, &out.SleptLastThirtyDays
		*out = new(float64)
		**out = **in
	}
	if in.SleptLastSevenDays != nil {
		in, out := &in.SleptLastSevenDays, &out.SleptLastSevenDays
		*out = new(float64)
		**out = **in
	}
	if in.ScheduledSleep != nil {
		in, out := &in.ScheduledSleep, &out.ScheduledSleep
		*out = new(int64)
		**out = **in
	}
	if in.ScheduledWakeup != nil {
		in, out := &in.ScheduledWakeup, &out.ScheduledWakeup
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepModeConfigStatus.
func (in *SleepModeConfigStatus) DeepCopy() *SleepModeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SleepModeConfigStatus)
	in.DeepCopyInto(out)
	return out
}
