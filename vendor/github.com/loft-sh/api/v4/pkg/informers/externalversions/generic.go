// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	virtualclusterv1 "github.com/loft-sh/api/v4/pkg/apis/virtualcluster/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=management.loft.sh, Version=v1
	case v1.SchemeGroupVersion.WithResource("agentauditevents"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().AgentAuditEvents().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("announcements"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Announcements().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("apps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Apps().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("backups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Backups().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Clusters().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("clusteraccesses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().ClusterAccesses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("clusterroletemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().ClusterRoleTemplates().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("configs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Configs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("convertvirtualclusterconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().ConvertVirtualClusterConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("devpodenvironmenttemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().DevPodEnvironmentTemplates().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("devpodworkspaceinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().DevPodWorkspaceInstances().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("devpodworkspacetemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().DevPodWorkspaceTemplates().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("directclusterendpointtokens"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().DirectClusterEndpointTokens().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Events().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("features"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Features().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ingressauthtokens"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().IngressAuthTokens().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("licenses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Licenses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("licensetokens"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().LicenseTokens().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("loftupgrades"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().LoftUpgrades().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("oidcclients"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().OIDCClients().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ownedaccesskeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().OwnedAccessKeys().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Projects().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("projectsecrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().ProjectSecrets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("redirecttokens"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().RedirectTokens().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("registervirtualclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().RegisterVirtualClusters().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("resetaccesskeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().ResetAccessKeys().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("runners"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Runners().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("selves"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Selves().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("selfsubjectaccessreviews"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().SelfSubjectAccessReviews().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("sharedsecrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().SharedSecrets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("spaceinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().SpaceInstances().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("spacetemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().SpaceTemplates().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("subjectaccessreviews"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().SubjectAccessReviews().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("tasks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Tasks().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("teams"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Teams().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("translatevclusterresourcenames"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().TranslateVClusterResourceNames().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().Users().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualclusterinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().VirtualClusterInstances().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualclustertemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Management().V1().VirtualClusterTemplates().Informer()}, nil

		// Group=storage.loft.sh, Version=v1
	case storagev1.SchemeGroupVersion.WithResource("accesskeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().AccessKeys().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("apps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().Apps().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().Clusters().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("clusteraccesses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().ClusterAccesses().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("clusterroletemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().ClusterRoleTemplates().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("devpodenvironmenttemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().DevPodEnvironmentTemplates().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("devpodworkspaceinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().DevPodWorkspaceInstances().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("devpodworkspacetemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().DevPodWorkspaceTemplates().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("networkpeers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().NetworkPeers().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().Projects().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("runners"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().Runners().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("sharedsecrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().SharedSecrets().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("spaceinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().SpaceInstances().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("spacetemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().SpaceTemplates().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("tasks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().Tasks().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("teams"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().Teams().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().Users().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("virtualclusterinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().VirtualClusterInstances().Informer()}, nil
	case storagev1.SchemeGroupVersion.WithResource("virtualclustertemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1().VirtualClusterTemplates().Informer()}, nil

		// Group=virtualcluster.loft.sh, Version=v1
	case virtualclusterv1.SchemeGroupVersion.WithResource("helmreleases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Virtualcluster().V1().HelmReleases().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
