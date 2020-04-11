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

package versioned

import (
	"fmt"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"

	azurebackendpoolsv1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/clientset/versioned/typed/azurebackendpool/v1"
	azureingressprohibitedtargetsv1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/clientset/versioned/typed/azureingressprohibitedtarget/v1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AzurebackendpoolsV1() azurebackendpoolsv1.AzurebackendpoolsV1Interface
	AzureingressprohibitedtargetsV1() azureingressprohibitedtargetsv1.AzureingressprohibitedtargetsV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	azurebackendpoolsV1             *azurebackendpoolsv1.AzurebackendpoolsV1Client
	azureingressprohibitedtargetsV1 *azureingressprohibitedtargetsv1.AzureingressprohibitedtargetsV1Client
}

// AzurebackendpoolsV1 retrieves the AzurebackendpoolsV1Client
func (c *Clientset) AzurebackendpoolsV1() azurebackendpoolsv1.AzurebackendpoolsV1Interface {
	return c.azurebackendpoolsV1
}

// AzureingressprohibitedtargetsV1 retrieves the AzureingressprohibitedtargetsV1Client
func (c *Clientset) AzureingressprohibitedtargetsV1() azureingressprohibitedtargetsv1.AzureingressprohibitedtargetsV1Interface {
	return c.azureingressprohibitedtargetsV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.azurebackendpoolsV1, err = azurebackendpoolsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.azureingressprohibitedtargetsV1, err = azureingressprohibitedtargetsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.azurebackendpoolsV1 = azurebackendpoolsv1.NewForConfigOrDie(c)
	cs.azureingressprohibitedtargetsV1 = azureingressprohibitedtargetsv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.azurebackendpoolsV1 = azurebackendpoolsv1.New(c)
	cs.azureingressprohibitedtargetsV1 = azureingressprohibitedtargetsv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
