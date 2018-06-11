package etcd

import (
	"encoding/json"

	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
)

// NotifierService is a service that performs writes to etcd.
type NotifierService struct {
	services.BaseService
	Path   string
	Client *Client
}

// NewNotifierService makes a etcdclient service.
func NewNotifierService(path string) (*NotifierService, error) {
	c, err := DialByConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("Cannot connect to ETCD server")
		return nil, err
	}
	service := &NotifierService{
		BaseService: services.BaseService{},
		Path:        path,
		Client:      NewClient(c),
	}
	return service, nil
}

// EtcdNotifierMarshal returns key/value string for given object
// TODO(Michal): use sink.Codec instead.
func (service *NotifierService) EtcdNotifierMarshal(schemaID string, objUUID string, obj interface{}) (string, []byte) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": schemaID,
		}).Debug("Create %s: Failed JSON Marshal", schemaID)
		return "", nil
	}

	objKey := "/" + service.Path + "/" + schemaID + "/" + objUUID
	return objKey, jsonBytes
}
