package vncapi

import (
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
)

// NotifierService is a service that proxies events to VNC API.
type NotifierService struct {
	services.BaseService
	client *client.HTTP
	log    *log.Entry
}

// Config is NotifierService config.
type Config struct {
	Endpoint string
}

// NewNotifierService makes a NotifierService service.
func NewNotifierService(c *Config) *NotifierService {
	client := client.NewHTTP(
		c.Endpoint,
		"",
		"",
		"",
		"",
		true,
		nil,
	)
	client.Init()

	return &NotifierService{
		client: client,
		log:    pkglog.NewLogger("vnc-api-notifier"),
	}
}
