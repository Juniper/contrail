// Package sync contains functionality that supplies etcd with data from PostgreSQL database.
package sync

import (
	"context"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/etcd"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/sync/replication"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	asfetcd "github.com/Juniper/asf/pkg/etcd"
)

const (
	syncID = "sync-service"
)

// Service represents Sync service.
type Service struct {
	watcher *replication.PostgresWatcher

	cancel context.CancelFunc
	log    *logrus.Entry
}

// NewService creates Sync service with given configuration.
// Close needs to be explicitly called on service teardown.
func NewService() (*Service, error) {
	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}

	c := determineCodecType()
	if c == nil {
		return nil, errors.New(`unknown codec set as "sync.storage"`)
	}

	etcdNotifierService, err := etcd.NewNotifierService(viper.GetString(asfetcd.ETCDPathVK), c)
	if err != nil {
		return nil, err
	}

	watcher, err := NewEventProducer(
		syncID,
		&services.ServiceEventProcessor{Service: etcdNotifierService},
		etcdNotifierService.Client,
	)
	if err != nil {
		return nil, err
	}

	return &Service{
		watcher: watcher,
		log:     logutil.NewLogger(syncID),
	}, nil
}

func determineCodecType() models.Codec {
	switch viper.GetString("sync.storage") {
	case models.JSONCodec.Key():
		return models.JSONCodec
	case models.ProtoCodec.Key():
		return models.ProtoCodec
	default:
		return nil
	}
}

func NewEventProducer(
	id string, processor services.EventProcessor, txn services.InTransactionDoer,
) (*replication.PostgresWatcher, error) {
	setViperDefaults()

	dbService, err := db.NewServiceFromConfig()
	if err != nil {
		return nil, err
	}
	handler := &eventChangeHandler{
		processor: &services.EventListProcessor{
			EventProcessor:    NewFQNameCache(processor),
			InTransactionDoer: txn,
		},
		decoder: dbService,
	}

	if !viper.GetBool("sync.dump") {
		return replication.NewPostgresWatcher(handler, dbService, replication.Slot(id), replication.NoDump())
	}
	return replication.NewPostgresWatcher(handler, dbService, replication.Slot(id))
}

func setViperDefaults() {
	viper.SetDefault("log_level", "debug")
	viper.SetDefault(asfetcd.ETCDDialTimeoutVK, "60s")
	viper.SetDefault("database.retry_period", "1s")
	viper.SetDefault("database.connection_retries", 10)
	viper.SetDefault("database.replication_status_timeout", "10s")
}

// Run runs Sync service.
func (s *Service) Run() error {
	s.log.Info("Running Sync service")
	ctx := context.Background()
	ctx, s.cancel = context.WithCancel(ctx)
	return s.watcher.Start(ctx)
}

// DumpDone returns a channel that is closed when dump is done.
func (s *Service) DumpDone() <-chan struct{} {
	return s.watcher.DumpDone()
}

// Close closes Sync service.
func (s *Service) Close() {
	s.log.Info("Closing Sync service")
	s.cancel()
}

type eventListProcessor interface {
	ProcessList(context.Context, *services.EventList) (*services.EventList, error)
}

type eventDecoder interface {
	DecodeRowEvent(operation, resourceName string, pk []string, properties map[string]interface{}) (*services.Event, error)
}

type eventChangeHandler struct {
	processor eventListProcessor
	decoder   eventDecoder
}

func (e *eventChangeHandler) Handle(ctx context.Context, changes []replication.Change) error {
	list := services.EventList{}
	for _, c := range changes {
		ev, err := e.decoder.DecodeRowEvent(
			changeOperationToServices(c.Operation()),
			c.Kind(),
			c.PK(),
			c.Data(),
		)
		if err != nil {
			return err
		}
		list.Events = append(list.Events, ev)
	}
	_, err := e.processor.ProcessList(ctx, &list)
	return err
}

func changeOperationToServices(op replication.ChangeOperation) string {
	switch op {
	case replication.CreateOperation:
		return services.OperationCreate
	case replication.UpdateOperation:
		return services.OperationUpdate
	case replication.DeleteOperation:
		return services.OperationDelete
	default:
		return services.OperationCreate
	}
}
