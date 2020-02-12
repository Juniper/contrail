package etcd

import (
	"path"

	"github.com/spf13/viper"
)

const (
	ETCDEndpointsVK          = "etcd.endpoints"
	ETCDDialTimeoutVK        = "etcd.dial_timeout"
	ETCDGRPCInsecureVK       = "etcd.grpc_insecure"
	ETCDPasswordVK           = "etcd.password"
	ETCDPathVK               = "etcd.path"
	ETCDTLSEnabledVK         = "etcd.tls.enabled"
	ETCDTLSCertificatePathVK = "etcd.tls.certificate_path"
	ETCDTLSKeyPathVK         = "etcd.tls.key_path"
	ETCDTLSTrustedCAPathVK   = "etcd.tls.trusted_ca_path"
	ETCDUsernameVK           = "etcd.username"
)

// ResourceKey constructs key for given resource type and pk.
// TODO(dfurman): pass ETCDPathVK value instead of reading it from the global configuration.
func ResourceKey(resourceType, pk string) string {
	return path.Join("/", viper.GetString(ETCDPathVK), resourceType, pk)
}
