package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateTagType performs Tag Type specific checks.
func (sv *ContrailTypeLogicService) CreateTagType(
	ctx context.Context, request *services.CreateTagTypeRequest,
) (*services.CreateTagTypeResponse, error) {
	var response *services.CreateTagTypeResponse

	if err := sv.validateTagType(request.GetTagType()); err != nil {
		return response, err
	}

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error
			response, err = sv.Next().CreateTagType(ctx, request)

			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) validateTagType(tt *models.TagType) error {
	if tt.TagTypeID != "" {
		return errutil.ErrorBadRequest("Tag Type ID is not settable")
	}

	return nil
}
