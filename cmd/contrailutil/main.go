package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrailutil"
	"github.com/Juniper/contrail/pkg/logutil"
)

func main() {
	err := contrailutil.ContrailUtil.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
