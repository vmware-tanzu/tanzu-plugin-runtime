// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package kubeconfig

// Note-IMPORTANT!! The below types are copied from the client-go library and modified a bit(removed fields like extensions not necessary for TAE kubeconfig)
// to provide kubeconfig access functionality without copying the k8s.io/client-go library which would bring in the other kubernetes dependencies
// and would impose dependency constraints on the plugin developers using the tanzu-plugin-runtime library

// Config holds the information needed to establish connection to remote kubernetes clusters as a given user
//
// (Note: !!modified Clusters,AuthInfos,Contexts from map to array to fit the Yaml marshaling and removed the extensions field)
type Config struct {
	Kind       string `json:"kind,omitempty" yaml:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	// Preferences holds general information to be use for cli interactions
	Preferences Preferences `json:"preferences" yaml:"preferences"`
	// Clusters is a map of referencable names to cluster configs
	Clusters []*Cluster `json:"clusters" yaml:"clusters"`
	// AuthInfos is a map of referencable names to user configs
	AuthInfos []*AuthInfo `json:"users" yaml:"users"`
	// Contexts is a map of referencable names to context configs
	Contexts []*Context `json:"contexts" yaml:"contexts"`
	// CurrentContext is the name of the context that you would like to use by default
	CurrentContext string `json:"current-context" yaml:"current-context"`
}

type Preferences struct {
	// +optional
	Colors bool `json:"colors,omitempty" yaml:"colors,omitempty"`
}

// Cluster contains information about how to communicate with a kubernetes cluster
type Cluster struct {
	Name    string `json:"name" yaml:"name"`
	Cluster struct {
		// Server is the address of the kubernetes cluster (https://hostname:port).
		Server string `json:"server" yaml:"server"`
		// TLSServerName is used to check server certificate. If TLSServerName is empty, the hostname used to contact the server is used.
		// +optional
		TLSServerName string `json:"tls-server-name,omitempty" yaml:"tls-server-name,omitempty"`
		// InsecureSkipTLSVerify skips the validity check for the server's certificate. This will make your HTTPS connections insecure.
		// +optional
		InsecureSkipTLSVerify bool `json:"insecure-skip-tls-verify,omitempty" yaml:"insecure-skip-tls-verify,omitempty"`
		// CertificateAuthority is the path to a cert file for the certificate authority.
		// +optional
		CertificateAuthority string `json:"certificate-authority,omitempty" yaml:"certificate-authority,omitempty"`
		// CertificateAuthorityData contains PEM-encoded certificate authority certificates. Overrides CertificateAuthority
		// +optional
		CertificateAuthorityData string `json:"certificate-authority-data,omitempty" yaml:"certificate-authority-data,omitempty"`
		// ProxyURL is the URL to the proxy to be used for all requests made by this
		// client. URLs with "http", "https", and "socks5" schemes are supported.  If
		// this configuration is not provided or the empty string, the client
		// attempts to construct a proxy configuration from http_proxy and
		// https_proxy environment variables. If these environment variables are not
		// set, the client does not attempt to proxy requests.
		//
		// socks5 proxying does not currently support spdy streaming endpoints (exec,
		// attach, port forward).
		// +optional
		ProxyURL string `json:"proxy-url,omitempty" yaml:"proxy-url,omitempty"`
		// DisableCompression allows client to opt-out of response compression for all requests to the server. This is useful
		// to speed up requests (specifically lists) when client-server network bandwidth is ample, by saving time on
		// compression (server-side) and decompression (client-side): https://github.com/kubernetes/kubernetes/issues/112296.
		// +optional
		DisableCompression bool `json:"disable-compression,omitempty" yaml:"disable-compression,omitempty"`
	} `json:"cluster" yaml:"cluster"`
}

// AuthInfo contains information that describes identity information.  This is use to tell the kubernetes cluster who you are.
type AuthInfo struct {
	Name     string `json:"name" yaml:"name"`
	AuthInfo struct {
		// ClientCertificate is the path to a client cert file for TLS.
		// +optional
		ClientCertificate string `json:"client-certificate,omitempty" yaml:"client-certificate,omitempty"`
		// ClientCertificateData contains PEM-encoded data from a client cert file for TLS. Overrides ClientCertificate
		// +optional
		ClientCertificateData string `json:"client-certificate-data,omitempty" yaml:"client-certificate-data,omitempty"`
		// ClientKey is the path to a client key file for TLS.
		// +optional
		ClientKey string `json:"client-key,omitempty" yaml:"client-key,omitempty"`
		// ClientKeyData contains PEM-encoded data from a client key file for TLS. Overrides ClientKey
		// +optional
		ClientKeyData string `json:"client-key-data,omitempty" yaml:"client-key-data,omitempty" datapolicy:"security-key"`
		// Token is the bearer token for authentication to the kubernetes cluster.
		// +optional
		Token string `json:"token,omitempty" yaml:"token,omitempty" datapolicy:"token"`
		// TokenFile is a pointer to a file that contains a bearer token (as described above).  If both Token and TokenFile are present, Token takes precedence.
		// +optional
		TokenFile string `json:"tokenFile,omitempty" yaml:"tokenFile,omitempty"`
		// Impersonate is the username to act-as.
		// +optional
		Impersonate string `json:"act-as,omitempty" yaml:"act-as,omitempty"`
		// ImpersonateUID is the uid to impersonate.
		// +optional
		ImpersonateUID string `json:"act-as-uid,omitempty" yaml:"act-as-uid,omitempty"`
		// ImpersonateGroups is the groups to impersonate.
		// +optional
		ImpersonateGroups []string `json:"act-as-groups,omitempty" yaml:"act-as-groups,omitempty"`
		// ImpersonateUserExtra contains additional information for impersonated user.
		// +optional
		ImpersonateUserExtra map[string][]string `json:"act-as-user-extra,omitempty" yaml:"act-as-user-extra,omitempty"`
		// Username is the username for basic authentication to the kubernetes cluster.
		// +optional
		Username string `json:"username,omitempty" yaml:"username,omitempty"`
		// Password is the password for basic authentication to the kubernetes cluster.
		// +optional
		Password string `json:"password,omitempty" yaml:"password,omitempty" datapolicy:"password"`
		// AuthProvider specifies a custom authentication plugin for the kubernetes cluster.
		// +optional
		AuthProvider *AuthProviderConfig `json:"auth-provider,omitempty" yaml:"auth-provider,omitempty"`
		// Exec specifies a custom exec-based authentication plugin for the kubernetes cluster.
		// +optional
		Exec *ExecConfig `json:"exec,omitempty" yaml:"exec,omitempty"`
	} `json:"user" yaml:"user"`
}

// Context is a tuple of references to a cluster (how do I communicate with a kubernetes cluster), a user (how do I identify myself), and a namespace (what subset of resources do I want to work with)
type Context struct {
	Name    string `json:"name" yaml:"name"`
	Context struct {
		// Cluster is the name of the cluster for this context
		Cluster string `json:"cluster" yaml:"cluster"`
		// AuthInfo is the name of the authInfo for this context
		AuthInfo string `json:"user" yaml:"user"`
		// Namespace is the default namespace to use on unspecified requests
		// +optional
		Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	} `json:"context" yaml:"context"`
}

// AuthProviderConfig holds the configuration for a specified auth provider.
type AuthProviderConfig struct {
	Name string `json:"name" yaml:"name"`
	// +optional
	Config map[string]string `json:"config,omitempty" yaml:"config,omitempty"`
}

// ExecConfig specifies a command to provide client credentials. The command is exec'd
// and outputs structured stdout holding credentials.
//
// See the client.authentication.k8s.io API group for specifications of the exact input
// and output format
type ExecConfig struct {
	// Command to execute.
	Command string `json:"command" yaml:"command"`
	// Arguments to pass to the command when executing it.
	// +optional
	Args []string `json:"args" yaml:"args"`
	// Env defines additional environment variables to expose to the process. These
	// are unioned with the host's environment, as well as variables client-go uses
	// to pass argument to the plugin.
	// +optional
	Env []ExecEnvVar `json:"env" yaml:"env"`

	// Preferred input version of the ExecInfo. The returned ExecCredentials MUST use
	// the same encoding version as the input.
	APIVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`

	// This text is shown to the user when the executable doesn't seem to be
	// present. For example, `brew install foo-cli` might be a good InstallHint for
	// foo-cli on Mac OS systems.
	InstallHint string `json:"installHint,omitempty" yaml:"installHint,omitempty"`

	// ProvideClusterInfo determines whether or not to provide cluster information,
	// which could potentially contain very large CA data, to this exec plugin as a
	// part of the KUBERNETES_EXEC_INFO environment variable. By default, it is set
	// to false. Package k8s.io/client-go/tools/auth/exec provides helper methods for
	// reading this environment variable.
	ProvideClusterInfo bool `json:"provideClusterInfo" yaml:"provideClusterInfo"`
}

// ExecEnvVar is used for setting environment variables when executing an exec-based
// credential plugin.
type ExecEnvVar struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}
