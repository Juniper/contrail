package main

import (
	"log"

	"github.com/Juniper/contrail/pkg/cmd/contrail"
)

func main() {
	err := contrail.Cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
