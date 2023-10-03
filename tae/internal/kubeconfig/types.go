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
	// Preferences holds general information to be used for cli interactions
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
	Name string `json:"name" yaml:"name"`
	// Note:!! Changed it to empty interface as runtime API is going to use the AuthInfo as is
	AuthInfo interface{} `json:"user" yaml:"user"`
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
