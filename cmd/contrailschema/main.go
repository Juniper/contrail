package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/cmd/contrailschema"
)

func main() {
	err := contrailschema.ContrailSchema.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
