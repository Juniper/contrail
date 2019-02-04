package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/cmd/contrailcli"
)

func main() {
	err := contrailcli.ContrailCLI.Execute()
	if err != nil {
		logrus.Fatal(err)
	}
}
