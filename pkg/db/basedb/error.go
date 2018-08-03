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

//HandleError converts db specific error.
func HandleError(err error) error {
	if err == nil {
		return nil
	}
	if err, ok := err.(*mysql.MySQLError); ok {
		switch err.Number {
		case mysqlUniqueViolation:
			return uniqueConstraintViolation()
		case mysqlRowIsReferenced:
			return foreignKeyConstraintViolation()
		case mysqlNoReferencedRow:
			return foreignKeyConstraintViolation()
		}
		log.Debugf("mysql error: [%d] %s", err.Number, err.Message)
	}
	if err, ok := err.(*pq.Error); ok {
		switch err.Code.Name() {
		case pgUniqueViolation:
			return uniqueConstraintViolation()
		case pgForeignKeyViolation:
			return foreignKeyConstraintViolation()
		}
		log.Debug("pq error:", err)
	}
	if err == sql.ErrNoRows {
		return common.ErrorNotFound
	}
	return err
}

func uniqueConstraintViolation() error {
	return common.ErrorConflictf("Resource conflict: unique constraint violation")
}

func foreignKeyConstraintViolation() error {
	return common.ErrorConflictf("Resource conflict: foreign key constraint violation")
}
