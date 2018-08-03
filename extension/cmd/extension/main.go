package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/cmd/contrail"

	"github.com/Juniper/contrail/extension/pkg/api"
)

func main() {
	apisrv.RegisterExtension(api.Init)
	err := contrail.Contrail.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
