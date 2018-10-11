package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/cmd/contrailcli"
)

func main() {
	err := contrailcli.ContrailCLI.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
