# Configuration API

## Details

- CFG: All existing types are stored in the existing configuration file ( ~/.config/tanzu/config.yaml on most systems, shortened as CFG for the rest of the document) Also we introduce an additional next gen configuration file CFG_NG as well as a CFG Metadata file (shortened as META).

- CFG_NG: Additional configuration file(~/config/tanzu/config_ng.yaml) introduced in CFG_NG. CFG has to be retained since that is the file that Legacy Plugins rely on today for configuration state.

- In this proposal the overall state of the CLI configuration can conceptually be thought of as the union of two non-intersecting parts (CFG, CFG_NG) of the larger logical config file.

- META: Over time, how the CLI interacts with the configuration can potentially evolve. This evolution will be triggered via a core CLI update, but has to be communicated to the plugins as well. This file, META (config-metadata.yaml a hidden/dot file, ) will serve the role of capturing this information. The file is not expected to be manipulated by the CLI user.

- Aspects of this proposal that plans to leverage META are:

- Configuration of the patch policy for different parts of the configuration.

- Determining when to transition to using a single configuration file (CFG_NG) to persist configuration state

Note: due to the fact that information in the CFG_NG file directly affects the manipulation of the actual configuration files, it is not practical to embed said information in the files being manipulated

### Runtime Config APIs

VVV Update list, remove references to 'Server'

``` go
func GetServer(name string) error
func SetServer(server Server) error
func DeleteServer(serverName string) error
func GetCurrentServer() error
func SetCurrentServer(currentServer string) error
func DeleteCurrentServer() error

func GetContext(name string) error
func HasContext(name string) error
func SetContext(context Context) error
func DeleteContext(name string) error
func GetCurrentContext(contextType ContextType) error
func SetCurrentContext(context Context) error
func DeleteCurrentContext(contextType ContextType) error

func IsFeatureEnabled(plugin, feature string) bool, error
func DeleteFeature(plugin, feature string) error
func SetFeature(plugin, feature, value string) error
func ConfigureDefaultFeatureFlagsIfMissing(plugin string, defaultFeatureFlags map[string]bool) error

func GetEnv(key string) error
func SetEnv(key, value string) error
func DeleteEnv(key string) error

func GetEdition() error
func SetEdition(edition string) error

func GetCLIPluginDiscoverySources() error
func SetCLIPluginDiscoverySource(pluginDiscovery PluginDiscovery) error
func DeleteCLIPluginDiscovery(name string) error
```

#### How to use the Config APIs

- Import the runtime/config package and use the API method as specified below

##### Example: Add new context - The new context will be stored in CFG_NG

``` go
import configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"

err := configapi.SetContext(newContextObject, booleanToSetThisContextAsCurrent)
if err != nil{
 // Failed to add new context
 fmt.Println(err.Error())
}
```

Ex:- Add new Server - The new server will be stored in CFG.

``` go
import configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"

err := configapi.SetServer(newServerObject, booleanToSetThisServerAsCurrent)
if err != nil{
 // Failed to add new server
 fmt.Println(err.Error())
}
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
