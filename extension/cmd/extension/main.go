package main

import (
	"github.com/Juniper/contrail/extension/pkg/api"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/cmd/contrail"
	log "github.com/sirupsen/logrus"
)

func main() {
	apisrv.RegisterExtension(api.Init)
	err := contrail.Contrail.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
