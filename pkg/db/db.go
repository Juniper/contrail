package db // nolint: golint

import (
	"context"
	"database/sql"
	"strings"

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

// Dump selects all data from every table and writes each row to ObjectWriter.
//
// Note that dumping the whole database using SELECT statements may take a lot
// of time and memory, increasing both server and database load thus it should
// be used as a first shot operation only.
//
// An example application of that function is loading initial database snapshot
// in Watcher.
func (db *Service) Dump(ctx context.Context, ow basedb.ObjectWriter) error {
	return db.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			return db.dump(ctx, ow)
		},
	)
}

type structure map[string]interface{}

func (s *structure) getPaths(prefix string) []string {
	var paths []string
	for k, v := range *s {
		p := prefix + "." + k
		switch o := v.(type) {
		case struct{}:
			paths = append(paths, p)
		case *structure:
			paths = append(paths, o.getPaths(p)...)
		}
	}
	return paths
}

func (s *structure) getChildPaths(path string) (paths []string) {
	innerStructure := s
	for _, segment := range strings.Split(path, ".") {
		switch o := (*innerStructure)[segment].(type) {
		case struct{}:
			return []string{path}
		case *structure:
			innerStructure = o
		}
	}
	return innerStructure.getPaths(path)
}
