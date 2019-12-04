package db

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// RelaxRef makes a reference not prevent the referenced resource from being deleted.
// When the referenced resource gets deleted, the reference will be deleted as well.
func (db *Service) RelaxRef(ctx context.Context, request *services.RelaxRefRequest) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		fromSchemaID, err := db.getSchemaID(ctx, request.GetUUID())
		if err != nil {
			return errors.Wrapf(err, "failed to relax reference from %v to %v: failed to get schema ID for `from`",
				request.GetUUID(), request.GetRefUUID())
		}

		toSchemaID, err := db.getSchemaID(ctx, request.GetRefUUID())
		if err != nil {
			return errors.Wrapf(err, "failed to relax reference from %v to %v: failed to get schema ID for `to`",
				request.GetUUID(), request.GetRefUUID())
		}

		if err := db.relaxRef(ctx, fromSchemaID, toSchemaID, request); err != nil {
			return err
		}
		return nil
	})
}

func (db *Service) getSchemaID(ctx context.Context, uuid string) (string, error) {
	metadata, err := db.GetMetadata(ctx, basemodels.Metadata{
		UUID: uuid,
	})
	if err != nil {
		return "", errors.Wrapf(err, "failed to get metadata for the resource with UUID '%v'", uuid)
	}

	return basemodels.KindToSchemaID(metadata.Type), nil
}

func (db *Service) relaxRef(ctx context.Context, schemaID, linkTo string, request *services.RelaxRefRequest) error {
	qb := db.QueryBuilders[schemaID]
	tx := basedb.GetTransaction(ctx)
	_, err := tx.ExecContext(ctx, qb.RelaxRefQuery(strings.ToLower(linkTo)), request.GetUUID(), request.GetRefUUID())
	if err != nil {
		err = basedb.FormatDBError(err)
		return errors.Wrapf(err, "failed to relax reference from %v %v to %v %v",
			schemaID, request.GetUUID(), linkTo, request.GetRefUUID())
	}
	return nil
}
