package types

import (
	"context"
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// CreateForwardingClass performs type specific validation
func (sv *ContrailTypeLogicService) CreateForwardingClass(
	ctx context.Context, request *services.CreateForwardingClassRequest,
) (*services.CreateForwardingClassResponse, error) {

	var response *services.CreateForwardingClassResponse
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			fcID := request.GetForwardingClass().ForwardingClassID
			err := checkForwardingClassID(ctx, sv, fcID)
			if err != nil {
				return err
			}

			response, err = sv.Next().CreateForwardingClass(ctx, request)
			return err
		})
	return response, err
}

// UpdateForwardingClass performs type specific validation
func (sv *ContrailTypeLogicService) UpdateForwardingClass(
	ctx context.Context,
	request *services.UpdateForwardingClassRequest,
) (*services.UpdateForwardingClassResponse, error) {

	var response *services.UpdateForwardingClassResponse
	forwardingClass := request.GetForwardingClass()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			databaseFC, err := sv.getForwardingClass(ctx, forwardingClass.GetUUID())
			if err != nil {
				return err
			}

			if databaseFC.ForwardingClassID != forwardingClass.ForwardingClassID {
				err = checkForwardingClassID(ctx, sv, forwardingClass.ForwardingClassID)
				if err != nil {
					return err
				}
			}

			response, err = sv.Next().UpdateForwardingClass(ctx, request)
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
					Values: []string{strconv.Itoa(int(fcID))},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	if listForwardingClassResponse.ForwardingClassCount != 0 {
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
