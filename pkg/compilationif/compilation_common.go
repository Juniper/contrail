package compilationif

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

// NewCompilationService makes a compilation service.
func NewCompilationService() *CompilationService {
	service := &CompilationService{
		BaseService: services.BaseService{},
	}
	service.Init()
	return service
}

// CompilationService
type CompilationService struct {
	services.BaseService
}

// HandleEtcdMessages
func (service *CompilationService) HandleEtcdMessages(ctx context.Context, oper int32, key, value string) {
	log.Printf("HandleEtcdMessages: Got Msg: Oper-%d, Key-%s, Value-%s", oper, key, value)
	subkeys := strings.Split(key, "/")
	if len(subkeys) != 4 {
		log.Errorf("HandleEtcdMessages: Malformed Key: %s", key)
		return
	}

	resourceType := subkeys[2]
	uuid := subkeys[3]

	err := service.HandleResource(ctx, oper, resourceType, uuid, value)
	if err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": resourceType,
		}).Debug("HandleEtcdMessages: Error Handling Etcd Message")
	}
}

func (service *CompilationService) HandleResource(ctx context.Context, oper int32, resourceType, uuid, value string) error {
	handle, ok := functions[resourceType]
	if !ok {
		return errors.Errorf("Unknown resource type: %s", resourceType)
	}
	return handle(service, ctx, oper, uuid, value)
}
