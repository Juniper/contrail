package main

import (
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/cmd/contrailutil"
)

func main() {
	err := contrailutil.ContrailUtil.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
