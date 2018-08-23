package basedb

import (
	"database/sql"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	mysqlUniqueViolation = 1062
	mysqlRowIsReferenced = 1451
	mysqlNoReferencedRow = 1452

	pgUniqueViolation     = "unique_violation"
	pgForeignKeyViolation = "foreign_key_violation"
)

//FormatDBError converts DB specific error.
func FormatDBError(err error) error {
	if err == nil {
		return nil
	}
	if err, ok := err.(*mysql.MySQLError); ok {
		publicErr, ok := convertMySQLError(err)
		if ok {
			log.Debugf("MySQL error: [%d] %s. Returning: %v", err.Number, err.Message, publicErr)
			return publicErr
		}
		log.Errorf("Unknown MySQL error: [%d] %s", err.Number, err.Message)
	}
	if err, ok := err.(*pq.Error); ok {
		publicErr, ok := convertPGError(err)
		if ok {
			log.Debugf("PostgreSQL error: %v. Returning: %v", err, publicErr)
			return publicErr
		}
		log.Error("Unknown PostgreSQL error:", err)
	}
	if err == sql.ErrNoRows {
		return common.ErrorNotFound
	}
	return err
}

func convertMySQLError(err *mysql.MySQLError) (error, bool) {
	switch err.Number {
	case mysqlUniqueViolation:
		return uniqueConstraintViolation(), true
	case mysqlRowIsReferenced:
		return foreignKeyConstraintViolation(), true
	case mysqlNoReferencedRow:
		return foreignKeyConstraintViolation(), true
	default:
		return nil, false
	}
}

func convertPGError(err *pq.Error) (error, bool) {
	switch err.Code.Name() {
	case pgUniqueViolation:
		return uniqueConstraintViolation(), true
	case pgForeignKeyViolation:
		return foreignKeyConstraintViolation(), true
	default:
		return nil, false
	}
}

func uniqueConstraintViolation() error {
	return common.ErrorConflictf("Resource conflict: unique constraint violation")
}

func foreignKeyConstraintViolation() error {
	return common.ErrorConflictf("Resource conflict: foreign key constraint violation")
}
