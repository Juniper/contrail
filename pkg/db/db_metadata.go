package db

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

// CreateMetaData creates fqname, uuid pair with type.
func (db *Service) CreateMetaData(ctx context.Context, metaData *models.MetaData) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		_, err := tx.Exec(
			"insert into metadata (uuid,type,fq_name) values ("+
				db.Dialect.values("uuid", "type", "fq_name")+");",
			metaData.UUID, metaData.Type, fqNameToString(metaData.FQName))
		err = handleError(err)
		return errors.Wrap(err, "failed to create metadata")
	})
}

// GetMetaData gets metadata from database.
func (db *Service) GetMetaData(ctx context.Context, uuid string, fqName []string) (*models.MetaData, error) {
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
			row = tx.QueryRow(query.String(), fqNameToString(fqName))
		} else {
			return fmt.Errorf("uuid and fqName unspecified")
		}
		err := row.Scan(&uuidString, &typeString, &fqNameString)
		if err != nil {
			return errors.Wrapf(handleError(err), "failed to get metadata")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &models.MetaData{
		UUID:   uuidString,
		FQName: parseFQName(fqNameString),
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
