package main

import (
	"log"

	"github.com/Juniper/contrail/pkg/cmd"
)

func main() {
	err := cmd.ContrailUtilCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
