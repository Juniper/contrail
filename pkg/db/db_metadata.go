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
			metaData.UUID, metaData.Type, models.FQNameToString(metaData.FQName))
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
			row = tx.QueryRow(query.String(), models.FQNameToString(fqName))
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
		FQName: models.ParseFQName(fqNameString),
		Type:   typeString,
	}, nil
}

type metadataQueryCtx int

func (db *Service) buildFilter(columns []string, filterValues ...[]string) string {
	var where string
	var filterQuery bytes.Buffer
	index := 0
	for i, column := range columns {
		writeStrings(&filterQuery, column, " in (")
		valuesQuery := ""
		valuesQuery, index = db.Dialect.valuesWithIndex(index, filterValues[i]...)
		writeStrings(&filterQuery, valuesQuery)
		writeStrings(&filterQuery, ")")
		writeStrings(&filterQuery, " or ")
	}
	writeStrings(&filterQuery, " false")
	where = filterQuery.String()

	return where
}

func (db *Service) stringsToValues(strings []string) []interface{} {
	var values []interface{}
	for _, s := range strings {
		values = append(values, &s)
	}
	return values
}

// ListMetadata gets metadata from database.
func (db *Service) ListMetadata(ctx context.Context, fqNameUUIDPairs []*models.FQNameUUIDPair) ([]*models.MetaData, error) {
	var metadatas []*models.MetaData

	if err := db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		var query bytes.Buffer
		query.WriteString("select uuid,type, fq_name from metadata where ")
		var fqNames []string
		var uuids []string

		for _, pair := range fqNameUUIDPairs {
			fqNames = append(fqNames, models.FQNameToString(pair.FQName))
			uuids = append(uuids, pair.UUID)
		}

		where := db.buildFilter([]string{"fq_name", "uuid"}, fqNames, uuids)
		query.WriteString(where)
		var values []interface{}
		values = append(values, db.stringsToValues(fqNames)...)
		values = append(values, db.stringsToValues(uuids)...)

		rows, err := tx.QueryContext(ctx, query.String(), values...)
		if err != nil {
			return err
		}

		for rows.Next() {
			metadata := &models.MetaData{}
			fqNameString := ""
			err := rows.Scan(&metadata.UUID, &metadata.Type, &fqNameString)
			err = handleError(err)
			if err != nil {
				return errors.Wrap(err, "metadata")
			}
			metadata.FQName = models.ParseFQName(fqNameString)
			metadatas = append(metadatas, metadata)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return metadatas, nil
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
