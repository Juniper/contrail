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

	ok, publicErr := convertError(err)
	if ok {
		log.Debugf("Database error: %v. Returning: %v", err, publicErr)
		return publicErr
	}
	log.Error("Unknown database error:", err)
	return err
}

func convertError(err error) (ok bool, publicErr error) {
	if err == sql.ErrNoRows {
		return true, common.ErrorNotFound
	}

	switch err.(type) {
	case *mysql.MySQLError:
		return convertMySQLError(err.(*mysql.MySQLError))
	case *pq.Error:
		return convertPGError(err.(*pq.Error))
	default:
		return false, err
	}
}

func convertMySQLError(err *mysql.MySQLError) (ok bool, publicErr error) {
	switch err.Number {
	case mysqlUniqueViolation:
		return true, uniqueConstraintViolation()
	case mysqlRowIsReferenced:
		return true, foreignKeyConstraintViolation()
	case mysqlNoReferencedRow:
		return true, foreignKeyConstraintViolation()
	default:
		return false, err
	}
}

func convertPGError(err *pq.Error) (ok bool, publicErr error) {
	switch err.Code.Name() {
	case pgUniqueViolation:
		return true, uniqueConstraintViolation()
	case pgForeignKeyViolation:
		return true, foreignKeyConstraintViolation()
	default:
		return false, err
	}
}

func uniqueConstraintViolation() error {
	return common.ErrorConflictf("Resource conflict: unique constraint violation")
}

func foreignKeyConstraintViolation() error {
	return common.ErrorConflictf("Resource conflict: foreign key constraint violation")
}
