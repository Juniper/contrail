package compilationif

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/plugins/contrail/dependencies"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/services"
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
	err := service.handleEtcdMessages(ctx, oper, key, value)
	if err != nil {
		// TODO Add args to the message
		log.WithError(err).Error("Failed to handle etcd message")
	}
}

func (service *CompilationService) handleEtcdMessages(ctx context.Context, oper int32, key, value string) error {
	processor := &services.ServiceEventProcessor{Service: service}
	// TODO Extract a separate service that handles dependencies
	event, err := etcd.ParseEvent(oper, key, []byte(value))
	if err != nil {
		return errors.Wrap(err, "failed to parse ETCD event")
	}

	allEvents := service.resolveDependencies(event)

	for _, event := range allEvents {
		_, err := processor.Process(ctx, event)
		if err != nil {
			// TODO Add event and causer event to the error
			return errors.Wrap(err, "failed to process dependent event")
		}
	}

	return nil
}

func (service *CompilationService) resolveDependencies(event *services.Event) []*services.Event {
	dependencyProcessor := dependencies.NewDependencyProcessor(ObjsCache)
	resource := event.GetResource()
	dependencyProcessor.Evaluate(resource, resource.Kind(), "Self")
	allResources := dependencyProcessor.GetResources()

	result := make([]*services.Event, 0)
	allResources.Range(func(rawKind, rawResourceMap interface{}) bool {
		kind, _ := rawKind.(string)
		resourceMap, _ := rawResourceMap.(*sync.Map)

		resourceMap.Range(func(rawUUID, rawResource interface{}) bool {
			uuid, _ := rawUUID.(string)

			result = append(result, makeEmptyUpdateEvent(kind, uuid, rawResource))
			return true
		})
		return true
	})
	return result
}

func makeEmptyUpdateEvent(goType string, uuid string, resource interface{}) *services.Event {
	return services.NewEvent(&services.EventOption{
		Kind:      goType,
		UUID:      uuid,
		Operation: services.OperationUpdate,
		Data:      map[string]interface{}{},
	})
}
