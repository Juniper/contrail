package contrailutil

import (
	"context"
	"fmt"
	"time"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/cassandra"
	etcdclient "github.com/Juniper/contrail/pkg/db/etcd"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	yamlType          = "yaml"
	cassandraType     = "cassandra"
	cassandraDumpType = "cassandra_dump"
	rdbmsType         = "rdbms"
	etcdType          = "etcd"
)

func init() {
	ContrailUtil.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&inType, "intype", "", "",
		`input type: "cassandra", "cassandra_dump", "yaml" and "rdbms" are supported`)
	convertCmd.Flags().StringVarP(&inFile, "in", "i", "", "Input file or Cassandra host")
	convertCmd.Flags().StringVarP(&outType, "outtype", "", "",
		`output type: "rdbms", "yaml" and "etcd" are supported`)
	convertCmd.Flags().StringVarP(&outFile, "out", "o", "", "Output file")
	convertCmd.Flags().IntVarP(&cassandraPort, "cassandra_port", "p", 9042, "Cassandra port")
	convertCmd.Flags().IntVarP(&cassandraTimeout, "cassandra_timeout", "t", 3600, "Cassandra timeout in seconds")
}

var inType, inFile string
var outType, outFile string
var cassandraPort, cassandraTimeout int

func readYAML() (*services.EventList, error) {
	var events services.EventList
	err := common.LoadFile(inFile, &events)
	return &events, err
}

func writeYAML(events *services.EventList) error {
	return common.SaveFile(outFile, events)
}

func readRDBMS() (*services.EventList, error) {
	dbService, err := db.NewServiceFromConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect DB")
	}
	ctx := context.Background()
	return services.Dump(ctx, dbService)
}

func writeRDBMS(events *services.EventList) error {
	dbService, err := db.NewServiceFromConfig()
	if err != nil {
		return errors.Wrap(err, "failed to connect DB")
	}
	ctx := context.Background()
	err = dbService.DoInTransaction(ctx,
		func(ctx context.Context) error {
			_, err = events.Process(ctx, dbService)
			return err
		})
	return err
}

func writeEtcd(events *services.EventList) error {
	etcdNotifierPath := viper.GetString("etcd.path")
	etcdNotifierService, err := etcdclient.NewNotifierService(etcdNotifierPath)
	if err != nil {
		return errors.Wrap(err, "failed to connect etcd")
	}
	etcdNotifierService.SetNext(&services.BaseService{})
	ctx := context.Background()
	_, err = events.Process(ctx, etcdNotifierService)
	return err
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert data format",
	Long:  `This command converts data formats from one to another`,
	Run: func(cmd *cobra.Command, args []string) {
		var events *services.EventList
		var err error
		switch inType {
		case cassandraType:
			events, err = cassandra.DumpCassandra(
				cassandra.Config{
					Host: inFile,
					Port: cassandraPort,
					Timeout: time.Duration(cassandraTimeout)*time.Second,
				})
		case cassandraDumpType:
			events, err = cassandra.ReadCassandraDump(inFile)
		case yamlType:
			events, err = readYAML()
		case rdbmsType:
			events, err = readRDBMS()
		default:
			err = fmt.Errorf("Unsupported input type %s", inType)
		}
		if err != nil {
			log.Fatal(err)
		}
		err = events.Sort()
		if err != nil {
			log.Fatal(err)
		}

		switch outType {
		case rdbmsType:
			err = writeRDBMS(events)
		case yamlType:
			err = writeYAML(events)
		case etcdType:
			err = writeEtcd(events)
		default:
			err = fmt.Errorf("Unsupported input type %s", inType)
			log.Fatal("Unsupported output type")
		}
		if err != nil {
			log.Fatal(err)
		}
	},
}
