package db

import (
	"context"
	"strings"

	"github.com/Juniper/asf/pkg/models"
	"github.com/pkg/errors"
)

// RelaxRef makes a reference not prevent the referenced resource from being deleted.
// When the referenced resource gets deleted, the reference will be deleted as well.
func (db *DB) RelaxRef(ctx context.Context, fromUUID, toUUID string) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		fromSchemaID, err := db.getSchemaID(ctx, fromUUID)
		if err != nil {
			return errors.Wrapf(err, "failed to relax reference from %v to %v: failed to get schema ID for `from`",
				fromUUID, toUUID)
		}

		toSchemaID, err := db.getSchemaID(ctx, toUUID)
		if err != nil {
			return errors.Wrapf(err, "failed to relax reference from %v to %v: failed to get schema ID for `to`",
				fromUUID, toUUID)
		}

		if err := db.relaxRef(ctx, fromSchemaID, toSchemaID, fromUUID, toUUID); err != nil {
			return err
		}
		return nil
	})
}

func (db *DB) getSchemaID(ctx context.Context, uuid string) (string, error) {
	metadata, err := db.GetMetadata(ctx, models.Metadata{UUID: uuid})
	if err != nil {
		return "", errors.Wrapf(err, "failed to get metadata for the resource with UUID '%v'", uuid)
	}

	return models.KindToSchemaID(metadata.Type), nil
}

func (db *DB) relaxRef(ctx context.Context, schemaID, linkTo string, fromUUID, toUUID string) error {
	qb := db.QueryBuilders[schemaID]
	tx := GetTransaction(ctx)
	if _, err := tx.ExecContext(ctx, qb.RelaxRefQuery(strings.ToLower(linkTo)), fromUUID, toUUID); err != nil {
		err = FormatDBError(err)
		return errors.Wrapf(err, "failed to relax reference from %v %v to %v %v", schemaID, fromUUID, linkTo, toUUID)
	}
	return nil
}
