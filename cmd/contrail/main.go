package main // nolint: golint

import (
	"github.com/Juniper/contrail/pkg/cmd/contrail"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := contrail.Contrail.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
