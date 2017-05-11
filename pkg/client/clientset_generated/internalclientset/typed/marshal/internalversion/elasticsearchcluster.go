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

package internalversion

import (
	marshal "github.com/jetstack-experimental/navigator/pkg/apis/marshal"
	scheme "github.com/jetstack-experimental/navigator/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ElasticsearchClustersGetter has a method to return a ElasticsearchClusterInterface.
// A group's client should implement this interface.
type ElasticsearchClustersGetter interface {
	ElasticsearchClusters(namespace string) ElasticsearchClusterInterface
}

// ElasticsearchClusterInterface has methods to work with ElasticsearchCluster resources.
type ElasticsearchClusterInterface interface {
	Create(*marshal.ElasticsearchCluster) (*marshal.ElasticsearchCluster, error)
	Update(*marshal.ElasticsearchCluster) (*marshal.ElasticsearchCluster, error)
	UpdateStatus(*marshal.ElasticsearchCluster) (*marshal.ElasticsearchCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*marshal.ElasticsearchCluster, error)
	List(opts v1.ListOptions) (*marshal.ElasticsearchClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *marshal.ElasticsearchCluster, err error)
	ElasticsearchClusterExpansion
}

// elasticsearchClusters implements ElasticsearchClusterInterface
type elasticsearchClusters struct {
	client rest.Interface
	ns     string
}

// newElasticsearchClusters returns a ElasticsearchClusters
func newElasticsearchClusters(c *MarshalClient, namespace string) *elasticsearchClusters {
	return &elasticsearchClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a elasticsearchCluster and creates it.  Returns the server's representation of the elasticsearchCluster, and an error, if there is any.
func (c *elasticsearchClusters) Create(elasticsearchCluster *marshal.ElasticsearchCluster) (result *marshal.ElasticsearchCluster, err error) {
	result = &marshal.ElasticsearchCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("elasticsearchclusters").
		Body(elasticsearchCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a elasticsearchCluster and updates it. Returns the server's representation of the elasticsearchCluster, and an error, if there is any.
func (c *elasticsearchClusters) Update(elasticsearchCluster *marshal.ElasticsearchCluster) (result *marshal.ElasticsearchCluster, err error) {
	result = &marshal.ElasticsearchCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearchclusters").
		Name(elasticsearchCluster.Name).
		Body(elasticsearchCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *elasticsearchClusters) UpdateStatus(elasticsearchCluster *marshal.ElasticsearchCluster) (result *marshal.ElasticsearchCluster, err error) {
	result = &marshal.ElasticsearchCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearchclusters").
		Name(elasticsearchCluster.Name).
		SubResource("status").
		Body(elasticsearchCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the elasticsearchCluster and deletes it. Returns an error if one occurs.
func (c *elasticsearchClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearchclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elasticsearchClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearchclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the elasticsearchCluster, and returns the corresponding elasticsearchCluster object, and an error if there is any.
func (c *elasticsearchClusters) Get(name string, options v1.GetOptions) (result *marshal.ElasticsearchCluster, err error) {
	result = &marshal.ElasticsearchCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ElasticsearchClusters that match those selectors.
func (c *elasticsearchClusters) List(opts v1.ListOptions) (result *marshal.ElasticsearchClusterList, err error) {
	result = &marshal.ElasticsearchClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elasticsearchClusters.
func (c *elasticsearchClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearchclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched elasticsearchCluster.
func (c *elasticsearchClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *marshal.ElasticsearchCluster, err error) {
	result = &marshal.ElasticsearchCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("elasticsearchclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
