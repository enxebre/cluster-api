/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	machinev1beta1 "sigs.k8s.io/cluster-api/pkg/apis/machine/v1beta1"
	clientset "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
	internalinterfaces "sigs.k8s.io/cluster-api/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1beta1 "sigs.k8s.io/cluster-api/pkg/client/listers_generated/machine/v1beta1"
)

// MachineDeploymentInformer provides access to a shared informer and lister for
// MachineDeployments.
type MachineDeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.MachineDeploymentLister
}

type machineDeploymentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMachineDeploymentInformer constructs a new informer for MachineDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineDeploymentInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMachineDeploymentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMachineDeploymentInformer constructs a new informer for MachineDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMachineDeploymentInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1beta1().MachineDeployments(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1beta1().MachineDeployments(namespace).Watch(options)
			},
		},
		&machinev1beta1.MachineDeployment{},
		resyncPeriod,
		indexers,
	)
}

func (f *machineDeploymentInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMachineDeploymentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *machineDeploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machinev1beta1.MachineDeployment{}, f.defaultInformer)
}

func (f *machineDeploymentInformer) Lister() v1beta1.MachineDeploymentLister {
	return v1beta1.NewMachineDeploymentLister(f.Informer().GetIndexer())
}
