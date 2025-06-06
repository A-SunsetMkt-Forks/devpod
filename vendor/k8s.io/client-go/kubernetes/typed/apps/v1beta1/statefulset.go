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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsappsv1beta1 "k8s.io/client-go/applyconfigurations/apps/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// StatefulSetsGetter has a method to return a StatefulSetInterface.
// A group's client should implement this interface.
type StatefulSetsGetter interface {
	StatefulSets(namespace string) StatefulSetInterface
}

// StatefulSetInterface has methods to work with StatefulSet resources.
type StatefulSetInterface interface {
	Create(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts v1.CreateOptions) (*appsv1beta1.StatefulSet, error)
	Update(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts v1.UpdateOptions) (*appsv1beta1.StatefulSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts v1.UpdateOptions) (*appsv1beta1.StatefulSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*appsv1beta1.StatefulSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*appsv1beta1.StatefulSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *appsv1beta1.StatefulSet, err error)
	Apply(ctx context.Context, statefulSet *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration, opts v1.ApplyOptions) (result *appsv1beta1.StatefulSet, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, statefulSet *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration, opts v1.ApplyOptions) (result *appsv1beta1.StatefulSet, err error)
	StatefulSetExpansion
}

// statefulSets implements StatefulSetInterface
type statefulSets struct {
	*gentype.ClientWithListAndApply[*appsv1beta1.StatefulSet, *appsv1beta1.StatefulSetList, *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration]
}

// newStatefulSets returns a StatefulSets
func newStatefulSets(c *AppsV1beta1Client, namespace string) *statefulSets {
	return &statefulSets{
		gentype.NewClientWithListAndApply[*appsv1beta1.StatefulSet, *appsv1beta1.StatefulSetList, *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration](
			"statefulsets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *appsv1beta1.StatefulSet { return &appsv1beta1.StatefulSet{} },
			func() *appsv1beta1.StatefulSetList { return &appsv1beta1.StatefulSetList{} },
			gentype.PrefersProtobuf[*appsv1beta1.StatefulSet](),
		),
	}
}
