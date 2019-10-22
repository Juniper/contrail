module github.com/Juniper/contrail

go 1.13

require (
	github.com/Juniper/asf v0.0.0-20200329072727-54d96252c742
	github.com/apparentlymart/go-cidr v1.0.1
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/flosch/pongo2 v0.0.0-20190707114632-bbf5a6c351f4
	github.com/gocql/gocql v0.0.0-20200324094621-6d895e38b0a5
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.3
	github.com/labstack/echo v3.3.10+incompatible
	github.com/mattn/goveralls v0.0.5
	github.com/mitchellh/gox v1.0.1
	github.com/pkg/errors v0.9.1
	github.com/pseudomuto/protoc-gen-doc v1.3.1
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.2
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.4.0
	golang.org/x/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0
	golang.org/x/tools v0.0.0-20200130002326-2f3ba24bd6e7
	google.golang.org/grpc v1.28.1
	gopkg.in/yaml.v2 v2.2.8
	sigs.k8s.io/yaml v1.2.0 // indirect
)

// workaround for https://github.com/gin-gonic/gin/issues/1673, etcd v3.2.2 dependency
replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go v0.0.0-20181204163529-d75b2dcb6bc8
