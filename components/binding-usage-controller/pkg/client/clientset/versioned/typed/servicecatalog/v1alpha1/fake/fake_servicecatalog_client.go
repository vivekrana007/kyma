// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kyma-project/kyma/components/binding-usage-controller/pkg/client/clientset/versioned/typed/servicecatalog/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeServicecatalogV1alpha1 struct {
	*testing.Fake
}

func (c *FakeServicecatalogV1alpha1) ServiceBindingUsages(namespace string) v1alpha1.ServiceBindingUsageInterface {
	return &FakeServiceBindingUsages{c, namespace}
}

func (c *FakeServicecatalogV1alpha1) UsageKinds() v1alpha1.UsageKindInterface {
	return &FakeUsageKinds{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeServicecatalogV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
