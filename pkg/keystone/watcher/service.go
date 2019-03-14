package watcher

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/logutil"
)

type Service struct {
	authURL   string
	apiServer *client.HTTP
	keystone  *client.Keystone
}

// NewKeystoneWatcherByConfig creates a service watcher that listen kieystone for project changes
func NewKeystoneWatcherByConfig() (*Service, error) {
	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}
	authURL := viper.GetString("keystone.authurl")
	if authURL == "" {
		return nil, errors.New("missing config option keystone.authurl needed by keystone watcher")
	}

	endpoint := viper.GetString("client.endpoint")
	insecure := viper.GetBool("insecure")
	api := &client.HTTP{
		Endpoint: endpoint,
		InSecure: insecure,
	}
	api.ID = viper.GetString("client.id")
	api.Password = viper.GetString("client.password")
	// XXX WIP !!!!!!!!!
	api.Scope = keystone.NewScope(c.DomainID, c.DomainName,
		c.ProjectID, c.ProjectName)
	return &Service{apiServer: api}, nil
}

// Watch starts listenting for project cheges in keystone
func (sv *Service) Watch() {
	for {
	}
}
