package db

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// RelaxRef makes a reference not prevent the referenced resource from being deleted.
// When the referenced resource gets deleted, the reference would be deleted as well.
func (db *Service) RelaxRef(ctx context.Context, request *services.RelaxRefRequest) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		fromSchemaID, err := db.getSchemaID(ctx, request.UUID)
		if err != nil {
			// TODO Add parameters to the error message?
			return errors.Wrap(err, "failed to relax reference: failed to fetch schema ID for `from`")
		}

		toSchemaID, err := db.getSchemaID(ctx, request.RefUUID)
		if err != nil {
			return errors.Wrap(err, "failed to relax reference: failed to fetch schema ID for `to`")
		}

		if err := db.relaxRef(ctx, fromSchemaID, toSchemaID, request); err != nil {
			return errors.Wrap(err, "failed to relax reference")
		}
		return nil
	})
}

func (db *Service) getSchemaID(ctx context.Context, uuid string) (string, error) {
	metadata, err := db.GetMetadata(ctx, basemodels.Metadata{
		UUID: uuid,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to fetch metadata")
	}

	return kindToSchemaID(metadata.Type), nil
}

// TODO Extract this to another package and make other packages use this.
func kindToSchemaID(kind string) string {
	return strings.Replace(kind, "-", "_", -1)
}

func (db *Service) relaxRef(ctx context.Context, schemaID, linkTo string, request *services.RelaxRefRequest) error {
	qb := db.QueryBuilders[schemaID]
	tx := basedb.GetTransaction(ctx)
	_, err := tx.ExecContext(ctx, qb.RelaxRefQuery(strings.ToLower(linkTo)), request.UUID, request.RefUUID)
	if err != nil {
		err = basedb.FormatDBError(err)
		return errors.Wrap(err, "failed to execute query")
	}
	return nil
}
