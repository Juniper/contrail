package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrailschema"
	"github.com/Juniper/contrail/pkg/log"
)

func main() {
	err := contrailschema.ContrailSchema.Execute()
	if err != nil {
		log.FatalWithStackTrace(err)
	}
}
