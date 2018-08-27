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
	inTransactionDoer services.InTransactionDoer
	client            *client.HTTP
	log               *log.Entry
}

// Config is NotifierService config.
type Config struct {
	Endpoint          string
	InTransactionDoer services.InTransactionDoer
}

// NewNotifierService makes a NotifierService.
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
		inTransactionDoer: c.InTransactionDoer,
		client:            client,
		log:               pkglog.NewLogger("vnc-api-notifier"),
	}
}
