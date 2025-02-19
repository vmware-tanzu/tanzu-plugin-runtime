module github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/testplugins/runtime-test-plugin-latest

go 1.22.0

toolchain go1.23.1

replace github.com/vmware-tanzu/tanzu-plugin-runtime => ./../../../../../tanzu-plugin-runtime

replace github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core => ../../core

require (
	github.com/spf13/cobra v1.9.1
	// Replaced above by latest version
	github.com/vmware-tanzu/tanzu-plugin-runtime v0.0.0-00010101000000-000000000000
	github.com/vmware-tanzu/tanzu-plugin-runtime/test/compatibility/core v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/alexflint/go-filemutex v1.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
)
