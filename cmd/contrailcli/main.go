package main

import (
	"github.com/Juniper/asf/pkg/cmd/asfcli"
	"github.com/Juniper/asf/pkg/logutil"
)

func main() {
	err := asfcli.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
