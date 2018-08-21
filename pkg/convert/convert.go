package convert

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/cassandra"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/services"
)

// Data source and destination types.
const (
	YAMLType          = "yaml"
	CassandraType     = "cassandra"
	CassandraDumpType = "cassandra_dump"
	RDBMSType         = "rdbms"
	EtcdType          = "etcd"
)

// Config for Convert command.
type Config struct {
	InType                  string
	InFile                  string
	OutType                 string
	OutFile                 string
	CassandraPort           int
	CassandraTimeout        int
	CassandraConnectTimeout int
	EtcdNotifierPath        string
}

// Convert converts data from one format to another.
func Convert(c *Config) error {
	events, err := readData(c)
	if err != nil {
		return errors.Wrapf(err, "reading events from %v failed", c.InType)
	}

	err = events.Sort()
	if err != nil {
		return errors.Wrap(err, "sorting events failed")
	}

	err = writeData(events, c)
	if err != nil {
		return errors.Wrapf(err, "writing events to %v failed", c.OutType)
	}
	return nil
}

func readData(c *Config) (*services.EventList, error) {
	switch c.InType {
	case CassandraType:
		cassCfg := cassandra.Config{
			Host:           c.InFile,
			Port:           c.CassandraPort,
			Timeout:        time.Duration(c.CassandraTimeout) * time.Second,
			ConnectTimeout: time.Duration(c.CassandraConnectTimeout) * time.Millisecond,
		}
		return cassandra.DumpCassandra(cassCfg)
	case CassandraDumpType:
		return cassandra.ReadCassandraDump(c.InFile)
	case YAMLType:
		return readYAML(c.InFile)
	case RDBMSType:
		return readRDBMS()
	default:
		return nil, errors.Errorf("unsupported input type %v", c.InType)
	}
}

func writeData(events *services.EventList, c *Config) error {
	switch c.OutType {
	case RDBMSType:
		return writeRDBMS(events)
	case YAMLType:
		return writeYAML(events, c.OutFile)
	case EtcdType:
		return writeEtcd(events, c.EtcdNotifierPath)
	default:
		return errors.Errorf("unsupported output type %v", c.OutType)
	}
}

func readYAML(inFile string) (*services.EventList, error) {
	var events services.EventList
	err := common.LoadFile(inFile, &events)
	return &events, err
}

func writeYAML(events *services.EventList, outFile string) error {
	return common.SaveFile(outFile, events)
}

func readRDBMS() (*services.EventList, error) {
	dbService, err := db.NewServiceFromConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect DB")
	}

	return services.Dump(context.Background(), dbService)
}

func writeRDBMS(events *services.EventList) error {
	dbService, err := db.NewServiceFromConfig()
	if err != nil {
		return errors.Wrap(err, "failed to connect to DB")
	}

	return dbService.DoInTransaction(
		context.Background(),
		func(ctx context.Context) error {
			_, err = events.Process(ctx, dbService)
			return err
		},
	)
}

func writeEtcd(events *services.EventList, etcdNotifierPath string) error {
	etcdNotifierService, err := etcd.NewNotifierService(etcdNotifierPath)
	if err != nil {
		return errors.Wrap(err, "failed to connect to etcd")
	}

	etcdNotifierService.SetNext(&services.BaseService{})

	_, err = events.Process(context.Background(), etcdNotifierService)
	return errors.Wrap(err, "processing events on etcdNotifierService failed")
}
