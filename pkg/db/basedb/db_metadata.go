package basedb

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/pkg/errors"
)

// CreateMetaData creates fqname, uuid pair with type.
func (db *BaseDB) CreateMetaData(ctx context.Context, metaData *basemodels.MetaData) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		_, err := tx.Exec(
			"insert into metadata (uuid,type,fq_name) values ("+
				db.Dialect.Values("uuid", "type", "fq_name")+");",
			metaData.UUID, metaData.Type, basemodels.FQNameToString(metaData.FQName))
		err = FormatDBError(err)
		return errors.Wrap(err, "failed to create metadata")
	})
}

// GetMetaData gets metadata from database.
func (db *BaseDB) GetMetaData(ctx context.Context, uuid string, fqName []string) (*basemodels.MetaData, error) {
	var uuidString, typeString, fqNameString string

	if err := db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		var query bytes.Buffer
		query.WriteString("select uuid,type, fq_name from metadata where ")
		var row *sql.Row
		var where string
		if uuid != "" {
			where = "uuid = " + db.Dialect.Placeholder(1)
			query.WriteString(where)
			row = tx.QueryRow(query.String(), uuid)
		} else if fqName != nil {
			where = "fq_name = " + db.Dialect.Placeholder(1)
			query.WriteString(where)
			row = tx.QueryRow(query.String(), basemodels.FQNameToString(fqName))
		} else {
			return fmt.Errorf("uuid and fqName unspecified")
		}
		err := row.Scan(&uuidString, &typeString, &fqNameString)
		if err != nil {
			return errors.Wrapf(FormatDBError(err), "failed to get metadata")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &basemodels.MetaData{
		UUID:   uuidString,
		FQName: basemodels.ParseFQName(fqNameString),
		Type:   typeString,
	}, nil
}

// DeleteMetaData deletes metadata by uuid.
func (db *BaseDB) DeleteMetaData(ctx context.Context, uuid string) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		_, err := tx.Exec("delete from metadata where uuid = "+db.Dialect.Placeholder(1), uuid)
		err = FormatDBError(err)
		return errors.Wrap(err, "failed to delete metadata")
	})
}
