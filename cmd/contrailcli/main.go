package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrailcli"
	"github.com/Juniper/contrail/pkg/logutil"
)

func main() {
	err := contrailcli.ContrailCLI.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
