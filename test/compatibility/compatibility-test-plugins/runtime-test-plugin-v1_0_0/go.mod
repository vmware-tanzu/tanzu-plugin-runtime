module github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/compatibility-test-plugins/runtime-test-plugin-v1_0_0

go 1.18

replace github.com/vmware-tanzu/tanzu-plugin-runtime => ./../../../../../tanzu-plugin-runtime

replace github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core => ../../core

replace github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/framework => ../../framework

require (
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.27.0
	github.com/spf13/cobra v1.6.1
	github.com/vmware-tanzu/tanzu-plugin-runtime v0.0.0-00010101000000-000000000000
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/aunum/log v0.0.0-20200821225356-38d2e2c8b489 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/juju/fslock v0.0.0-20160525022230-4d5c94c67b4b // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
