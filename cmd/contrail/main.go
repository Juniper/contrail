package main

import (
	"log"
	//Import MySQL DB driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/Juniper/contrail/pkg/cmd/contrail"
)

func main() {
	err := contrail.Cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
