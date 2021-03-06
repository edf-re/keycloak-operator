// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	keycloakv1alpha1 "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	versioned "github.com/keycloak/keycloak-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/keycloak/keycloak-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/keycloak/keycloak-operator/pkg/client/listers/keycloak/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KeycloakRealmInformer provides access to a shared informer and lister for
// KeycloakRealms.
type KeycloakRealmInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KeycloakRealmLister
}

type keycloakRealmInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKeycloakRealmInformer constructs a new informer for KeycloakRealm type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKeycloakRealmInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKeycloakRealmInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKeycloakRealmInformer constructs a new informer for KeycloakRealm type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKeycloakRealmInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KeycloakV1alpha1().KeycloakRealms(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KeycloakV1alpha1().KeycloakRealms(namespace).Watch(context.TODO(), options)
			},
		},
		&keycloakv1alpha1.KeycloakRealm{},
		resyncPeriod,
		indexers,
	)
}

func (f *keycloakRealmInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKeycloakRealmInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *keycloakRealmInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&keycloakv1alpha1.KeycloakRealm{}, f.defaultInformer)
}

func (f *keycloakRealmInformer) Lister() v1alpha1.KeycloakRealmLister {
	return v1alpha1.NewKeycloakRealmLister(f.Informer().GetIndexer())
}
