package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrailschema"
	"github.com/Juniper/asf/pkg/logutil"
)

func main() {
	err := contrailschema.ContrailSchema.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
