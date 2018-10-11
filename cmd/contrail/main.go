package main // nolint: golint

import (
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/cmd/contrail"
)

func main() {
	err := contrail.Contrail.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
