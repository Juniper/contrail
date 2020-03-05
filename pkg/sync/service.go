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

type watchCloser interface {
	Watch(context.Context) error
	DumpDone() <-chan struct{}
	Close()
}

// Service represents Sync service.
type Service struct {
	watcher watchCloser
	log     *logrus.Entry
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

	watcher, err := createWatcher(syncID, &services.EventListProcessor{
		EventProcessor:    NewFQNameCache(&services.ServiceEventProcessor{Service: etcdNotifierService}),
		InTransactionDoer: etcdNotifierService.Client,
	})
	if err != nil {
		return nil, err
	}

	return &Service{
		watcher: watcher,
		log:     logutil.NewLogger(syncID),
	}, nil
}

func setViperDefaults() {
	viper.SetDefault("log_level", "debug")
	viper.SetDefault(asfetcd.ETCDDialTimeoutVK, "60s")
	viper.SetDefault("database.retry_period", "1s")
	viper.SetDefault("database.connection_retries", 10)
	viper.SetDefault("database.replication_status_timeout", "10s")
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

func createWatcher(id string, processor eventProcessor) (watchCloser, error) {
	setViperDefaults()

	conn, err := replication.NewPostgresConnection()
	if err != nil {
		return nil, err
	}

	dbService, err := db.NewServiceFromConfig()
	if err != nil {
		return nil, err
	}
	handler := &EventChangeHandler{processor: processor, decoder: dbService}

	conf := replication.PostgresSubscriptionConfig{
		Slot:          replication.SlotName(id),
		Publication:   replication.PostgreSQLPublicationName,
		StatusTimeout: viper.GetDuration("database.replication_status_timeout"),
	}

	return replication.NewPostgresWatcher(conf, conn, handler, viper.GetBool("sync.dump")), nil
}

// Run runs Sync service.
func (s *Service) Run() error {
	s.log.Info("Running Sync service")
	return s.watcher.Watch(context.Background())
}

// DumpDone returns a channel that is closed when dump is done.
func (s *Service) DumpDone() <-chan struct{} {
	return s.watcher.DumpDone()
}

// Close closes Sync service.
func (s *Service) Close() {
	s.log.Info("Closing Sync service")
	s.watcher.Close()
}
