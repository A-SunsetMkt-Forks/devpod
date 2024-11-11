// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	versioned "github.com/loft-sh/api/v4/pkg/clientset/versioned"
	internalinterfaces "github.com/loft-sh/api/v4/pkg/informers/externalversions/internalinterfaces"
	v1 "github.com/loft-sh/api/v4/pkg/listers/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TeamInformer provides access to a shared informer and lister for
// Teams.
type TeamInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TeamLister
}

type teamInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTeamInformer constructs a new informer for Team type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTeamInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTeamInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTeamInformer constructs a new informer for Team type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTeamInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().Teams().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().Teams().Watch(context.TODO(), options)
			},
		},
		&storagev1.Team{},
		resyncPeriod,
		indexers,
	)
}

func (f *teamInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTeamInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *teamInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1.Team{}, f.defaultInformer)
}

func (f *teamInformer) Lister() v1.TeamLister {
	return v1.NewTeamLister(f.Informer().GetIndexer())
}
