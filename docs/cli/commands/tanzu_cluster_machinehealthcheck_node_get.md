## tanzu cluster machinehealthcheck node get

Get a MachineHealthCheck object of the nodes of a cluster

### Synopsis

Get a MachineHealthCheck object of the nodes for the given cluster

```
tanzu cluster machinehealthcheck node get CLUSTER_NAME [flags]
```

### Options

```
  -h, --help               help for get
  -m, --mhc-name string    Name of the MachineHealthCheck object
  -n, --namespace string   The namespace where the MachineHealthCheck object was created.
```

### Options inherited from parent commands

```
      --log-file string   Log file path
  -v, --verbose int32     Number for the log level verbosity(0-9)
```

### SEE ALSO

* [tanzu cluster machinehealthcheck node](tanzu_cluster_machinehealthcheck_node.md)	 - MachineHealthCheck operations for the nodes of a cluster

###### Auto generated by spf13/cobra on 14-Sep-2022