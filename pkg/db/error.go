package db

import (
	"github.com/Juniper/contrail/pkg/common"
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	mysqlUniqueViolation = 1062
	pgUniqueViolation    = "unique_violation"
)

func handleError(err error) error {
	if err == nil {
		return nil
	}
	if err, ok := err.(*mysql.MySQLError); ok {
		switch err.Number {
		case mysqlUniqueViolation:
			return common.ErrorConflict
		}
		log.Debug("mysql error:", err.Number, err.Message)
	}
	if err, ok := err.(*pq.Error); ok {
		switch err.Code.Name() {
		case pgUniqueViolation:
			return common.ErrorConflict
		}
		log.Debug("pq error:", err.Code.Name())
	}
	return err
}
