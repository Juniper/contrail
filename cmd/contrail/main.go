package main

import (
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/cmd/contrail"
)

func main() {
	if err := contrail.Run(); err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
