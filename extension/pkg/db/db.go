package db

import (
	"database/sql"

	"github.com/Juniper/contrail/extension/pkg/services"
	"github.com/Juniper/contrail/pkg/db/basedb"
)

//Service struct
type Service struct {
	services.BaseService
	basedb.BaseDB
}

//NewService makes a DB service.
func NewService(db *sql.DB, dialect string) *Service {
	dbService := &Service{
		BaseDB: basedb.NewBaseDB(db, dialect),
	}
	dbService.initQueryBuilders()
	return dbService
}
