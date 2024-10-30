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

package v1beta1

import (
	"context"

	v1beta1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	batchv1beta1 "k8s.io/client-go/applyconfigurations/batch/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CronJobsGetter has a method to return a CronJobInterface.
// A group's client should implement this interface.
type CronJobsGetter interface {
	CronJobs(namespace string) CronJobInterface
}

// CronJobInterface has methods to work with CronJob resources.
type CronJobInterface interface {
	Create(ctx context.Context, cronJob *v1beta1.CronJob, opts v1.CreateOptions) (*v1beta1.CronJob, error)
	Update(ctx context.Context, cronJob *v1beta1.CronJob, opts v1.UpdateOptions) (*v1beta1.CronJob, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, cronJob *v1beta1.CronJob, opts v1.UpdateOptions) (*v1beta1.CronJob, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.CronJob, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.CronJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CronJob, err error)
	Apply(ctx context.Context, cronJob *batchv1beta1.CronJobApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.CronJob, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, cronJob *batchv1beta1.CronJobApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.CronJob, err error)
	CronJobExpansion
}

// cronJobs implements CronJobInterface
type cronJobs struct {
	*gentype.ClientWithListAndApply[*v1beta1.CronJob, *v1beta1.CronJobList, *batchv1beta1.CronJobApplyConfiguration]
}

// newCronJobs returns a CronJobs
func newCronJobs(c *BatchV1beta1Client, namespace string) *cronJobs {
	return &cronJobs{
		gentype.NewClientWithListAndApply[*v1beta1.CronJob, *v1beta1.CronJobList, *batchv1beta1.CronJobApplyConfiguration](
			"cronjobs",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.CronJob { return &v1beta1.CronJob{} },
			func() *v1beta1.CronJobList { return &v1beta1.CronJobList{} },
			gentype.PrefersProtobuf[*v1beta1.CronJob](),
		),
	}
}
