module runtime-test-plugin-0-11

go 1.16

replace (
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/helpers => ./../../compatibility-test-plugins/helpers
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework => ./../../framework
)

replace (
	sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v1.0.1
	sigs.k8s.io/kind => sigs.k8s.io/kind v0.11.1
)

require (
	github.com/spf13/cobra v1.6.1
	github.com/vmware-tanzu/tanzu-framework v0.11.6
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/helpers v0.0.0-00010101000000-000000000000
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1

)

require (
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.26.0
	sigs.k8s.io/controller-runtime v0.12.3 // indirect
)
