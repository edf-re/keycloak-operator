// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/keycloak/keycloak-operator/pkg/client/clientset/versioned/typed/keycloak/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKeycloakV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKeycloakV1alpha1) Keycloaks(namespace string) v1alpha1.KeycloakInterface {
	return &FakeKeycloaks{c, namespace}
}

func (c *FakeKeycloakV1alpha1) KeycloakBackups(namespace string) v1alpha1.KeycloakBackupInterface {
	return &FakeKeycloakBackups{c, namespace}
}

func (c *FakeKeycloakV1alpha1) KeycloakClients(namespace string) v1alpha1.KeycloakClientInterface {
	return &FakeKeycloakClients{c, namespace}
}

func (c *FakeKeycloakV1alpha1) KeycloakRealms(namespace string) v1alpha1.KeycloakRealmInterface {
	return &FakeKeycloakRealms{c, namespace}
}

func (c *FakeKeycloakV1alpha1) KeycloakUsers(namespace string) v1alpha1.KeycloakUserInterface {
	return &FakeKeycloakUsers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKeycloakV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
