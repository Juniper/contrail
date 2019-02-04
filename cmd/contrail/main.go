package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/cmd/contrail"
)

func main() {
	err := contrail.Contrail.Execute()
	if err != nil {
		logrus.Fatal(err)
	}
}
