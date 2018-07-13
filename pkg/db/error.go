package db

import (
	"database/sql"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	mysqlUniqueViolation     = 1062
	mysqlForeignKeyViolation = 1451

	pgUniqueViolation     = "unique_violation"
	pgForeignKeyViolation = "foreign_key_violation"
)

func handleError(err error) error {
	if err == nil {
		return nil
	}
	if err, ok := err.(*mysql.MySQLError); ok {
		log.Debugf("mysql error: [%d] %s", err.Number, err.Message)
		switch err.Number {
		case mysqlUniqueViolation, mysqlForeignKeyViolation:
			return common.ErrorConflict
		}
	}
	if err, ok := err.(*pq.Error); ok {
		log.Debug("pq error:", err)
		switch err.Code.Name() {
		case pgUniqueViolation, pgForeignKeyViolation:
			return common.ErrorConflict
		}
	}
	if err == sql.ErrNoRows {
		return common.ErrorNotFound
	}
	return err
}
