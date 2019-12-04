module github.com/Juniper/asf

go 1.13

require (
	github.com/DavidCai1993/etcd-lock v0.0.0-20171006032119-32f65b8d019a
	github.com/ExpansiveWorlds/instrumentedsql v0.0.0-20171218214018-45abb4b1947d
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/apparentlymart/go-cidr v1.0.1
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/codegangsta/cli v1.20.0 // indirect
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/elazarl/go-bindata-assetfs v1.0.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20190707114632-bbf5a6c351f4
	github.com/go-openapi/spec v0.19.3
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/gocql/gocql v0.0.0-20191018090344-07ace3bab0f8
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.3.1
	github.com/google/uuid v1.1.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.0+incompatible
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/kyleconroy/pgoutput v0.1.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/lib/pq v1.2.0
	github.com/mattn/go-shellwords v1.0.6
	github.com/mattn/go-sqlite3 v1.11.0 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v0.0.0-20191009025716-f1972eb1d1f5 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/stretchr/testify v1.4.0
	github.com/valyala/fasttemplate v1.1.0 // indirect
	github.com/volatiletech/inflect v0.0.0-20170731032912-e7201282ae8d // indirect
	github.com/volatiletech/sqlboiler v3.5.0+incompatible
	github.com/yudai/gotty v2.0.0-alpha.3+incompatible
	github.com/yudai/hcl v0.0.0-20151013225006-5fa2393b3552 // indirect
	go.opencensus.io v0.22.1 // indirect
	golang.org/x/net v0.0.0-20191014212845-da9a3fd4c582
	google.golang.org/api v0.10.0 // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/yaml.v2 v2.2.4
	sigs.k8s.io/yaml v1.1.0 // indirect
)

// github.com/codegangsta/cli was moved, but github.com/yudai/gotty uses obsolete import path
replace github.com/codegangsta/cli v1.20.0 => github.com/urfave/cli v1.20.0
