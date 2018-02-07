package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrailutil"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := contrailutil.ContrailUtil.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
