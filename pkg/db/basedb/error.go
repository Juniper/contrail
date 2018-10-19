package basedb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/errutil"
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

	if publicErr := getPublicError(err); publicErr != nil {
		log.Debugf("Database error: %v. Returning: %v", err, publicErr)
		return publicErr
	}
	log.Error("Unknown database error:", err)
	return err
}

// getPublicError returns an error with an API error code and a high-level error message.
// If err is not recognized, nil is returned.
func getPublicError(err error) error {
	if err == sql.ErrNoRows {
		return errutil.ErrorNotFound
	}

	switch err.(type) {
	case *mysql.MySQLError:
		return getPublicMySQLError(err.(*mysql.MySQLError))
	case *pq.Error:
		return getPublicPGError(err.(*pq.Error))
	default:
		return nil
	}
}

func getPublicMySQLError(err *mysql.MySQLError) error {
	switch err.Number {
	case mysqlUniqueViolation:
		return uniqueConstraintViolation()
	case mysqlRowIsReferenced:
		return foreignKeyConstraintViolation()
	case mysqlNoReferencedRow:
		return foreignKeyConstraintViolation()
	default:
		return nil
	}
}

func getPublicPGError(err *pq.Error) error {
	switch err.Code.Name() {
	case pgUniqueViolation:
		return uniqueConstraintViolation()
	case pgForeignKeyViolation:
		return foreignKeyConstraintViolation()
	default:
		return nil
	}
}

func uniqueConstraintViolation() error {
	return errutil.ErrorConflictf("Resource conflict: unique constraint violation")
}

func foreignKeyConstraintViolation() error {
	return errutil.ErrorConflictf("Resource conflict: foreign key constraint violation")
}
