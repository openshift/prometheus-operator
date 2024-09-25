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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ResourceQuotaLister helps list ResourceQuotas.
// All objects returned here must be treated as read-only.
type ResourceQuotaLister interface {
	// List lists all ResourceQuotas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ResourceQuota, err error)
	// ResourceQuotas returns an object that can list and get ResourceQuotas.
	ResourceQuotas(namespace string) ResourceQuotaNamespaceLister
	ResourceQuotaListerExpansion
}

// resourceQuotaLister implements the ResourceQuotaLister interface.
type resourceQuotaLister struct {
	listers.ResourceIndexer[*v1.ResourceQuota]
}

// NewResourceQuotaLister returns a new ResourceQuotaLister.
func NewResourceQuotaLister(indexer cache.Indexer) ResourceQuotaLister {
	return &resourceQuotaLister{listers.New[*v1.ResourceQuota](indexer, v1.Resource("resourcequota"))}
}

// ResourceQuotas returns an object that can list and get ResourceQuotas.
func (s *resourceQuotaLister) ResourceQuotas(namespace string) ResourceQuotaNamespaceLister {
	return resourceQuotaNamespaceLister{listers.NewNamespaced[*v1.ResourceQuota](s.ResourceIndexer, namespace)}
}

// ResourceQuotaNamespaceLister helps list and get ResourceQuotas.
// All objects returned here must be treated as read-only.
type ResourceQuotaNamespaceLister interface {
	// List lists all ResourceQuotas in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ResourceQuota, err error)
	// Get retrieves the ResourceQuota from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ResourceQuota, error)
	ResourceQuotaNamespaceListerExpansion
}

// resourceQuotaNamespaceLister implements the ResourceQuotaNamespaceLister
// interface.
type resourceQuotaNamespaceLister struct {
	listers.ResourceIndexer[*v1.ResourceQuota]
}
