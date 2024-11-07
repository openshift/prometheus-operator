// Copyright The prometheus-operator Authors
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ThanosRulerSpecApplyConfiguration represents a declarative configuration of the ThanosRulerSpec type for use
// with apply.
type ThanosRulerSpecApplyConfiguration struct {
	Version                            *string                                         `json:"version,omitempty"`
	PodMetadata                        *EmbeddedObjectMetadataApplyConfiguration       `json:"podMetadata,omitempty"`
	Image                              *string                                         `json:"image,omitempty"`
	ImagePullPolicy                    *corev1.PullPolicy                              `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets                   []corev1.LocalObjectReference                   `json:"imagePullSecrets,omitempty"`
	Paused                             *bool                                           `json:"paused,omitempty"`
	Replicas                           *int32                                          `json:"replicas,omitempty"`
	NodeSelector                       map[string]string                               `json:"nodeSelector,omitempty"`
	Resources                          *corev1.ResourceRequirements                    `json:"resources,omitempty"`
	Affinity                           *corev1.Affinity                                `json:"affinity,omitempty"`
	Tolerations                        []corev1.Toleration                             `json:"tolerations,omitempty"`
	TopologySpreadConstraints          []corev1.TopologySpreadConstraint               `json:"topologySpreadConstraints,omitempty"`
	SecurityContext                    *corev1.PodSecurityContext                      `json:"securityContext,omitempty"`
	DNSPolicy                          *monitoringv1.DNSPolicy                         `json:"dnsPolicy,omitempty"`
	DNSConfig                          *PodDNSConfigApplyConfiguration                 `json:"dnsConfig,omitempty"`
	PriorityClassName                  *string                                         `json:"priorityClassName,omitempty"`
	ServiceAccountName                 *string                                         `json:"serviceAccountName,omitempty"`
	Storage                            *StorageSpecApplyConfiguration                  `json:"storage,omitempty"`
	Volumes                            []corev1.Volume                                 `json:"volumes,omitempty"`
	VolumeMounts                       []corev1.VolumeMount                            `json:"volumeMounts,omitempty"`
	ObjectStorageConfig                *corev1.SecretKeySelector                       `json:"objectStorageConfig,omitempty"`
	ObjectStorageConfigFile            *string                                         `json:"objectStorageConfigFile,omitempty"`
	ListenLocal                        *bool                                           `json:"listenLocal,omitempty"`
	QueryEndpoints                     []string                                        `json:"queryEndpoints,omitempty"`
	QueryConfig                        *corev1.SecretKeySelector                       `json:"queryConfig,omitempty"`
	AlertManagersURL                   []string                                        `json:"alertmanagersUrl,omitempty"`
	AlertManagersConfig                *corev1.SecretKeySelector                       `json:"alertmanagersConfig,omitempty"`
	RuleSelector                       *metav1.LabelSelectorApplyConfiguration         `json:"ruleSelector,omitempty"`
	RuleNamespaceSelector              *metav1.LabelSelectorApplyConfiguration         `json:"ruleNamespaceSelector,omitempty"`
	EnforcedNamespaceLabel             *string                                         `json:"enforcedNamespaceLabel,omitempty"`
	ExcludedFromEnforcement            []ObjectReferenceApplyConfiguration             `json:"excludedFromEnforcement,omitempty"`
	PrometheusRulesExcludedFromEnforce []PrometheusRuleExcludeConfigApplyConfiguration `json:"prometheusRulesExcludedFromEnforce,omitempty"`
	LogLevel                           *string                                         `json:"logLevel,omitempty"`
	LogFormat                          *string                                         `json:"logFormat,omitempty"`
	PortName                           *string                                         `json:"portName,omitempty"`
	EvaluationInterval                 *monitoringv1.Duration                          `json:"evaluationInterval,omitempty"`
	Retention                          *monitoringv1.Duration                          `json:"retention,omitempty"`
	Containers                         []corev1.Container                              `json:"containers,omitempty"`
	InitContainers                     []corev1.Container                              `json:"initContainers,omitempty"`
	TracingConfig                      *corev1.SecretKeySelector                       `json:"tracingConfig,omitempty"`
	TracingConfigFile                  *string                                         `json:"tracingConfigFile,omitempty"`
	Labels                             map[string]string                               `json:"labels,omitempty"`
	AlertDropLabels                    []string                                        `json:"alertDropLabels,omitempty"`
	ExternalPrefix                     *string                                         `json:"externalPrefix,omitempty"`
	RoutePrefix                        *string                                         `json:"routePrefix,omitempty"`
	GRPCServerTLSConfig                *TLSConfigApplyConfiguration                    `json:"grpcServerTlsConfig,omitempty"`
	AlertQueryURL                      *string                                         `json:"alertQueryUrl,omitempty"`
	MinReadySeconds                    *uint32                                         `json:"minReadySeconds,omitempty"`
	AlertRelabelConfigs                *corev1.SecretKeySelector                       `json:"alertRelabelConfigs,omitempty"`
	AlertRelabelConfigFile             *string                                         `json:"alertRelabelConfigFile,omitempty"`
	HostAliases                        []HostAliasApplyConfiguration                   `json:"hostAliases,omitempty"`
	AdditionalArgs                     []ArgumentApplyConfiguration                    `json:"additionalArgs,omitempty"`
	Web                                *ThanosRulerWebSpecApplyConfiguration           `json:"web,omitempty"`
}

// ThanosRulerSpecApplyConfiguration constructs a declarative configuration of the ThanosRulerSpec type for use with
// apply.
func ThanosRulerSpec() *ThanosRulerSpecApplyConfiguration {
	return &ThanosRulerSpecApplyConfiguration{}
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithVersion(value string) *ThanosRulerSpecApplyConfiguration {
	b.Version = &value
	return b
}

// WithPodMetadata sets the PodMetadata field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodMetadata field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithPodMetadata(value *EmbeddedObjectMetadataApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	b.PodMetadata = value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithImage(value string) *ThanosRulerSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithImagePullPolicy sets the ImagePullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePullPolicy field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithImagePullPolicy(value corev1.PullPolicy) *ThanosRulerSpecApplyConfiguration {
	b.ImagePullPolicy = &value
	return b
}

// WithImagePullSecrets adds the given value to the ImagePullSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImagePullSecrets field.
func (b *ThanosRulerSpecApplyConfiguration) WithImagePullSecrets(values ...corev1.LocalObjectReference) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.ImagePullSecrets = append(b.ImagePullSecrets, values[i])
	}
	return b
}

// WithPaused sets the Paused field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Paused field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithPaused(value bool) *ThanosRulerSpecApplyConfiguration {
	b.Paused = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithReplicas(value int32) *ThanosRulerSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *ThanosRulerSpecApplyConfiguration) WithNodeSelector(entries map[string]string) *ThanosRulerSpecApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithResources(value corev1.ResourceRequirements) *ThanosRulerSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithAffinity(value corev1.Affinity) *ThanosRulerSpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *ThanosRulerSpecApplyConfiguration) WithTolerations(values ...corev1.Toleration) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithTopologySpreadConstraints adds the given value to the TopologySpreadConstraints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TopologySpreadConstraints field.
func (b *ThanosRulerSpecApplyConfiguration) WithTopologySpreadConstraints(values ...corev1.TopologySpreadConstraint) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.TopologySpreadConstraints = append(b.TopologySpreadConstraints, values[i])
	}
	return b
}

// WithSecurityContext sets the SecurityContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityContext field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithSecurityContext(value corev1.PodSecurityContext) *ThanosRulerSpecApplyConfiguration {
	b.SecurityContext = &value
	return b
}

// WithDNSPolicy sets the DNSPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSPolicy field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithDNSPolicy(value monitoringv1.DNSPolicy) *ThanosRulerSpecApplyConfiguration {
	b.DNSPolicy = &value
	return b
}

// WithDNSConfig sets the DNSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSConfig field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithDNSConfig(value *PodDNSConfigApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	b.DNSConfig = value
	return b
}

// WithPriorityClassName sets the PriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PriorityClassName field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithPriorityClassName(value string) *ThanosRulerSpecApplyConfiguration {
	b.PriorityClassName = &value
	return b
}

// WithServiceAccountName sets the ServiceAccountName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountName field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithServiceAccountName(value string) *ThanosRulerSpecApplyConfiguration {
	b.ServiceAccountName = &value
	return b
}

// WithStorage sets the Storage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Storage field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithStorage(value *StorageSpecApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	b.Storage = value
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *ThanosRulerSpecApplyConfiguration) WithVolumes(values ...corev1.Volume) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.Volumes = append(b.Volumes, values[i])
	}
	return b
}

// WithVolumeMounts adds the given value to the VolumeMounts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the VolumeMounts field.
func (b *ThanosRulerSpecApplyConfiguration) WithVolumeMounts(values ...corev1.VolumeMount) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.VolumeMounts = append(b.VolumeMounts, values[i])
	}
	return b
}

// WithObjectStorageConfig sets the ObjectStorageConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObjectStorageConfig field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithObjectStorageConfig(value corev1.SecretKeySelector) *ThanosRulerSpecApplyConfiguration {
	b.ObjectStorageConfig = &value
	return b
}

// WithObjectStorageConfigFile sets the ObjectStorageConfigFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObjectStorageConfigFile field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithObjectStorageConfigFile(value string) *ThanosRulerSpecApplyConfiguration {
	b.ObjectStorageConfigFile = &value
	return b
}

// WithListenLocal sets the ListenLocal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ListenLocal field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithListenLocal(value bool) *ThanosRulerSpecApplyConfiguration {
	b.ListenLocal = &value
	return b
}

// WithQueryEndpoints adds the given value to the QueryEndpoints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the QueryEndpoints field.
func (b *ThanosRulerSpecApplyConfiguration) WithQueryEndpoints(values ...string) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.QueryEndpoints = append(b.QueryEndpoints, values[i])
	}
	return b
}

// WithQueryConfig sets the QueryConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the QueryConfig field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithQueryConfig(value corev1.SecretKeySelector) *ThanosRulerSpecApplyConfiguration {
	b.QueryConfig = &value
	return b
}

// WithAlertManagersURL adds the given value to the AlertManagersURL field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AlertManagersURL field.
func (b *ThanosRulerSpecApplyConfiguration) WithAlertManagersURL(values ...string) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.AlertManagersURL = append(b.AlertManagersURL, values[i])
	}
	return b
}

// WithAlertManagersConfig sets the AlertManagersConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AlertManagersConfig field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithAlertManagersConfig(value corev1.SecretKeySelector) *ThanosRulerSpecApplyConfiguration {
	b.AlertManagersConfig = &value
	return b
}

// WithRuleSelector sets the RuleSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RuleSelector field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithRuleSelector(value *metav1.LabelSelectorApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	b.RuleSelector = value
	return b
}

// WithRuleNamespaceSelector sets the RuleNamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RuleNamespaceSelector field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithRuleNamespaceSelector(value *metav1.LabelSelectorApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	b.RuleNamespaceSelector = value
	return b
}

// WithEnforcedNamespaceLabel sets the EnforcedNamespaceLabel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnforcedNamespaceLabel field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithEnforcedNamespaceLabel(value string) *ThanosRulerSpecApplyConfiguration {
	b.EnforcedNamespaceLabel = &value
	return b
}

// WithExcludedFromEnforcement adds the given value to the ExcludedFromEnforcement field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExcludedFromEnforcement field.
func (b *ThanosRulerSpecApplyConfiguration) WithExcludedFromEnforcement(values ...*ObjectReferenceApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExcludedFromEnforcement")
		}
		b.ExcludedFromEnforcement = append(b.ExcludedFromEnforcement, *values[i])
	}
	return b
}

// WithPrometheusRulesExcludedFromEnforce adds the given value to the PrometheusRulesExcludedFromEnforce field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PrometheusRulesExcludedFromEnforce field.
func (b *ThanosRulerSpecApplyConfiguration) WithPrometheusRulesExcludedFromEnforce(values ...*PrometheusRuleExcludeConfigApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPrometheusRulesExcludedFromEnforce")
		}
		b.PrometheusRulesExcludedFromEnforce = append(b.PrometheusRulesExcludedFromEnforce, *values[i])
	}
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithLogLevel(value string) *ThanosRulerSpecApplyConfiguration {
	b.LogLevel = &value
	return b
}

// WithLogFormat sets the LogFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogFormat field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithLogFormat(value string) *ThanosRulerSpecApplyConfiguration {
	b.LogFormat = &value
	return b
}

// WithPortName sets the PortName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortName field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithPortName(value string) *ThanosRulerSpecApplyConfiguration {
	b.PortName = &value
	return b
}

// WithEvaluationInterval sets the EvaluationInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EvaluationInterval field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithEvaluationInterval(value monitoringv1.Duration) *ThanosRulerSpecApplyConfiguration {
	b.EvaluationInterval = &value
	return b
}

// WithRetention sets the Retention field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Retention field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithRetention(value monitoringv1.Duration) *ThanosRulerSpecApplyConfiguration {
	b.Retention = &value
	return b
}

// WithContainers adds the given value to the Containers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Containers field.
func (b *ThanosRulerSpecApplyConfiguration) WithContainers(values ...corev1.Container) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.Containers = append(b.Containers, values[i])
	}
	return b
}

// WithInitContainers adds the given value to the InitContainers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InitContainers field.
func (b *ThanosRulerSpecApplyConfiguration) WithInitContainers(values ...corev1.Container) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.InitContainers = append(b.InitContainers, values[i])
	}
	return b
}

// WithTracingConfig sets the TracingConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TracingConfig field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithTracingConfig(value corev1.SecretKeySelector) *ThanosRulerSpecApplyConfiguration {
	b.TracingConfig = &value
	return b
}

// WithTracingConfigFile sets the TracingConfigFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TracingConfigFile field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithTracingConfigFile(value string) *ThanosRulerSpecApplyConfiguration {
	b.TracingConfigFile = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ThanosRulerSpecApplyConfiguration) WithLabels(entries map[string]string) *ThanosRulerSpecApplyConfiguration {
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAlertDropLabels adds the given value to the AlertDropLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AlertDropLabels field.
func (b *ThanosRulerSpecApplyConfiguration) WithAlertDropLabels(values ...string) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		b.AlertDropLabels = append(b.AlertDropLabels, values[i])
	}
	return b
}

// WithExternalPrefix sets the ExternalPrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalPrefix field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithExternalPrefix(value string) *ThanosRulerSpecApplyConfiguration {
	b.ExternalPrefix = &value
	return b
}

// WithRoutePrefix sets the RoutePrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RoutePrefix field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithRoutePrefix(value string) *ThanosRulerSpecApplyConfiguration {
	b.RoutePrefix = &value
	return b
}

// WithGRPCServerTLSConfig sets the GRPCServerTLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GRPCServerTLSConfig field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithGRPCServerTLSConfig(value *TLSConfigApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	b.GRPCServerTLSConfig = value
	return b
}

// WithAlertQueryURL sets the AlertQueryURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AlertQueryURL field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithAlertQueryURL(value string) *ThanosRulerSpecApplyConfiguration {
	b.AlertQueryURL = &value
	return b
}

// WithMinReadySeconds sets the MinReadySeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReadySeconds field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithMinReadySeconds(value uint32) *ThanosRulerSpecApplyConfiguration {
	b.MinReadySeconds = &value
	return b
}

// WithAlertRelabelConfigs sets the AlertRelabelConfigs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AlertRelabelConfigs field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithAlertRelabelConfigs(value corev1.SecretKeySelector) *ThanosRulerSpecApplyConfiguration {
	b.AlertRelabelConfigs = &value
	return b
}

// WithAlertRelabelConfigFile sets the AlertRelabelConfigFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AlertRelabelConfigFile field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithAlertRelabelConfigFile(value string) *ThanosRulerSpecApplyConfiguration {
	b.AlertRelabelConfigFile = &value
	return b
}

// WithHostAliases adds the given value to the HostAliases field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HostAliases field.
func (b *ThanosRulerSpecApplyConfiguration) WithHostAliases(values ...*HostAliasApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHostAliases")
		}
		b.HostAliases = append(b.HostAliases, *values[i])
	}
	return b
}

// WithAdditionalArgs adds the given value to the AdditionalArgs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdditionalArgs field.
func (b *ThanosRulerSpecApplyConfiguration) WithAdditionalArgs(values ...*ArgumentApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAdditionalArgs")
		}
		b.AdditionalArgs = append(b.AdditionalArgs, *values[i])
	}
	return b
}

// WithWeb sets the Web field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Web field is set to the value of the last call.
func (b *ThanosRulerSpecApplyConfiguration) WithWeb(value *ThanosRulerWebSpecApplyConfiguration) *ThanosRulerSpecApplyConfiguration {
	b.Web = value
	return b
}
