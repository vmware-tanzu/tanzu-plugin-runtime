// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package types provides all types, validator, helpers methods to create Runtime API commands
package types

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// FeatureCli allows a feature to be set at the CLI level (globally) rather than for a single plugin
	FeatureCli string = "cli"

	EditionStandard  = "tkg"
	EditionCommunity = "tce"
)

// EditionSelector allows selecting edition versions based on config file
type EditionSelector string

// VersionSelectorLevel allows selecting plugin versions based on semver properties
type VersionSelectorLevel string

// ServerType is the type of server.
// Note: Shall be deprecated in a future version. Superseded by ContextType.
type ServerType string

// ContextType is the type of the context (control plane).
type ContextType string

const (
	// ManagementClusterServerType is a management cluster server.
	// Note: Shall be deprecated in a future version. Superseded by CtxTypeK8s.
	ManagementClusterServerType ServerType = "managementcluster"

	// GlobalServerType is a global control plane server.
	// Note: Shall be deprecated in a future version. Superseded by CtxTypeTMC.
	GlobalServerType ServerType = "global"

	// CtxTypeK8s is a kubernetes cluster API server.
	CtxTypeK8s ContextType = "k8s"

	// CtxTypeTMC is a Tanzu Mission Control server.
	CtxTypeTMC ContextType = "tmc"
)

// Target is the namespace of the CLI to which plugin is applicable
type Target string

const (
	// TargetK8s is a kubernetes target of the CLI
	TargetK8s Target = "kubernetes"

	//nolint: deadcode, unused, varcheck
	targetK8s Target = "k8s"

	// TargetTMC is a Tanzu Mission Control target of the CLI
	TargetTMC Target = "mission-control"

	//nolint: deadcode, unused, varcheck
	targetTMC Target = "tmc"

	// TargetGlobal is used for plugins that are not associated with any target
	TargetGlobal Target = "global"

	// TargetUnknown specifies that the target is not currently known
	TargetUnknown Target = ""
)

// ManagementClusterServerOpts is the configuration for a management cluster kubeconfig.
// Note: Shall be deprecated in a future version. Superseded by ClusterServer.
type ManagementClusterServerOpts struct {
	// Endpoint for the login.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// Path to the kubeconfig.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The context to use (if required), defaults to current.
	Context string `json:"context,omitempty" yaml:"context,omitempty"`
}

// ClusterServerOpts contains the configuration for a kubernetes cluster (kubeconfig).
type ClusterServerOpts struct {
	// Endpoint for the login.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// Path to the kubeconfig.
	Path string `json:"path,omitempty" yaml:"path,omitempty,omitempty"`

	// The kubernetes context to use (if required), defaults to current.
	Context string `json:"context,omitempty" yaml:"context,omitempty"`

	// Denotes whether this server is a management cluster or not (workload cluster).
	IsManagementCluster bool `json:"isManagementCluster,omitempty" yaml:"isManagementCluster,omitempty"`
}

// GlobalServerOpts is the configuration for a global server.
type GlobalServerOpts struct {
	// Endpoint for the server.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// Auth for the global server.
	Auth GlobalServerAuthOpts `json:"auth,omitempty" yaml:"auth,omitempty"`
}

// GlobalServerAuthOpts is authentication for a global server.
type GlobalServerAuthOpts struct {
	// Issuer url for IDP, compliant with OIDC Metadata Discovery.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	// UserName is the authorized user the token is assigned to.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`

	// Permissions are roles assigned to the user.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// AccessToken is the current access token based on the context.
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// IDToken is the current id token based on the context scoped to the CLI.
	IDToken string `json:"IDToken,omitempty" yaml:"IDToken,omitempty"`

	// RefreshToken will be stored only in case of api-token login flow.
	RefreshToken string `json:"refresh_token,omitempty" yaml:"refresh_token,omitempty"`

	// Expiration times of the token.
	Expiration metav1.Time `json:"expiration,omitempty" yaml:"expiration,omitempty"`

	// Type of the token (user or client).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// FeatureMap is simply a hash table, but needs an explicit type to be an object in another hash map (cf ClientOptions.Features)
type FeatureMap map[string]string

// EnvMap is simply a hash table, but needs an explicit type to be an object in another hash map (cf ClientOptions.Env)
type EnvMap map[string]string

// GCPDiscoveryOpts provides a plugin discovery mechanism via a Google Cloud Storage
// bucket with a manifest.yaml file.
type GCPDiscoveryOpts struct {
	// Name is a name of the discovery
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Bucket is a Google Cloud Storage bucket.
	// E.g., tanzu-cli
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
	// BasePath is a URI path that is prefixed to the object name/path.
	// E.g., plugins/cluster
	ManifestPath string `json:"manifestPath" yaml:"manifestPath,omitempty"`
}

// OCIDiscoveryOpts provides a plugin discovery mechanism via a OCI Image Registry
type OCIDiscoveryOpts struct {
	// Name is a name of the discovery
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Image is an OCI compliant image. Which include DNS-compatible registry name,
	// a valid URI path(MAY contain zero or more ‘/’) and a valid tag.
	// E.g., harbor.my-domain.local/tanzu-cli/plugins-manifest:latest
	// Contains a directory containing YAML files, each of which contains single
	// CLIPlugin API resource.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`
}

// GenericRESTDiscoveryOpts provides a plugin discovery mechanism via any REST API
// endpoint. The fully qualified list URL is constructed as
// `https://{Endpoint}/{BasePath}` and the get plugin URL is constructed as .
// `https://{Endpoint}/{BasePath}/{Plugin}`.
type GenericRESTDiscoveryOpts struct {
	// Name is a name of the discovery
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Endpoint is the REST API server endpoint.
	// E.g., api.my-domain.local
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	// BasePath is the base URL path of the plugin discovery API.
	// E.g., /v1alpha1/cli/plugins
	BasePath string `json:"basePath,omitempty" yaml:"basePath,omitempty"`
}

// KubernetesDiscoveryOpts provides a plugin discovery mechanism via the Kubernetes API server.
type KubernetesDiscoveryOpts struct {
	// Name is a name of the discovery
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Path to the kubeconfig.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
	// The context to use (if required), defaults to current.
	Context string `json:"context,omitempty" yaml:"context,omitempty"`
	// Version of the CLIPlugins API to query.
	// E.g., v1alpha1
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

// LocalDiscoveryOpts is a artifact discovery endpoint utilizing a local host OS.
type LocalDiscoveryOpts struct {
	// Name is a name of the discovery
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Path is a local path pointing to directory
	// containing YAML files, each of which contains single
	// CLIPlugin API resource.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}

// PluginRepositoryOpts is a CLI plugin repository
type PluginRepositoryOpts struct {
	// GCPPluginRepository is a plugin repository that utilizes GCP cloud storage.
	GCPPluginRepository *GCPPluginRepositoryOpts `json:"gcpPluginRepository,omitempty" yaml:"gcpPluginRepository,omitempty"`
}

// GCPPluginRepositoryOpts is a plugin repository that utilizes GCP cloud storage.
type GCPPluginRepositoryOpts struct {
	// Name of the repository.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// BucketName is the name of the bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// RootPath within the bucket.
	RootPath string `json:"rootPath,omitempty" yaml:"rootPath,omitempty"`
}

// ServerOpts connection.
// Note: Shall be deprecated in a future version. Superseded by Context.
type ServerOpts struct {
	// Name of the server.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Type of the endpoint.
	Type ServerType `json:"type,omitempty" yaml:"type,omitempty"`

	// GlobalOpts if the server is global.
	GlobalOpts *GlobalServerOpts `json:"globalOpts,omitempty" yaml:"globalOpts,omitempty"`

	// ManagementClusterOpts if the server is a management cluster.
	ManagementClusterOpts *ManagementClusterServerOpts `json:"managementClusterOpts,omitempty" yaml:"managementClusterOpts,omitempty"`

	// DiscoverySources determines from where to discover plugins
	// associated with this server
	DiscoverySources []PluginDiscoveryOpts `json:"discoverySources,omitempty" yaml:"discoverySources,omitempty"`
}

// ContextOpts configuration for a control plane. This can one of the following,
// 1. Kubernetes Cluster
// 2. Tanzu Mission Control endpoint
// ContextOpts is the super set of parameters for Context, for all Runtime Versions. Based on the Runtime Version CtxOptions attributes may change(mandatory/optional/Not applicable).
// Command Helper functions (i.e. NewSetContextCommand) validate the supplied inputOptions as per RuntimeVersion and make sure all required attributes are set.
// If not Command Helper functions will throw error and Test fails.
type ContextOpts struct {
	// Name of the context.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Type of the context.
	Type ContextType `json:"type,omitempty" yaml:"type,omitempty"`

	// Target of the context.
	Target Target `json:"target,omitempty" yaml:"target,omitempty"`

	// GlobalOpts if the context is a global control plane (e.g., TMC).
	GlobalOpts *GlobalServerOpts `json:"globalOpts,omitempty" yaml:"globalOpts,omitempty"`

	// ClusterOpts if the context is a kubernetes cluster.
	ClusterOpts *ClusterServerOpts `json:"clusterOpts,omitempty" yaml:"clusterOpts,omitempty"`

	// DiscoverySources determines from where to discover plugins
	// associated with this context.
	DiscoverySources []PluginDiscoveryOpts `json:"discoverySources,omitempty" yaml:"discoverySources,omitempty"`
}

// ClientOptionsOpts are the client specific options.
type ClientOptionsOpts struct {
	// CLI options specific to the CLI.
	CLI      *CLIOptionsOpts       `json:"cli,omitempty" yaml:"cli,omitempty"`
	Features map[string]FeatureMap `json:"features,omitempty" yaml:"features,omitempty"`
	Env      map[string]string     `json:"env,omitempty" yaml:"env,omitempty"`
}

// CLIOptionsOpts are options for the CLI.
type CLIOptionsOpts struct {
	// Repositories are the plugin repositories.
	Repositories []PluginRepositoryOpts `json:"repositories,omitempty" yaml:"repositories,omitempty"`
	// DiscoverySources determines from where to discover stand-alone plugins
	DiscoverySources []PluginDiscoveryOpts `json:"discoverySources,omitempty" yaml:"discoverySources,omitempty"`
	// UnstableVersionSelector determined which version tags are allowed
	UnstableVersionSelector VersionSelectorLevel `json:"unstableVersionSelector,omitempty" yaml:"unstableVersionSelector,omitempty"`
	// Edition
	Edition EditionSelector `json:"edition,omitempty" yaml:"edition,omitempty"`
	// BOMRepo is the root repository URL used to resolve the compatibiilty file
	// and bill of materials. An example URL is projects.registry.vmware.com/tkg.
	BOMRepo string `json:"bomRepo,omitempty" yaml:"bomRepo,omitempty"`
	// CompatibilityFilePath is the path, from the BOM repo, to download and access the compatibility file.
	// the compatibility file is used for resolving the bill of materials for creating clusters.
	CompatibilityFilePath string `json:"compatibilityFilePath,omitempty" yaml:"compatibilityFilePath,omitempty"`
}

// PluginDiscoveryOpts contains a specific distribution mechanism. Only one of the
// configs must be set.
type PluginDiscoveryOpts struct {
	// GCPStorage is set if the plugins are to be discovered via Google Cloud Storage.
	GCP *GCPDiscoveryOpts `json:"gcp,omitempty" yaml:"gcp,omitempty"`
	// OCIDiscovery is set if the plugins are to be discovered via an OCI Image Registry.
	OCI *OCIDiscoveryOpts `json:"oci,omitempty" yaml:"oci,omitempty"`
	// GenericRESTDiscovery is set if the plugins are to be discovered via a REST API endpoint.
	REST *GenericRESTDiscoveryOpts `json:"rest,omitempty" yaml:"rest,omitempty"`
	// KubernetesDiscovery is set if the plugins are to be discovered via the Kubernetes API server.
	Kubernetes *KubernetesDiscoveryOpts `json:"k8s,omitempty" yaml:"k8s,omitempty"`
	// LocalDiscovery is set if the plugins are to be discovered via Local Manifest fast.
	Local *LocalDiscoveryOpts `json:"local,omitempty" yaml:"local,omitempty"`
	// ContextType the discovery source is associated with (applicable only for stand-alone plugins).
	ContextType ContextType `json:"contextType,omitempty" yaml:"contextType,omitempty"`
}

// ClientConfigOpts is the Schema for the configs API
type ClientConfigOpts struct {
	metav1.TypeMeta   `json:",omitempty" yaml:",omitempty"`
	metav1.ObjectMeta `json:",omitempty" yaml:",omitempty"`

	// KnownServers available.
	// Note: Shall be deprecated in a future version. Superseded by KnownContexts.
	KnownServers []*ServerOpts `json:"servers,omitempty" yaml:"servers,omitempty"`

	// CurrentServer in use.
	// Note: Shall be deprecated in a future version. Superseded by CurrentContext.
	CurrentServer string `json:"current,omitempty" yaml:"current,omitempty"`

	// KnownContexts available.
	KnownContexts []*ContextOpts `json:"contexts,omitempty" yaml:"contexts,omitempty"`

	// CurrentContext for every type.
	CurrentContext map[string]string `json:"currentContext,omitempty" yaml:"currentContext,omitempty"`

	// CurrentContext for every type.
	CurrentContextWithContextType map[ContextType]string `json:"currentContextContextType,omitempty" yaml:"currentContextContextType,omitempty"`

	// CurrentContext for every type.
	CurrentContextWithTarget map[Target]string `json:"currentContextTarget,omitempty" yaml:"currentContextTarget,omitempty"`

	// ClientOptions are client specific options.
	ClientOptions *ClientOptionsOpts `json:"clientOptions,omitempty" yaml:"clientOptions,omitempty"`
}
