package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/cmd/contrailutil"
)

func main() {
	err := contrailutil.ContrailUtil.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
