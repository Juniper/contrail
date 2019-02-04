package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/cmd/contrailschema"
)

func main() {
	err := contrailschema.ContrailSchema.Execute()
	if err != nil {
		logrus.Fatal(err)
	}
}
