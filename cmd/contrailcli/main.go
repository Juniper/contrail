package main

import (
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/cmd/contrailcli"
)

func main() {
	err := contrailcli.ContrailCLI.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
