/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	marshal "github.com/jetstack-experimental/navigator/pkg/apis/marshal"
	internalclientset "github.com/jetstack-experimental/navigator/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/jetstack-experimental/navigator/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/jetstack-experimental/navigator/pkg/client/listers_generated/marshal/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ElasticsearchClusterInformer provides access to a shared informer and lister for
// ElasticsearchClusters.
type ElasticsearchClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ElasticsearchClusterLister
}

type elasticsearchClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newElasticsearchClusterInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Marshal().ElasticsearchClusters(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Marshal().ElasticsearchClusters(v1.NamespaceAll).Watch(options)
			},
		},
		&marshal.ElasticsearchCluster{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *elasticsearchClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&marshal.ElasticsearchCluster{}, newElasticsearchClusterInformer)
}

func (f *elasticsearchClusterInformer) Lister() internalversion.ElasticsearchClusterLister {
	return internalversion.NewElasticsearchClusterLister(f.Informer().GetIndexer())
}
