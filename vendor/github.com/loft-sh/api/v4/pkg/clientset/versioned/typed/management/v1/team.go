// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// TeamsGetter has a method to return a TeamInterface.
// A group's client should implement this interface.
type TeamsGetter interface {
	Teams() TeamInterface
}

// TeamInterface has methods to work with Team resources.
type TeamInterface interface {
	Create(ctx context.Context, team *managementv1.Team, opts metav1.CreateOptions) (*managementv1.Team, error)
	Update(ctx context.Context, team *managementv1.Team, opts metav1.UpdateOptions) (*managementv1.Team, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, team *managementv1.Team, opts metav1.UpdateOptions) (*managementv1.Team, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*managementv1.Team, error)
	List(ctx context.Context, opts metav1.ListOptions) (*managementv1.TeamList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *managementv1.Team, err error)
	ListClusters(ctx context.Context, teamName string, options metav1.GetOptions) (*managementv1.TeamClusters, error)
	ListAccessKeys(ctx context.Context, teamName string, options metav1.GetOptions) (*managementv1.TeamAccessKeys, error)

	TeamExpansion
}

// teams implements TeamInterface
type teams struct {
	*gentype.ClientWithList[*managementv1.Team, *managementv1.TeamList]
}

// newTeams returns a Teams
func newTeams(c *ManagementV1Client) *teams {
	return &teams{
		gentype.NewClientWithList[*managementv1.Team, *managementv1.TeamList](
			"teams",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *managementv1.Team { return &managementv1.Team{} },
			func() *managementv1.TeamList { return &managementv1.TeamList{} },
		),
	}
}

// ListClusters takes name of the team, and returns the corresponding managementv1.TeamClusters object, and an error if there is any.
func (c *teams) ListClusters(ctx context.Context, teamName string, options metav1.GetOptions) (result *managementv1.TeamClusters, err error) {
	result = &managementv1.TeamClusters{}
	err = c.GetClient().Get().
		Resource("teams").
		Name(teamName).
		SubResource("clusters").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// ListAccessKeys takes name of the team, and returns the corresponding managementv1.TeamAccessKeys object, and an error if there is any.
func (c *teams) ListAccessKeys(ctx context.Context, teamName string, options metav1.GetOptions) (result *managementv1.TeamAccessKeys, err error) {
	result = &managementv1.TeamAccessKeys{}
	err = c.GetClient().Get().
		Resource("teams").
		Name(teamName).
		SubResource("accesskeys").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}
