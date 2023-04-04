# Configuration API

## Overview

Config library provides various API methods to store DiscoverySources, CLI, Features and environment settings in config.yaml (CFG) and Contexts in config-ng.yaml(CFG_NG) configuration files respectively.

## Details

- CFG: All existing types are stored in the existing configuration file ( ~/.config/tanzu/config.yaml on most systems, shortened as CFG for the rest of the document) Also we introduce an additional next gen configuration file CFG_NG as well as a CFG Metadata file (shortened as META).

- CFG_NG: Additional configuration file(~/config/tanzu/config_ng.yaml) introduced to store new configuration related types. CFG has to be retained since that is the file that Legacy Plugins rely on today for configuration state.

- In this proposal the overall state of the CLI configuration can conceptually be thought of as the union of two non-intersecting parts (CFG, CFG_NG) of the larger logical config file.

- META: Over time, how the CLI interacts with the configuration can potentially evolve. This evolution will be triggered via a core CLI update, but has to be communicated to the plugins as well. This file, META (config-metadata.yaml a hidden/dot file, ) will serve the role of capturing this information. The file is not expected to be manipulated by the CLI user.

- Aspects of this proposal that plans to leverage META are:

- Configuration of the patch policy for different parts of the configuration.

- Determining when to transition to using a single configuration file (CFG_NG) to persist configuration state

Note: due to the fact that information in the CFG_NG file directly affects the manipulation of the actual configuration files, it is not practical to embed said information in the files being manipulated

### Available Runtime Config APIs

``` go
// Context APIs
func GetContext(name string) (context Context, error)
func AddContext(context Context, setCurrent bool) error
func SetContext(context Context, setCurrent bool) error
func DeleteContext(name string) error
func RemoveContext(name string) error
func ContextExists(name string) (bool, error)
func GetCurrentContext(contextType ContextType) error
func GetAllCurrentContextsMap() (map[configtypes.Target]*configtypes.Context, error)
func GetAllCurrentContextsList() ([]string, error)
func SetCurrentContext(context Context) error
func RemoveCurrentContext(contextType ContextType) error
func EndpointFromContext(s *configtypes.Context) (endpoint string, err error)

// Feature APIs
func IsFeatureEnabled(plugin, key string) (bool, error)
func DeleteFeature(plugin, key string) error
func SetFeature(plugin, key, value string) error
func ConfigureDefaultFeatureFlagsIfMissing(plugin string, defaultFeatureFlags map[string]bool) error
func IsFeatureActivated(feature string) bool

// Env APIs
func GetAllEnvs() (map[string]string, error)
func GetEnv(key string) (string, error)
func SetEnv(key, value string) error
func DeleteEnv(key string) error
func GetEnvConfigurations() map[string]string

// Edition APIs
func GetEdition() (string, error)
func SetEdition(val string) (err error)

// Discovery Sources APIs
func GetCLIDiscoverySources() ([]configtypes.PluginDiscovery, error)
func GetCLIDiscoverySource(name string) (*configtypes.PluginDiscovery, error)
func SetCLIDiscoverySources(discoverySources []configtypes.PluginDiscovery) error
func SetCLIDiscoverySource(discoverySource configtypes.PluginDiscovery) error
func DeleteCLIDiscoverySource(name string) error

// ClientConfig APIs
func ClientConfigPath() (path string, err error)
func ClientConfigNextGenPath() (path string, err error)
func AcquireTanzuConfigNextGenLock()
func ReleaseTanzuConfigNextGenLock()
func AcquireTanzuConfigLock()
func ReleaseTanzuConfigLock()
func LocalDir() (path string, err error)
func DeleteClientConfigNextGen() error

// Config Metadata APIs
func GetMetadata() (*configtypes.Metadata, error)
func GetConfigMetadata() (*configtypes.ConfigMetadata, error)
func GetConfigMetadataPatchStrategy() (map[string]string, error)
func SetConfigMetadataPatchStrategy(key, value string) error
func SetConfigMetadataPatchStrategies(patchStrategies map[string]string) error
func CfgMetadataFilePath() (path string, err error)
func AcquireTanzuMetadataLock()
func ReleaseTanzuMetadataLock()

// Config Metadata Settings APIs
func GetConfigMetadataSettings() (map[string]string, error)
func GetConfigMetadataSetting(key string) (string, error)
func IsConfigMetadataSettingsEnabled(key string) (bool, error)
func UseUnifiedConfig() (bool, error)
func DeleteConfigMetadataSetting(key string) error
func SetConfigMetadataSetting(key, value string) error
```

#### How to use the Config APIs

- Import the runtime/config package and use the API method as specified below

##### Example: Add new context - The new context will be stored in CFG_NG

If a plugin wants to add the context it should use the SetContext API from
[context-related APIs](https://github.com/vmware-tanzu/tanzu-plugin-runtime/blob/main/config/contexts.go)
in the Tanzu Plugin Runtime library to ensure forward compatibility.
example, to add a context use the below snippet:

``` go
import configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"

err := configapi.SetContext(newContextObject, booleanToSetThisContextAsCurrent)
if err != nil{
 // Failed to add new context
 fmt.Println(err.Error())
}

// Fetch the context from the CFG_NG by context name
context, err := configapi.GetContext(name)

// Delete existing context from CFG_NG by context name
err := configapi.DeleteContext(name)
```

##### Example: Get an existing context - The matching context that is stored in CFG_NG

If a plugin wants to get the context it should use the GetContext API from
[context-related APIs](https://github.com/vmware-tanzu/tanzu-plugin-runtime/blob/main/config/contexts.go)
in the Tanzu Plugin Runtime library to ensure forward compatibility.
example, to get the context use the below snippet:

``` go
import configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"

// Fetch the context from the CFG_NG by context name
context, err := configapi.GetContext(name)
```

##### Example: Delete context - delete an existing context from CFG_NG

If a plugin wants to delete a context it should use the DeleteContext API from
[context-related APIs](https://github.com/vmware-tanzu/tanzu-plugin-runtime/blob/main/config/contexts.go)
in the Tanzu Plugin Runtime library to ensure forward compatibility.
example, to delete the context use the below snippet:

``` go
import configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"

// Delete existing context from CFG_NG by context name
err := configapi.DeleteContext(name)
```

##### Example: Retrieve context information for a specific target

If a plugin wants to access the context it should use the
[context-related APIs](https://github.com/vmware-tanzu/tanzu-plugin-runtime/blob/main/config/contexts.go)
in the Tanzu Plugin Runtime library to ensure forward compatibility. For
example, to get the current active context use the below snippet:

```go
import (
  config "github.com/vmware-tanzu/tanzu-plugin-runtime/config"
  cfgtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

ctx, err := config.GetCurrentContext(cfgtypes.TargetK8s)
