package main

import (
	"github.com/Juniper/asf/pkg/cmd/asfgen"
	"github.com/Juniper/asf/pkg/logutil"
)

func main() {
	err := asfgen.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
