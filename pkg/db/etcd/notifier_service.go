package etcd

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// NotifierService is a service that performs writes to etcd.
type NotifierService struct {
	services.BaseService
	Path   string
	Client *Client
	Codec  models.Codec
	log    *logrus.Entry
}

// NewNotifierService creates a etcd Notifier Service.
func NewNotifierService(path string, codec models.Codec) (*NotifierService, error) {
	ec, err := NewClientByViper("etcd-notifier")
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to etcd server")
	}

	service := &NotifierService{
		Path:   path,
		Client: ec,
		Codec:  codec,
		log:    logutil.NewLogger("etcd-notifier"),
	}
	return service, nil
}
