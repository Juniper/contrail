package types

import (
	"context"
	"strconv"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateForwardingClass performs type specific validation.
func (sv *ContrailTypeLogicService) CreateForwardingClass(
	ctx context.Context, request *services.CreateForwardingClassRequest,
) (*services.CreateForwardingClassResponse, error) {

	var response *services.CreateForwardingClassResponse
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			fcID := request.GetForwardingClass().GetForwardingClassID()
			err := checkForwardingClassID(ctx, sv, fcID)
			if err != nil {
				return err
			}

			response, err = sv.BaseService.CreateForwardingClass(ctx, request)
			return err
		})
	return response, err
}

// UpdateForwardingClass performs type specific validation.
func (sv *ContrailTypeLogicService) UpdateForwardingClass(
	ctx context.Context,
	request *services.UpdateForwardingClassRequest,
) (*services.UpdateForwardingClassResponse, error) {
	var response *services.UpdateForwardingClassResponse

	forwardingClass := request.GetForwardingClass()
	fm := request.GetFieldMask()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			databaseFC, err := sv.getForwardingClass(ctx, forwardingClass.GetUUID())
			if err != nil {
				return err
			}

			reqFrwdClassID := forwardingClass.GetForwardingClassID()
			if basemodels.FieldMaskContains(&fm, models.ForwardingClassFieldForwardingClassID) &&
				databaseFC.GetForwardingClassID() != reqFrwdClassID {

				err = checkForwardingClassID(ctx, sv, reqFrwdClassID)
				if err != nil {
					return err
				}
			}

			response, err = sv.BaseService.UpdateForwardingClass(ctx, request)
			return err
		})
	return response, err
}

func checkForwardingClassID(ctx context.Context, sv *ContrailTypeLogicService, fcID int64) error {
	listForwardingClassResponse, err := sv.ReadService.ListForwardingClass(ctx, &services.ListForwardingClassRequest{
		Spec: &baseservices.ListSpec{
			Filters: []*baseservices.Filter{
				{
					Key:    models.ForwardingClassFieldForwardingClassID,
					Values: []string{strconv.FormatInt(fcID, 10)},
				},
			},
			Fields: []string{models.ForwardingClassFieldDisplayName},
		},
	})
	if err != nil {
		return err
	}

	if len(listForwardingClassResponse.ForwardingClasss) > 0 {
		return errutil.ErrorBadRequestf("Forwarding class %s is configured with a id %d",
			listForwardingClassResponse.ForwardingClasss[0].DisplayName, fcID)
	}

	return nil
}

func (sv *ContrailTypeLogicService) getForwardingClass(
	ctx context.Context, id string,
) (*models.ForwardingClass, error) {

	forwardingClassResponse, err := sv.ReadService.GetForwardingClass(
		ctx,
		&services.GetForwardingClassRequest{
			ID: id,
		},
	)

	return forwardingClassResponse.GetForwardingClass(), err
}
