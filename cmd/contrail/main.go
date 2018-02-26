package main

import (
	//Import MySQL DB driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/Juniper/contrail/pkg/cmd/contrail"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := contrail.Contrail.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
