package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/cmd/contrailutil"
)

func main() {
	err := contrailutil.ContrailUtil.Execute()
	if err != nil {
		logrus.Fatal(err)
	}
}
