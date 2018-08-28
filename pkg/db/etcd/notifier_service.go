package etcd

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// NotifierService is a service that performs writes to etcd.
type NotifierService struct {
	services.BaseService
	Path   string
	Client *Client
	Codec  models.Codec
	log    *log.Entry
}

// NewNotifierService makes a etcdclient service.
func NewNotifierService(path string, codec models.Codec) (*NotifierService, error) {
	c, err := DialByConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to etcd server")
	}

	service := &NotifierService{
		Path:   path,
		Client: NewClient(c),
		Codec:  codec,
		log:    pkglog.NewLogger("etcd-notifier"),
	}
	return service, nil
}
