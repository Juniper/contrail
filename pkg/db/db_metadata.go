package db

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

// MetaData represents resource meta data.
//
// TODO(Michal): services.ContrailService relying on this struct caused circular
// 	dependency. Move MetaData definition to db package after extracting metadata
// 	dependent code from service package.
type MetaData = services.MetaData

// CreateMetaData creates fqname, uuid pair with type.
func (db *Service) CreateMetaData(ctx context.Context, metaData *MetaData) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		_, err := tx.Exec(
			"insert into metadata (uuid,type,fq_name) values ("+
				db.Dialect.values("uuid", "type", "fq_name")+");",
			metaData.UUID, metaData.Type, models.FQNameToString(metaData.FQName))
		err = handleError(err)
		return errors.Wrap(err, "failed to create metadata")
	})
}

// GetMetaData gets metadata from database.
func (db *Service) GetMetaData(ctx context.Context, uuid string, fqName []string) (*MetaData, error) {
	var uuidString, typeString, fqNameString string

	if err := db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		var query bytes.Buffer
		query.WriteString("select uuid,type, fq_name from metadata where ")
		var row *sql.Row
		var where string
		if uuid != "" {
			where = "uuid = " + db.Dialect.placeholder(1)
			query.WriteString(where)
			row = tx.QueryRow(query.String(), uuid)
		} else if fqName != nil {
			where = "fq_name = " + db.Dialect.placeholder(1)
			query.WriteString(where)
			row = tx.QueryRow(query.String(), models.FQNameToString(fqName))
		} else {
			return fmt.Errorf("uuid and fqName unspecified ")
		}
		err := row.Scan(&uuidString, &typeString, &fqNameString)
		if err != nil {
			return errors.Wrapf(handleError(err), "failed to get metadata (%v")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &MetaData{
		UUID:   uuidString,
		FQName: models.ParseFQName(fqNameString),
		Type:   typeString,
	}, nil
}

// DeleteMetaData deletes metadata by uuid.
func (db *Service) DeleteMetaData(ctx context.Context, uuid string) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		_, err := tx.Exec("delete from metadata where uuid = "+db.Dialect.placeholder(1), uuid)
		err = handleError(err)
		return errors.Wrap(err, "failed to delete metadata")
	})
}
