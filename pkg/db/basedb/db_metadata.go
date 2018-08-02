package basedb

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/common"
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
	if uuid == "" && len(fqName) == 0 {
		return nil, fmt.Errorf("uuid and fqName unspecified")
	}

	metadatas, err := db.ListMetadata(ctx, []*basemodels.FQNameUUIDPair{
		{FQName: fqName, UUID: uuid},
	})

	if err != nil {
		return nil, errors.Wrapf(FormatDBError(err), "failed to get metadata")
	}

	if len(metadatas) == 1 {
		return metadatas[0], nil
	}

	return nil, common.ErrorNotFound
}

func (db *BaseDB) buildFilter(columns []string, filterValues ...[]string) string {
	var where string
	var filterQuery bytes.Buffer
	index := 0
	WriteStrings(&filterQuery, " false")
	for i, column := range columns {
		if len(filterValues[i]) > 0 {
			WriteStrings(&filterQuery, " or ")
			WriteStrings(&filterQuery, column, " in (")
			valuesQuery := ""
			values := filterValues[i]
			valuesQuery, index = db.Dialect.ValuesWithIndex(index, values...)
			WriteStrings(&filterQuery, valuesQuery)
			WriteStrings(&filterQuery, ")")
		}
	}
	where = filterQuery.String()

	return where
}

func (db *BaseDB) stringsToValues(strings []string) []interface{} {
	var values []interface{}
	for _, s := range strings {
		values = append(values, s)
	}
	return values
}

// ListMetadata gets metadata from database.
func (db *BaseDB) ListMetadata(
	ctx context.Context, fqNameUUIDPairs []*basemodels.FQNameUUIDPair,
) ([]*basemodels.MetaData, error) {
	var metadatas []*basemodels.MetaData

	if err := db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		var query bytes.Buffer
		query.WriteString("select uuid,type, fq_name from metadata where ")
		var fqNames []string
		var uuids []string

		for _, pair := range fqNameUUIDPairs {
			if len(pair.FQName) > 0 {
				fqNames = append(fqNames, basemodels.FQNameToString(pair.FQName))
			}
			if pair.UUID != "" {
				uuids = append(uuids, pair.UUID)
			}
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
			metadata := &basemodels.MetaData{}
			fqNameString := ""
			err := rows.Scan(&metadata.UUID, &metadata.Type, &fqNameString)
			err = FormatDBError(err)
			if err != nil {
				return errors.Wrap(err, "couldn't get metadatas")
			}
			metadata.FQName = basemodels.ParseFQName(fqNameString)
			metadatas = append(metadatas, metadata)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return metadatas, nil
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
