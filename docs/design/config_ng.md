# Config API

The format of the Tanzu CLI configuration has evolved over time, with different generations of core CLI and plugins having possibly different expectations on what it contains.
There is a lack of clear specification on the compatibility guarantee each revision of the configuration provides to existing and future consumers.
This has led to modifications by some consumers that rendered the configuration unusable to others.
Tanzu CLI Runtime expect plugin developers to publish their plugins at different release cadences, and expect a generation of these plugins to interoperate for a substantial period of time.
During this period we should allow for flexibility in introducing compatible changes to the configuration definition with the least amount of impact to developers.
Access to the configuration file by the plugin should be done through APIs provided by the plugin runtime library. Similarly, access to the same configuration by the CLI user are through built-in commands 'tanzu config….' provided by the core CLI, which internally uses the APIs provided by plugin runtime library.

## Config API Rules

Changes within the same major version of configuration definition are backward compatible with older versions sharing the same major version, by adhering to the following rules:

- Existing fields should never be removed.
- Semantics of existing fields defined should not change.
- New fields can be introduced, as long as their presence in the config file is never assumed. This means each field can either be optional or, if mandatory, is initialized an explicit value or a default one before use.

All plugins built to interact with the same major version should be able to interoperate without fear of corrupting the CLI configuration for one another. To achieve this, modifications to the configuration have to be done in a safe manner, as described below.

The new implementation supports unambiguous and safe create/update/delete operations of different parts of the CFG and CFG_NG.

Provided SetSomeConfigSection(), GetSomeConfigSection(), DeleteSomeConfigSection() functions in the config interface to C/U/D a specific stanza of the CFG and CFG_NG.

Each SetSomeConfigSection(), GetSomeConfigSection(), DeleteSomeConfigSection() function only modifies the parts of the stanza it knows about, and leave everything else untouched.

YAML modification are done as non-disruptively as possible. In particular order of sections, elements in list are maintained
and comments associated with YAML nodes are retained.

## Granularity of Config Mutation

- Nondestructive modifications of config file via StoreClientConfig.
- Provided more explicit Config API that satisfies compatibility guarantee.

### Replace Strategy

Introduced new META to configure the metadata related info for config handling.

Example

``` yaml
configMetadata:
  patchStrategy:
    contexts.clusterOpts.annotation: "replace"
```

To the above metadata the Merge on setting the contexts will replace/delete the `clusterOpts.annotation` field if not specified in the change object.

By Default Set APIs will use Merge strategy on all the yaml nodes.

### Multi File Approach

Introduce Multi Config files CFG and CFG_NG to manage client config data.

Introduce new APIs to META to manage what goes into CFG_NG.

contexts and currentContexts will be moved to CFG_NG.

New type settings has been added to META to store any internal settings related to config operations which are controlled by core cli.

`useUnifiedConfig`  Used to move all the config related operations to CFG_NG. By default, this setting is deactivated and will be enabled in next future versions of cli.

Existing getClientConfigNode and getClientConfigNodeNoLock  function implementation is updated to handle multiple config files.

New factory and filesytem files are added to perform operations on CFG_NG.

TANZU_CONFIG_NEXT_GEN Env key is defined to explicitly set CFG_NG file.

#### Read from Config

- Check for useUnifiedConfig settings in META.
- if enabled Read `CFG_NG` and construct yaml node.
- If deactivated go to 2.
- Construct root node with cfgItems from CFG and next gen nodes from CFG_NG.
- Return the multi config root node.

#### Write to Config

- Check for useUnifiedConfig settings in META.
- if enabled Write to CFG_NG.
- If deactivated go to 2.
- Create root cfg node and root next gen node.
- Process through the change node and construct the root cfg node and root next gen node based on cfgItems.
- Write the root cfg node to CFG.
- Write the root next gen node to CFG_NG.
- Also populate the legacyClientConfig directory .tanzu.

Example:- Default behaviour

``` yaml
configMetadata:
  settings:
    useUnifiedConfig: false
```

Outcome :- Only contexts and currentContexts read and write to CFG_NG

`tanzu config get` cmd implementation is changed to handle multi files. now it will retrieve the combination of both CFG and CFG_NG

Example :-

``` yaml
configMetadata:
    settings:
        useUnifiedConfig: true
```

Outcome :- All config item read and write to CFG_NG

contexts and current context are now stored in new CFG_NG.

### Merge vs Replace mutations

Below is an example of

Starting state of configuration, (possibly written by older CLIs)

``` yaml
discoverySources:
- oci:
    image: staging.v.com/standalone-tkg-plugins:latest
    name: standalone
```

New client update source entry, new fields are introduced

``` yaml
discoverySources:
- oci:
    image: projects.v.com/standalone-tkg-plugins:latest
    name: standalone
    caCert: XXX
```

Older client (which is only aware of the image/name fields) update the standalone source to staging:

(merge)

``` yaml
discoverySources:
  - oci:
      image: staging.v.com/standalone-tkg-plugins:latest
      name: standalone
      caCert: XXX
```

(replace)

``` yaml
discoverySources:
 - oci:
    image: staging..com/standalone-tkg-plugins:latest
    name: standalone
```

(the list item is replaced, with the input, wiping out caCert field)

When encountering unknown fields during modification of a configuration entry, the default behavior of the new config writing logic is to leave them unperturbed (the merge behavior). An alternative implementation could be to remove those fields (essentially replace the entire config resource with the provided input).
Which update behavior makes sense is situation dependent, and may also depend on how the definition of a configuration resource evolves over time.

To provide the flexibility to employ either of the two behaviors, both the merge and replace capabilities will be implemented in the plugin runtime, with the replace logic only activated on specific sections of the CLI configuration as designated by the patch policy configuration in META
