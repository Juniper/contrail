package main

import (
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/cmd/contrailschema"
)

func main() {
	err := contrailschema.ContrailSchema.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
