package db

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// Chown changes ownership of an object given its uuid and type.
func (db *Service) Chown(ctx context.Context, uuid string, objectType string, owner string) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		var fm types.FieldMask
		basemodels.FieldMaskAppend(&fm, models.VirtualNetworkFieldPerms2, models.PermType2FieldOwner)

		event, err := services.NewEvent(&services.EventOption{
			UUID:      uuid,
			Kind:      objectType,
			Operation: services.OperationUpdate,
			Data: map[string]interface{}{
				"perms2": map[string]interface{}{
					"owner": owner,
				},
			},
			FieldMask: &fm,
		})
		if err != nil {
			return errors.Wrapf(err, "failed to change the owner of '%v' with UUID '%v'", objectType, uuid)
		}

		_, err = event.Process(ctx, db)
		return errors.Wrapf(err, "failed to change the owner of '%v' with UUID '%v'", objectType, uuid)
	})
}
