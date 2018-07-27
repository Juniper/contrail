package main

import (
	"github.com/Juniper/contrail/extention/pkg/api"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/cmd/contrail"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := contrail.Contrail.Execute()
	apisrv.RegisterExtension(api.Init)
	if err != nil {
		log.Fatal(err)
	}
}
