package db

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
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

func (db *Service) count(ctx context.Context, qb *basedb.QueryBuilder, spec *baseservices.ListSpec) (count int64, err error) {
	query, values := qb.CountQuery(auth.GetAuthCTX(ctx), spec)

	tx := basedb.GetTransaction(ctx)
	row := tx.QueryRowContext(ctx, query, values...)

	if err = row.Scan(&count); err != nil {
		return 0, errors.Wrap(basedb.FormatDBError(err), "count query failed")
	}

	return count, nil
}

// checkPolicy check ownership of resources.
func (db *Service) checkPolicy(ctx context.Context, qb *basedb.QueryBuilder, uuid string) (err error) {
	tx := basedb.GetTransaction(ctx)
	auth := auth.GetAuthCTX(ctx)

	selectQuery := qb.SelectAuthQuery(auth.IsAdmin())

	var count int

	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return basedb.FormatDBError(err)
		}
		row.Scan(&count)
		if count == 0 {
			return errutil.ErrorNotFound
		}
	} else {
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return basedb.FormatDBError(err)
		}
		row.Scan(&count)
		if count == 0 {
			return errutil.ErrorNotFound
		}
	}
	return nil
}

type childObject interface {
	GetUUID() string
	GetParentUUID() string
	GetParentType() string
}

func (db *Service) createParentReference(
	ctx context.Context,
	obj childObject,
	qb *basedb.QueryBuilder,
	possibleParents []string,
	optional bool,
) (err error) {
	parentSchemaID := basemodels.KindToSchemaID(obj.GetParentType())

	if !format.ContainsString(possibleParents, parentSchemaID) {
		if optional {
			return nil
		}
		return errutil.ErrorBadRequest("invalid parent type")
	}

	tx := basedb.GetTransaction(ctx)
	_, err = tx.ExecContext(ctx, qb.CreateParentRefQuery(parentSchemaID), obj.GetUUID(), obj.GetParentUUID())

	return errors.Wrapf(basedb.FormatDBError(err), "creating resource %T with UUID '%v' in DB failed", obj, obj.GetUUID())
}

func (db *Service) createRef(
	ctx context.Context,
	fromID, toID string,
	fromSchemaID, toSchemaID string,
	attrs ...interface{},
) error {
	qb := db.QueryBuilders[fromSchemaID]
	tx := basedb.GetTransaction(ctx)

	_, err := tx.ExecContext(ctx, qb.CreateRefQuery(toSchemaID), append([]interface{}{fromID, toID}, attrs...)...)
	if err != nil {
		return errors.Wrapf(basedb.FormatDBError(err), "{{ reference.GoName }}Ref create failed for object {{ schema.JSONSchema.GoName }} with UUID: '%v' and ref UUID '%v': ", fromID, toID)
	}
	return nil
}

func (db *Service) deleteRef(
	ctx context.Context,
	fromID, toID string,
	fromSchemaID, toSchemaID string,
) error {
	query := db.QueryBuilders[fromSchemaID].DeleteRefQuery(toSchemaID)
	tx := basedb.GetTransaction(ctx)

	_, err := tx.ExecContext(ctx, query, fromID, toID)

	return errors.Wrapf(basedb.FormatDBError(err), "{{ reference.GoName }}Ref delete failed for object {{ schema.JSONSchema.GoName }} with UUID: '%v' and ref UUID '%v': ", fromID, toID)
}
