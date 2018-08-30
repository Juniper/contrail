package basedb

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/pkg/errors"
)

// CreateMetaData creates fqname, uuid pair with type.
func (db *BaseDB) CreateMetaData(ctx context.Context, metaData *basemodels.MetaData) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		fqNameStr, err := fqNameToString(metaData.FQName)
		if err != nil {
			return errors.Wrap(err, "failed to stringify fq_name")
		}
		_, err = tx.Exec(
			"insert into metadata (uuid,type,fq_name) values ("+
				db.Dialect.Values("uuid", "type", "fq_name")+");",
			metaData.UUID, metaData.Type, fqNameStr)
		err = FormatDBError(err)
		return errors.Wrap(err, "failed to create metadata")
	})
}

// GetMetaData gets metadata from database. It requires type and either uuid or fq_name.
func (db *BaseDB) GetMetaData(
	ctx context.Context,
	m basemodels.MetaData,
) (*basemodels.MetaData, error) {
	metadatas, err := db.ListMetadata(ctx, []*basemodels.MetaData{&m})

	if err != nil {
		return nil, errors.Wrapf(FormatDBError(err), "failed to get metadata")
	}

	if len(metadatas) == 1 {
		return metadatas[0], nil
	}

	return nil, common.ErrorNotFound
}

func (db *BaseDB) buildMetadataFilter(
	metaDatas []*basemodels.MetaData,
) (where string, values []interface{}, err error) {
	index := 0
	var filters []string
	var t, uuid, fqName string
	for _, m := range metaDatas {
		if (m.Type == "" || len(m.FQName) == 0) && m.UUID == "" {
			return "", nil, fmt.Errorf("uuid or pair of fq_name and type is required")
		}

		if m.UUID != "" {
			uuid, index = db.Dialect.ValuesWithIndex(index, m.UUID)
			values = append(values, m.UUID)
			filters = append(filters, " ( uuid = "+uuid+" ) ")
		} else {
			t, index = db.Dialect.ValuesWithIndex(index, m.Type)
			values = append(values, m.Type)
			fqNameStr, err := fqNameToString(m.FQName)
			if err != nil {
				return "", nil, errors.Wrapf(err, "failed to stringify fq_name")
			}
			fqName, index = db.Dialect.ValuesWithIndex(index, fqNameStr)
			values = append(values, fqNameStr)
			filters = append(filters, " ( type = "+t+" and fq_name = "+fqName+" ) ")
		}
	}
	where = strings.Join(filters, " or ")
	return where, values, nil
}

// ListMetadata gets metadata from database.
func (db *BaseDB) ListMetadata(
	ctx context.Context,
	m []*basemodels.MetaData,
) ([]*basemodels.MetaData, error) {
	var metadatas []*basemodels.MetaData

	var query bytes.Buffer
	query.WriteString("select uuid,type, fq_name from metadata where ")

	where, values, err := db.buildMetadataFilter(m)
	if err != nil {
		return nil, err
	}
	query.WriteString(where)

	if err := db.DoInTransaction(ctx, func(ctx context.Context) error {
		tx := GetTransaction(ctx)

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
			fqName, err := ParseFQName(fqNameString)
			if err != nil {
				return err
			}
			metadata.FQName = fqName
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
