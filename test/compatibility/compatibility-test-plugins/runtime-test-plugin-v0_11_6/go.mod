module github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/runtime-test-plugin-v0_11_6

go 1.16

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/k14s/kbld => github.com/anujc25/carvel-kbld v0.31.0-update-vendir
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core => ./../../core
	sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v1.0.1
	sigs.k8s.io/kind => sigs.k8s.io/kind v0.11.1
)

require (
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.16.0
	github.com/spf13/cobra v1.6.1
	github.com/vmware-tanzu/tanzu-framework v0.11.6
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)
