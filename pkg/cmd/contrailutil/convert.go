package contrailutil

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/convert"
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

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert data format",
	Long:  `This command converts data formats from one to another`,
	Run: func(cmd *cobra.Command, args []string) {
		err := convert.Convert(&convert.Config{
			InType:           inType,
			InFile:           inFile,
			OutType:          outType,
			OutFile:          outFile,
			CassandraPort:    cassandraPort,
			CassandraTimeout: cassandraTimeout,
			EtcdNotifierPath: viper.GetString("etcd.path"),
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}
