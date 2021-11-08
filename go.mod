module github.com/prometheus-operator/prometheus-operator

go 1.17

require (
	github.com/alecthomas/units v0.0.0-20210927113745-59d0afb8317a
	github.com/blang/semver/v4 v4.0.0
	github.com/brancz/kube-rbac-proxy v0.8.0
	github.com/docker/distribution v2.7.1+incompatible
	github.com/evanphx/json-patch/v5 v5.6.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-kit/log v0.2.0
	github.com/go-openapi/swag v0.19.15
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-cmp v0.5.6
	github.com/hashicorp/go-version v1.3.0
	github.com/kylelemons/godebug v1.1.0
	github.com/mitchellh/hashstructure v1.1.0
	github.com/oklog/run v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus-community/prom-label-proxy v0.4.0
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.52.0
	github.com/prometheus-operator/prometheus-operator/pkg/client v0.52.0
	github.com/prometheus/alertmanager v0.23.0
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.32.1
	github.com/prometheus/prometheus v1.8.2-0.20211005150130-f29caccc4255
	github.com/stretchr/testify v1.7.0
	github.com/thanos-io/thanos v0.23.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/protobuf v1.27.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.22.3
	k8s.io/apiextensions-apiserver v0.22.3
	k8s.io/apimachinery v0.22.3
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/component-base v0.22.3
	k8s.io/klog/v2 v2.30.0
	k8s.io/utils v0.0.0-20210930125809-cb0fa318a74b
)

replace (
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring => ./pkg/apis/monitoring
	github.com/prometheus-operator/prometheus-operator/pkg/client => ./pkg/client
	// A replace directive is needed for k8s.io/client-go because Cortex (which
	// is an indirect dependency through Thanos) has a requirement on v12.0.0.
	k8s.io/client-go => k8s.io/client-go v0.22.3
	k8s.io/klog/v2 => github.com/simonpasquier/klog-gokit/v3 v3.0.0
)
