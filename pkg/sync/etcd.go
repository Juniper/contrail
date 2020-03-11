package sync

import (
	"github.com/Juniper/asf/pkg/sync"
	"github.com/Juniper/contrail/pkg/etcd"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	asfetcd "github.com/Juniper/asf/pkg/etcd"
)

const syncCodecPath = "sync.storage"

// NewEtcdFeeder returns a new PostgresWatcher that is configured to feed etcd with object updates.
func NewEtcdFeeder(id string) (*sync.PostgresWatcher, error) {
	c := determineCodecType()
	if c == nil {
		return nil, errors.Errorf(`unknown codec set as "%s"`, syncCodecPath)
	}

	etcdNotifierService, err := etcd.NewNotifierService(viper.GetString(asfetcd.ETCDPathVK), c)
	if err != nil {
		return nil, err
	}

	return NewEventProducer(
		id,
		&services.ServiceEventProcessor{Service: etcdNotifierService},
		etcdNotifierService.Client,
	)
}

func determineCodecType() models.Codec {
	switch viper.GetString(syncCodecPath) {
	case models.JSONCodec.Key():
		return models.JSONCodec
	case models.ProtoCodec.Key():
		return models.ProtoCodec
	default:
		return nil
	}
}
