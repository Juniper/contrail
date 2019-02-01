package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrailutil"
	"github.com/Juniper/contrail/pkg/log"
)

func main() {
	err := contrailutil.ContrailUtil.Execute()
	if err != nil {
		log.FatalWithStackTrace(err)
	}
}
