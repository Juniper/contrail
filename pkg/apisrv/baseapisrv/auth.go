package baseapisrv

import (
	"github.com/Juniper/asf/pkg/keystone"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func (s *Server) authPlugins(authSkipPaths []string) (plugins []APIPlugin, err error) {
	if viper.GetString("keystone.authurl") != "" {
		plugins = append(plugins, NewAuthPluginByViper(authSkipPaths))
	} else if viper.GetBool("no_auth") {
		plugins = append(plugins, noAuthPlugin{})
	}
	return plugins, nil
}

// AuthPlugin authenticates requests to the server with Keystone.
type AuthPlugin struct {
	m *keystone.AuthMiddleware
}

// NewAuthPluginByViper creates an AuthPlugin based on global Viper configuration.
func NewAuthPluginByViper(skipPaths []string) *AuthPlugin {
	authURL := viper.GetString("keystone.authurl")
	insecure := viper.GetBool("keystone.insecure")

	return &AuthPlugin{
		m: keystone.NewAuthMiddleware(authURL, insecure, skipPaths),
	}
}

// RegisterHTTPAPI registers authentication middleware for most endpoints in the server.
func (p *AuthPlugin) RegisterHTTPAPI(r HTTPRouter) {
	r.Use(p.m.HTTPMiddleware)
}

// RegisterGRPCAPI registers an authentication interceptor for GRPC.
func (p *AuthPlugin) RegisterGRPCAPI(r GRPCRouter) {
	r.AddServerOptions(grpc.UnaryInterceptor(p.m.GRPCInterceptor))
}
