package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrailcli"
	"github.com/Juniper/contrail/pkg/log"
)

func main() {
	err := contrailcli.ContrailCLI.Execute()
	if err != nil {
		log.FatalWithStackTrace(err)
	}
}
