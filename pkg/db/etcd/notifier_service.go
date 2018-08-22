package etcd

import (
	"encoding/json"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
)

// NotifierService is a service that performs writes to etcd.
type NotifierService struct {
	services.BaseService
	Path   string
	Client *Client
	log    *log.Entry
}

// NewNotifierService makes a etcdclient service.
func NewNotifierService(path string) (*NotifierService, error) {
	c, err := DialByConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to etcd server")
	}

	service := &NotifierService{
		BaseService: services.BaseService{},
		Path:        path,
		Client:      NewClient(c),
		log:         pkglog.NewLogger("etcd-notifier"),
	}
	return service, nil
}

// EtcdNotifierMarshal returns key/value string for given object
// TODO(Michal): use sink.Codec instead.
func (ns *NotifierService) EtcdNotifierMarshal(schemaID string, objUUID string, obj interface{}) (string, []byte) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": schemaID,
		}).Debug("Create %s: Failed JSON Marshal", schemaID)
		return "", nil
	}

	objKey := "/" + ns.Path + "/" + schemaID + "/" + objUUID
	return objKey, jsonBytes
}
