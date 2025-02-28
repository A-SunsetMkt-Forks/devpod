// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// DevPodWorkspaceInstancesGetter has a method to return a DevPodWorkspaceInstanceInterface.
// A group's client should implement this interface.
type DevPodWorkspaceInstancesGetter interface {
	DevPodWorkspaceInstances(namespace string) DevPodWorkspaceInstanceInterface
}

// DevPodWorkspaceInstanceInterface has methods to work with DevPodWorkspaceInstance resources.
type DevPodWorkspaceInstanceInterface interface {
	Create(ctx context.Context, devPodWorkspaceInstance *v1.DevPodWorkspaceInstance, opts metav1.CreateOptions) (*v1.DevPodWorkspaceInstance, error)
	Update(ctx context.Context, devPodWorkspaceInstance *v1.DevPodWorkspaceInstance, opts metav1.UpdateOptions) (*v1.DevPodWorkspaceInstance, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DevPodWorkspaceInstance, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DevPodWorkspaceInstanceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DevPodWorkspaceInstance, err error)
	GetState(ctx context.Context, devPodWorkspaceInstanceName string, options metav1.GetOptions) (*v1.DevPodWorkspaceInstanceState, error)
	SetState(ctx context.Context, devPodWorkspaceInstanceName string, devPodWorkspaceInstanceState *v1.DevPodWorkspaceInstanceState, opts metav1.CreateOptions) (*v1.DevPodWorkspaceInstanceState, error)
	Troubleshoot(ctx context.Context, devPodWorkspaceInstanceName string, options metav1.GetOptions) (*v1.DevPodWorkspaceInstanceTroubleshoot, error)

	DevPodWorkspaceInstanceExpansion
}

// devPodWorkspaceInstances implements DevPodWorkspaceInstanceInterface
type devPodWorkspaceInstances struct {
	*gentype.ClientWithList[*v1.DevPodWorkspaceInstance, *v1.DevPodWorkspaceInstanceList]
}

// newDevPodWorkspaceInstances returns a DevPodWorkspaceInstances
func newDevPodWorkspaceInstances(c *ManagementV1Client, namespace string) *devPodWorkspaceInstances {
	return &devPodWorkspaceInstances{
		gentype.NewClientWithList[*v1.DevPodWorkspaceInstance, *v1.DevPodWorkspaceInstanceList](
			"devpodworkspaceinstances",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.DevPodWorkspaceInstance { return &v1.DevPodWorkspaceInstance{} },
			func() *v1.DevPodWorkspaceInstanceList { return &v1.DevPodWorkspaceInstanceList{} }),
	}
}

// GetState takes name of the devPodWorkspaceInstance, and returns the corresponding v1.DevPodWorkspaceInstanceState object, and an error if there is any.
func (c *devPodWorkspaceInstances) GetState(ctx context.Context, devPodWorkspaceInstanceName string, options metav1.GetOptions) (result *v1.DevPodWorkspaceInstanceState, err error) {
	result = &v1.DevPodWorkspaceInstanceState{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("devpodworkspaceinstances").
		Name(devPodWorkspaceInstanceName).
		SubResource("state").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// SetState takes the representation of a devPodWorkspaceInstanceState and creates it.  Returns the server's representation of the devPodWorkspaceInstanceState, and an error, if there is any.
func (c *devPodWorkspaceInstances) SetState(ctx context.Context, devPodWorkspaceInstanceName string, devPodWorkspaceInstanceState *v1.DevPodWorkspaceInstanceState, opts metav1.CreateOptions) (result *v1.DevPodWorkspaceInstanceState, err error) {
	result = &v1.DevPodWorkspaceInstanceState{}
	err = c.GetClient().Post().
		Namespace(c.GetNamespace()).
		Resource("devpodworkspaceinstances").
		Name(devPodWorkspaceInstanceName).
		SubResource("state").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(devPodWorkspaceInstanceState).
		Do(ctx).
		Into(result)
	return
}

// Troubleshoot takes name of the devPodWorkspaceInstance, and returns the corresponding v1.DevPodWorkspaceInstanceTroubleshoot object, and an error if there is any.
func (c *devPodWorkspaceInstances) Troubleshoot(ctx context.Context, devPodWorkspaceInstanceName string, options metav1.GetOptions) (result *v1.DevPodWorkspaceInstanceTroubleshoot, err error) {
	result = &v1.DevPodWorkspaceInstanceTroubleshoot{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("devpodworkspaceinstances").
		Name(devPodWorkspaceInstanceName).
		SubResource("troubleshoot").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}
