## tanzu package install

Install a package

```
tanzu package install INSTALLED_PACKAGE_NAME --package-name PACKAGE_NAME --version VERSION [flags]
```

### Examples

```

    # Install package contour with installed package name as 'contour-pkg' with specified version and without waiting for package reconciliation to complete 	
    tanzu package install contour-pkg --package-name contour.tanzu.vmware.com --namespace test-ns --version 1.15.1-tkg.1-vmware1 --wait=false
	
    # Install package contour with kubeconfig flag and waiting for package reconciliation to complete	
    tanzu package install contour-pkg --package-name contour.tanzu.vmware.com --namespace test-ns --version 1.15.1-tkg.1-vmware1 --kubeconfig path/to/kubeconfig
```

### Options

```
      --create-namespace              Create namespace if the target namespace does not exist, optional
  -h, --help                          help for install
  -n, --namespace string              Namespace indicates the location of the repository from which the package is retrieved (default "default")
  -p, --package-name string           Name of the package to be installed
      --poll-interval duration        Time interval between subsequent polls of package reconciliation status, optional (default 1s)
      --poll-timeout duration         Timeout value for polls of package reconciliation status, optional (default 15m0s)
      --service-account-name string   Name of an existing service account used to install underlying package contents, optional
  -f, --values-file string            The path to the configuration values file, optional
  -v, --version string                Version of the package to be installed
      --wait                          Wait for the package reconciliation to complete, optional. To disable wait, specify --wait=false (default true)
```

### Options inherited from parent commands

```
      --kubeconfig string   The path to the kubeconfig file, optional
      --verbose int32       Number for the log level verbosity(0-9)
```

### SEE ALSO

* [tanzu package](tanzu_package.md)	 - Tanzu package management

###### Auto generated by spf13/cobra on 14-Sep-2022