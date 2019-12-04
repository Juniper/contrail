package db

import (
	"context"
	"database/sql"

	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
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
	return NewService(sqlDB), nil
}

//NewService makes a DB service.
func NewService(db *sql.DB) *Service {
	dbService := &Service{
		BaseService: services.BaseService{},
		BaseDB:      basedb.NewBaseDB(db),
	}
	dbService.initQueryBuilders()
	return dbService
}

// delete deletes a resource
func (db *Service) delete(
	ctx context.Context,
	qb *basedb.QueryBuilder,
	uuid string,
	backrefFields map[string][]string,
) error {
	if err := db.checkPolicy(ctx, qb, uuid); err != nil {
		return err
	}

	tx := basedb.GetTransaction(ctx)

	for backref := range backrefFields {
		_, err := tx.ExecContext(ctx, qb.DeleteRelaxedBackrefsQuery(backref), uuid)
		if err != nil {
			return errors.Wrapf(
				basedb.FormatDBError(err),
				"deleting all relaxed references from %s to resource with UUID '%v' from DB failed",
				backref, uuid,
			)
		}
	}

	_, err := tx.ExecContext(ctx, qb.DeleteQuery(), uuid)
	if err != nil {
		err = basedb.FormatDBError(err)
		return errors.Wrapf(err, "deleting resource with UUID '%v' from DB failed", uuid)
	}

	return db.DeleteMetadata(ctx, uuid)
}

func (db *Service) count(
	ctx context.Context, qb *basedb.QueryBuilder, spec *baseservices.ListSpec,
) (count int64, err error) {
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

	var row *sql.Row
	if auth.IsAdmin() {
		row = tx.QueryRowContext(ctx, selectQuery, uuid)
	} else {
		row = tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
	}

	var count int
	err = row.Scan(&count)
	if err != nil {
		return basedb.FormatDBError(err)
	}
	if count == 0 {
		return errutil.ErrorNotFound
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

	return errors.Wrapf(
		basedb.FormatDBError(err), "creating resource %T with UUID '%v' in DB failed", obj, obj.GetUUID(),
	)
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
		return errors.Wrapf(
			basedb.FormatDBError(err),
			"%s_ref create failed for object %s with UUID: '%v' and ref UUID '%v'",
			toSchemaID, fromSchemaID, fromID, toID,
		)
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

	return errors.Wrapf(
		basedb.FormatDBError(err),
		"%s_ref delete failed for object %s with UUID: '%v' and ref UUID '%v'",
		toSchemaID, fromSchemaID, fromID, toID,
	)
}
