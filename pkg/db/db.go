package db

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/services"
)

//Service for DB.
type Service struct {
	services.BaseService
	basedb.BaseDB
}

//NewServiceFromConfig makes db service from viper config.
func NewServiceFromConfig() (*Service, error) {
	sqlDB, err := basedb.ConnectDB()
	if err != nil {
		return nil, errors.Wrap(err, "Init DB failed")
	}
	return NewService(sqlDB, viper.GetString("database.dialect")), nil
}

//NewService makes a DB service.
func NewService(db *sql.DB, dialect string) *Service {
	dbService := &Service{
		BaseService: services.BaseService{},
		BaseDB:      basedb.NewBaseDB(db, dialect),
	}
	dbService.initQueryBuilders()
	return dbService
}
