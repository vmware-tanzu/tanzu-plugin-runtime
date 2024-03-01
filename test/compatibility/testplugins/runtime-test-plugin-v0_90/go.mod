module github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/testplugins/runtime-test-plugin-v0_90

go 1.18

replace github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core => ../../core

require (
	github.com/spf13/cobra v1.7.0
	github.com/vmware-tanzu/tanzu-plugin-runtime v0.90.0
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core v0.90.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/juju/fslock v0.0.0-20160525022230-4d5c94c67b4b // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
