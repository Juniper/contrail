package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrail"
	"github.com/Juniper/asf/pkg/logutil"
)

func main() {
	err := contrail.Contrail.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
