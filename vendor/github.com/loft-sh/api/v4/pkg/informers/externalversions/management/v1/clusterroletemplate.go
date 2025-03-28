// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apismanagementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	versioned "github.com/loft-sh/api/v4/pkg/clientset/versioned"
	internalinterfaces "github.com/loft-sh/api/v4/pkg/informers/externalversions/internalinterfaces"
	managementv1 "github.com/loft-sh/api/v4/pkg/listers/management/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterRoleTemplateInformer provides access to a shared informer and lister for
// ClusterRoleTemplates.
type ClusterRoleTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() managementv1.ClusterRoleTemplateLister
}

type clusterRoleTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterRoleTemplateInformer constructs a new informer for ClusterRoleTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterRoleTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterRoleTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterRoleTemplateInformer constructs a new informer for ClusterRoleTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterRoleTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1().ClusterRoleTemplates().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1().ClusterRoleTemplates().Watch(context.TODO(), options)
			},
		},
		&apismanagementv1.ClusterRoleTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterRoleTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterRoleTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterRoleTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apismanagementv1.ClusterRoleTemplate{}, f.defaultInformer)
}

func (f *clusterRoleTemplateInformer) Lister() managementv1.ClusterRoleTemplateLister {
	return managementv1.NewClusterRoleTemplateLister(f.Informer().GetIndexer())
}
