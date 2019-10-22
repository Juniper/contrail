module github.com/Juniper/contrail

go 1.12

require (
	github.com/DavidCai1993/etcd-lock v0.0.0-20171006032119-32f65b8d019a // indirect
	github.com/Juniper/asf v0.0.0-20200329072727-54d96252c742
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/apparentlymart/go-cidr v1.0.1
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/aws/aws-sdk-go v1.16.6 // indirect
	github.com/bitly/go-hostpool v0.0.0-20171023180738-a3a6125de932 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/coreos/etcd v3.3.2+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20190707114632-bbf5a6c351f4
	github.com/gocql/gocql v0.0.0-20180530083731-3c37daec2f4d
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.3.1
	github.com/google/uuid v1.1.1 // indirect
	github.com/hashicorp/go-getter v0.0.0-20181213035916-be39683deade // indirect
	github.com/hashicorp/go-uuid v1.0.0 // indirect
	github.com/hashicorp/hcl2 v0.0.0-20181215005721-253da47fd604 // indirect
	github.com/hashicorp/hil v0.0.0-20170627220502-fa9f258a9250 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform v0.11.11
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/joho/godotenv v1.2.0
	github.com/juju/loggo v0.0.0-20190526231331-6e530bcce5d8 // indirect
	github.com/juju/testing v0.0.0-20191001232224-ce9dec17d28b // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/mattn/go-shellwords v1.0.6
	github.com/mattn/goveralls v0.0.3
	github.com/mitchellh/cli v1.0.0 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/gox v1.0.1
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/posener/complete v1.2.1 // indirect
	github.com/pseudomuto/protoc-gen-doc v1.2.0
	github.com/pseudomuto/protokit v0.2.0 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.2.1
	github.com/streadway/amqp v0.0.0-20180528204448-e5adc2ada8b8
	github.com/stretchr/testify v1.4.0
	github.com/ugorji/go v1.1.7 // indirect
	github.com/zclconf/go-cty v0.0.0-20181210194043-3a13afc12792 // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0
	golang.org/x/tools v0.0.0-20190614205625-5aca471b1d59
	google.golang.org/grpc v1.27.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/resty.v1 v1.12.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

// workaround for https://github.com/gin-gonic/gin/issues/1673, etcd v3.2.2 dependency
replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go v0.0.0-20181204163529-d75b2dcb6bc8
