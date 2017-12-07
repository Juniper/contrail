package main

import (
	"log"

	"github.com/Juniper/contrail/pkg/cmd/contrailutil"
)

func main() {
	err := contrailutil.Cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
