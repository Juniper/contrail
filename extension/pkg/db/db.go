package db

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/extension/pkg/services"
	"github.com/Juniper/contrail/pkg/db/basedb"
)

//Service struct
type Service struct {
	services.BaseService
	basedb.BaseDB
}

//NewServiceFromConfig makes db service from viper config.
func NewServiceFromConfig() (*Service, error) {
	sqlDB, err := basedb.ConnectDB()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect database")
	}
	return NewService(sqlDB, viper.GetString("database.dialect")), nil
}

//NewService makes a DB service.
func NewService(db *sql.DB, dialect string) *Service {
	dbService := &Service{
		BaseDB: basedb.NewBaseDB(db, dialect),
	}
	dbService.initQueryBuilders()
	return dbService
}
