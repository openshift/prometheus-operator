// Copyright 2016 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sutil

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/cespare/xxhash/v2"
	promversion "github.com/prometheus/common/version"
	appsv1 "k8s.io/api/apps/v1"
	authv1 "k8s.io/api/authorization/v1"
	v1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/client-go/discovery"
	clientappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	clientauthv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	clientv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	clientdiscoveryv1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	monitoringv1beta1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1beta1"
)

// KubeConfigEnv (optionally) specify the location of kubeconfig file.
const KubeConfigEnv = "KUBECONFIG"

const StatusCleanupFinalizerName = "monitoring.coreos.com/status-cleanup"

var invalidDNS1123Characters = regexp.MustCompile("[^-a-z0-9]+")

var scheme = runtime.NewScheme()

func init() {
	utilruntime.Must(monitoringv1.SchemeBuilder.AddToScheme(scheme))
	utilruntime.Must(monitoringv1alpha1.SchemeBuilder.AddToScheme(scheme))
	utilruntime.Must(monitoringv1beta1.SchemeBuilder.AddToScheme(scheme))
}

// PodRunningAndReady returns whether a pod is running and each container has
// passed it's ready state.
func PodRunningAndReady(pod v1.Pod) (bool, error) {
	switch pod.Status.Phase {
	case v1.PodFailed, v1.PodSucceeded:
		return false, fmt.Errorf("pod completed with phase %s", pod.Status.Phase)
	case v1.PodRunning:
		for _, cond := range pod.Status.Conditions {
			if cond.Type != v1.PodReady {
				continue
			}
			return cond.Status == v1.ConditionTrue, nil
		}
		return false, fmt.Errorf("pod ready condition not found")
	}
	return false, nil
}

type ClusterConfig struct {
	Host           string
	TLSConfig      rest.TLSClientConfig
	AsUser         string
	KubeconfigPath string
}

func NewClusterConfig(config ClusterConfig) (*rest.Config, error) {
	var cfg *rest.Config
	var err error

	if config.KubeconfigPath == "" {
		config.KubeconfigPath = os.Getenv(KubeConfigEnv)
	}

	if config.KubeconfigPath != "" {
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		cfg, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{}).ClientConfig()
		if err != nil {
			return nil, fmt.Errorf("error creating config from %s: %w", config.KubeconfigPath, err)
		}
	} else {
		if len(config.Host) == 0 {
			if cfg, err = rest.InClusterConfig(); err != nil {
				return nil, err
			}
		} else {
			cfg = &rest.Config{
				Host: config.Host,
			}
			hostURL, err := url.Parse(config.Host)
			if err != nil {
				return nil, fmt.Errorf("error parsing host url %s: %w", config.Host, err)
			}
			if hostURL.Scheme == "https" {
				cfg.TLSClientConfig = config.TLSConfig
			}
		}
	}

	cfg.QPS = 100
	cfg.Burst = 100

	cfg.UserAgent = fmt.Sprintf("PrometheusOperator/%s", promversion.Version)
	cfg.Impersonate.UserName = config.AsUser

	return cfg, nil
}

// ResourceAttribute represents authorization attributes to check on a given resource.
type ResourceAttribute struct {
	Resource string
	Name     string
	Group    string
	Version  string
	Verbs    []string
}

// IsAllowed returns whether the user (e.g. the operator's service account) has
// been granted the required RBAC attributes.
// It returns true when the conditions are met for the namespaces (an empty
// namespace value means "all").
// The second return value returns the list of permissions that are missing if
// the requirements aren't met.
func IsAllowed(
	ctx context.Context,
	ssarClient clientauthv1.SelfSubjectAccessReviewInterface,
	namespaces []string,
	attributes ...ResourceAttribute,
) (bool, []error, error) {
	if len(attributes) == 0 {
		return false, nil, fmt.Errorf("resource attributes must not be empty")
	}

	if len(namespaces) == 0 {
		namespaces = []string{v1.NamespaceAll}
	}

	var missingPermissions []error
	for _, ns := range namespaces {
		for _, ra := range attributes {
			for _, verb := range ra.Verbs {
				resourceAttributes := authv1.ResourceAttributes{
					Verb:     verb,
					Group:    ra.Group,
					Version:  ra.Version,
					Resource: ra.Resource,
					// An empty name value means "all" resources.
					Name: ra.Name,
					// An empty namespace value means "all" for namespace-scoped resources.
					Namespace: ns,
				}

				// Special case for SAR on namespaces resources: Namespace and
				// Name need to be equal.
				if resourceAttributes.Group == "" && resourceAttributes.Resource == "namespaces" && resourceAttributes.Name != "" && resourceAttributes.Namespace == "" {
					resourceAttributes.Namespace = resourceAttributes.Name
				}

				ssar := &authv1.SelfSubjectAccessReview{
					Spec: authv1.SelfSubjectAccessReviewSpec{
						ResourceAttributes: &resourceAttributes,
					},
				}

				// FIXME(simonpasquier): retry in case of server-side errors.
				ssarResponse, err := ssarClient.Create(ctx, ssar, metav1.CreateOptions{})
				if err != nil {
					return false, nil, err
				}

				if !ssarResponse.Status.Allowed {
					var (
						reason   error
						resource = ra.Resource
					)
					if ra.Name != "" {
						resource += "/" + ra.Name
					}

					switch ns {
					case v1.NamespaceAll:
						reason = fmt.Errorf("missing %q permission on resource %q (group: %q) for all namespaces", verb, resource, ra.Group)
					default:
						reason = fmt.Errorf("missing %q permission on resource %q (group: %q) for namespace %q", verb, resource, ra.Group, ns)
					}

					missingPermissions = append(missingPermissions, reason)
				}
			}
		}
	}

	return len(missingPermissions) == 0, missingPermissions, nil
}

func IsResourceNotFoundError(err error) bool {
	se, ok := err.(*apierrors.StatusError)
	if !ok {
		return false
	}

	if se.Status().Code == http.StatusNotFound && se.Status().Reason == metav1.StatusReasonNotFound {
		return true
	}

	return false
}

func CreateOrUpdateService(ctx context.Context, sclient clientv1.ServiceInterface, svc *v1.Service) (*v1.Service, error) {
	var ret *v1.Service

	// As stated in the RetryOnConflict's documentation, the returned error shouldn't be wrapped.
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		service, err := sclient.Get(ctx, svc.Name, metav1.GetOptions{})
		if err != nil {
			if !apierrors.IsNotFound(err) {
				return err
			}

			ret, err = sclient.Create(ctx, svc, metav1.CreateOptions{})
			return err
		}

		// Apply immutable fields from the existing service.
		svc.Spec.IPFamilies = service.Spec.IPFamilies
		svc.Spec.IPFamilyPolicy = service.Spec.IPFamilyPolicy
		svc.Spec.ClusterIP = service.Spec.ClusterIP
		svc.Spec.ClusterIPs = service.Spec.ClusterIPs

		svc.SetOwnerReferences(mergeOwnerReferences(service.GetOwnerReferences(), svc.GetOwnerReferences()))
		mergeMetadata(&svc.ObjectMeta, service.ObjectMeta)

		ret, err = sclient.Update(ctx, svc, metav1.UpdateOptions{})
		return err
	})

	return ret, err
}

// CreateOrUpdateEndpoints creates or updates an endpoint resource.
//
//nolint:staticcheck // Ignore SA1019 Endpoints is marked as deprecated.
func CreateOrUpdateEndpoints(ctx context.Context, eclient clientv1.EndpointsInterface, eps *v1.Endpoints) error {
	// As stated in the RetryOnConflict's documentation, the returned error shouldn't be wrapped.
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		endpoints, err := eclient.Get(ctx, eps.Name, metav1.GetOptions{})
		if err != nil {
			if !apierrors.IsNotFound(err) {
				return err
			}

			_, err = eclient.Create(ctx, eps, metav1.CreateOptions{})
			return err
		}

		mergeMetadata(&eps.ObjectMeta, endpoints.ObjectMeta)

		_, err = eclient.Update(ctx, eps, metav1.UpdateOptions{})
		return err
	})
}

func CreateOrUpdateEndpointSlice(ctx context.Context, c clientdiscoveryv1.EndpointSliceInterface, eps *discoveryv1.EndpointSlice) error {
	// As stated in the RetryOnConflict's documentation, the returned error shouldn't be wrapped.
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		if eps.Name == "" {
			_, err := c.Create(ctx, eps, metav1.CreateOptions{})
			return err
		}

		endpoints, err := c.Get(ctx, eps.Name, metav1.GetOptions{})
		if err != nil {
			if !apierrors.IsNotFound(err) {
				return err
			}

			_, err = c.Create(ctx, eps, metav1.CreateOptions{})
			return err
		}

		mergeMetadata(&eps.ObjectMeta, endpoints.ObjectMeta)

		_, err = c.Update(ctx, eps, metav1.UpdateOptions{})
		return err
	})
}

// UpdateStatefulSet merges metadata of existing StatefulSet with new one and updates it.
func UpdateStatefulSet(ctx context.Context, sstClient clientappsv1.StatefulSetInterface, sset *appsv1.StatefulSet) error {
	// As stated in the RetryOnConflict's documentation, the returned error shouldn't be wrapped.
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		existingSset, err := sstClient.Get(ctx, sset.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}

		mergeMetadata(&sset.ObjectMeta, existingSset.ObjectMeta)
		// Propagate annotations set by kubectl on spec.template.annotations. e.g performing a rolling restart.
		mergeKubectlAnnotations(&existingSset.Spec.Template.ObjectMeta, sset.Spec.Template.ObjectMeta)

		_, err = sstClient.Update(ctx, sset, metav1.UpdateOptions{})
		return err
	})
}

// UpdateDaemonSet merges metadata of existing DaemonSet with new one and updates it.
func UpdateDaemonSet(ctx context.Context, dmsClient clientappsv1.DaemonSetInterface, dset *appsv1.DaemonSet) error {
	// As stated in the RetryOnConflict's documentation, the returned error shouldn't be wrapped.
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		existingDset, err := dmsClient.Get(ctx, dset.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}

		mergeMetadata(&dset.ObjectMeta, existingDset.ObjectMeta)
		// Propagate annotations set by kubectl on spec.template.annotations. e.g performing a rolling restart.
		mergeKubectlAnnotations(&existingDset.Spec.Template.ObjectMeta, dset.Spec.Template.ObjectMeta)

		_, err = dmsClient.Update(ctx, dset, metav1.UpdateOptions{})
		return err
	})
}

// CreateOrUpdateSecret merges metadata of existing Secret with new one and updates it.
func CreateOrUpdateSecret(ctx context.Context, secretClient clientv1.SecretInterface, desired *v1.Secret) error {
	// As stated in the RetryOnConflict's documentation, the returned error shouldn't be wrapped.
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		existingSecret, err := secretClient.Get(ctx, desired.Name, metav1.GetOptions{})
		if err != nil {
			if !apierrors.IsNotFound(err) {
				return err
			}

			_, err = secretClient.Create(ctx, desired, metav1.CreateOptions{})
			return err
		}

		mutated := existingSecret.DeepCopyObject().(*v1.Secret)
		mergeMetadata(&desired.ObjectMeta, mutated.ObjectMeta)
		if apiequality.Semantic.DeepEqual(existingSecret, desired) {
			return nil
		}
		_, err = secretClient.Update(ctx, desired, metav1.UpdateOptions{})
		return err
	})
}

// IsAPIGroupVersionResourceSupported checks if given groupVersion and resource is supported by the cluster.
func IsAPIGroupVersionResourceSupported(discoveryCli discovery.DiscoveryInterface, groupVersion schema.GroupVersion, resource string) (bool, error) {
	apiResourceList, err := discoveryCli.ServerResourcesForGroupVersion(groupVersion.String())
	if err != nil {
		if IsResourceNotFoundError(err) {
			return false, nil
		}

		return false, err
	}

	for _, apiResource := range apiResourceList.APIResources {
		if resource == apiResource.Name {
			return true, nil
		}
	}

	return false, nil
}

// ResourceNamer knows how to generate valid names for various Kubernetes resources.
type ResourceNamer struct {
	prefix string
}

// NewResourceNamerWithPrefix returns a ResourceNamer that adds a prefix
// followed by an hyphen character to all resource names.
func NewResourceNamerWithPrefix(p string) ResourceNamer {
	return ResourceNamer{prefix: p}
}

func (rn ResourceNamer) sanitizedLabel(name string) string {
	if rn.prefix != "" {
		name = strings.TrimRight(rn.prefix, "-") + "-" + name
	}

	name = strings.ToLower(name)
	name = invalidDNS1123Characters.ReplaceAllString(name, "-")
	name = strings.Trim(name, "-")

	return name
}

func isValidDNS1123Label(name string) error {
	if errs := validation.IsDNS1123Label(name); len(errs) > 0 {
		return errors.New(strings.Join(errs, ","))
	}

	return nil
}

// UniqueDNS1123Label returns a name that is a valid DNS-1123 label.
// The returned name has a hash-based suffix to ensure uniqueness in case the
// input name exceeds the 63-chars limit.
func (rn ResourceNamer) UniqueDNS1123Label(name string) (string, error) {
	// Hash the name and append the 8 first characters of the hash
	// value to the resulting name to ensure that 2 names longer than
	// DNS1123LabelMaxLength return unique names.
	// E.g. long-63-chars-abc, long-63-chars-XYZ may be added to
	// name since they are trimmed at long-63-chars, there will be 2
	// resource entries with the same name.
	// In practice, the hash is computed for the full name then trimmed to
	// the first 8 chars and added to the end:
	// * long-63-chars-abc -> first-54-chars-deadbeef
	// * long-63-chars-XYZ -> first-54-chars-d3adb33f
	xxh := xxhash.New()
	if _, err := xxh.Write([]byte(name)); err != nil {
		return "", err
	}

	h := fmt.Sprintf("-%x", xxh.Sum64())
	h = h[:9]

	name = rn.sanitizedLabel(name)

	if len(name) > validation.DNS1123LabelMaxLength-9 {
		name = name[:validation.DNS1123LabelMaxLength-9]
	}

	name = name + h
	if errs := validation.IsDNS1123Label(name); len(errs) > 0 {
		return "", errors.New(strings.Join(errs, ","))
	}

	return name, isValidDNS1123Label(name)
}

// DNS1123Label returns a name that is a valid DNS-1123 label.
// It will sanitize a name, removing invalid characters and if
// the name is bigger than 63 chars it will truncate it.
func (rn ResourceNamer) DNS1123Label(name string) (string, error) {
	name = rn.sanitizedLabel(name)

	if len(name) > validation.DNS1123LabelMaxLength {
		name = name[:validation.DNS1123LabelMaxLength]
	}

	return name, isValidDNS1123Label(name)
}

// AddTypeInformationToObject adds TypeMeta information to a runtime.Object based upon the loaded scheme.Scheme
// See https://github.com/kubernetes/client-go/issues/308#issuecomment-700099260
func AddTypeInformationToObject(obj runtime.Object) error {
	gvks, _, err := scheme.ObjectKinds(obj)
	if err != nil {
		return fmt.Errorf("missing apiVersion or kind and cannot assign it; %w", err)
	}

	for _, gvk := range gvks {
		if len(gvk.Kind) == 0 {
			continue
		}
		if len(gvk.Version) == 0 || gvk.Version == runtime.APIVersionInternal {
			continue
		}
		obj.GetObjectKind().SetGroupVersionKind(gvk)
		break
	}

	return nil
}

func mergeOwnerReferences(oldObj []metav1.OwnerReference, newObj []metav1.OwnerReference) []metav1.OwnerReference {
	existing := make(map[metav1.OwnerReference]bool)
	for _, ownerRef := range oldObj {
		existing[ownerRef] = true
	}
	for _, ownerRef := range newObj {
		if _, ok := existing[ownerRef]; !ok {
			oldObj = append(oldObj, ownerRef)
		}
	}
	return oldObj
}

// mergeMetadata takes labels and annotations from the old resource and merges
// them into the new resource. If a key is present in both resources, the new
// resource wins. It also copies the ResourceVersion from the old resource to
// the new resource to prevent update conflicts.
func mergeMetadata(newObj *metav1.ObjectMeta, oldObj metav1.ObjectMeta) {
	newObj.ResourceVersion = oldObj.ResourceVersion

	newObj.SetLabels(mergeMaps(newObj.Labels, oldObj.Labels))
	newObj.SetAnnotations(mergeMaps(newObj.Annotations, oldObj.Annotations))
}

func mergeMaps(newObj map[string]string, oldObj map[string]string) map[string]string {
	return mergeMapsByPrefix(newObj, oldObj, "")
}

func mergeKubectlAnnotations(from *metav1.ObjectMeta, to metav1.ObjectMeta) {
	from.SetAnnotations(mergeMapsByPrefix(from.Annotations, to.Annotations, "kubectl.kubernetes.io/"))
}

func mergeMapsByPrefix(from map[string]string, to map[string]string, prefix string) map[string]string {
	if to == nil {
		to = make(map[string]string)
	}

	if from == nil {
		from = make(map[string]string)
	}

	for k, v := range from {
		if strings.HasPrefix(k, prefix) {
			to[k] = v
		}
	}

	return to
}

func UpdateDNSConfig(podSpec *v1.PodSpec, config *monitoringv1.PodDNSConfig) {
	if config == nil {
		return
	}

	dnsConfig := v1.PodDNSConfig{
		Nameservers: config.Nameservers,
		Searches:    config.Searches,
	}

	for _, opt := range config.Options {
		dnsConfig.Options = append(dnsConfig.Options, v1.PodDNSConfigOption{
			Name:  opt.Name,
			Value: opt.Value,
		})
	}

	podSpec.DNSConfig = &dnsConfig
}

func UpdateDNSPolicy(podSpec *v1.PodSpec, dnsPolicy *monitoringv1.DNSPolicy) {
	if dnsPolicy == nil {
		return
	}

	podSpec.DNSPolicy = v1.DNSPolicy(*dnsPolicy)
}

// This function is responsible for the following:
//
// Verify that the service exists in the resource's namespace
// If it does not exist, fail the reconciliation.
//
// If the ServiceName is specified and a service with the same name exists in the same namespace as the
// resource, ensure that the custom governing service's selector matches the
// labels.
// If it is not selected, fail the reconciliation
// Warning: the function will panic if the resource's ServiceName is nil..
func EnsureCustomGoverningService(ctx context.Context, namespace string, serviceName string, svcClient clientv1.ServiceInterface, selectorLabels map[string]string) error {
	// Check if the custom governing service is defined in the same namespace and selects the Prometheus pod.
	svc, err := svcClient.Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get custom governing service %s/%s: %w", namespace, serviceName, err)
	}

	svcSelector, err := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{MatchLabels: svc.Spec.Selector})
	if err != nil {
		return fmt.Errorf("failed to parse the selector labels for custom governing service %s/%s: %w", namespace, serviceName, err)
	}

	if !svcSelector.Matches(labels.Set(selectorLabels)) {
		return fmt.Errorf("custom governing service %s/%s with selector %q does not select pods with labels %q",
			namespace, serviceName, svcSelector.String(), labels.Set(selectorLabels).String())
	}
	return nil
}

// AddFinalizerPatch generates the JSON patch payload which adds the finalizer to the object's metadata.
// If the finalizer is already present, it returns an empty []byte slice.
func FinalizerAddPatch(finalizers []string, finalizerName string) ([]byte, error) {
	if slices.Contains(finalizers, finalizerName) {
		return []byte{}, nil
	}
	if len(finalizers) == 0 {
		patch := []map[string]interface{}{
			{
				"op":    "add",
				"path":  "/metadata/finalizers",
				"value": []string{finalizerName},
			},
		}
		return json.Marshal(patch)
	}
	patch := []map[string]interface{}{
		{
			"op":    "add",
			"path":  "/metadata/finalizers/-",
			"value": finalizerName,
		},
	}
	return json.Marshal(patch)
}

// FinalizerDeletePatch generates the JSON patch payload which deletes the finalizer from the object's metadata.
// If the finalizer is not present, it returns nil.
func FinalizerDeletePatch(finalizers []string, finalizerName string) ([]byte, error) {
	for i, f := range finalizers {
		if f == finalizerName {
			patch := []map[string]interface{}{
				{
					"op":   "remove",
					"path": fmt.Sprintf("/metadata/finalizers/%d", i),
				},
			}
			return json.Marshal(patch)
		}
	}
	return nil, nil
}

func HasStatusCleanupFinalizer(obj metav1.Object) bool {
	return slices.Contains(obj.GetFinalizers(), StatusCleanupFinalizerName)
}
