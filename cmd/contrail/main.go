package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrail"
	"github.com/Juniper/contrail/pkg/log"
)

func main() {
	err := contrail.Contrail.Execute()
	if err != nil {
		log.FatalWithStackTrace(err)
	}
}
