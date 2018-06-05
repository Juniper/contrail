package etcd

import (
	"encoding/json"

	"github.com/Juniper/contrail/pkg/serviceif"
	log "github.com/sirupsen/logrus"
)

type NotifierService struct {
	serviceif.BaseService
	Path   string
	Etcdcl *Client
}

// NewNotifierService makes a etcdclient service.
func NewNotifierService(path string) (*NotifierService, error) {
	etcdcl, err := Dial()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("Cannot connect to ETCD server")
		return nil, err
	}
	service := &NotifierService{
		BaseService: serviceif.BaseService{},
		Path:        path,
		Etcdcl:      etcdcl,
	}
	return service, nil
}

// EtcdNotifierMarshal returns key/value string for given object
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
