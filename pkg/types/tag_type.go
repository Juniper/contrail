package types

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	tagTypeIDPoolKey = "tag_type_id"
)

// CreateTagType validates if TagTypeID is not set, allocates one and creates new Tag Type.
func (sv *ContrailTypeLogicService) CreateTagType(
	ctx context.Context,
	request *services.CreateTagTypeRequest,
) (*services.CreateTagTypeResponse, error) {
	var response *services.CreateTagTypeResponse
	tt := request.GetTagType()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			if tt.GetTagTypeID() != "" {
				return errutil.ErrorBadRequestf("Tag Type ID is not settable")
			}

			id, err := sv.IntPoolAllocator.AllocateInt(ctx, tagTypeIDPoolKey)
			if err != nil {
				return err
			}
			tt.TagTypeID = fmt.Sprintf("0x%04x", id)

			response, err = sv.BaseService.CreateTagType(ctx, request)
			return err
		})

	return response, err
}

// DeleteTagType deallocates TagTypeID and deletes Tag Type.
func (sv *ContrailTypeLogicService) DeleteTagType(
	ctx context.Context,
	request *services.DeleteTagTypeRequest,
) (*services.DeleteTagTypeResponse, error) {
	var response *services.DeleteTagTypeResponse

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			tt, err := sv.getTagType(ctx, request.GetID())
			if err != nil {
				return err
			}

			val, err := strconv.ParseInt(tt.TagTypeID, 0, 64)
			if err != nil {
				return errutil.ErrorBadRequestf("convert hex to int error for (%v): %s", tt.TagTypeID, err)
			}

			if err = sv.IntPoolAllocator.DeallocateInt(ctx, tagTypeIDPoolKey, val); err != nil {
				return errutil.ErrorBadRequestf("cannot deallocate int (%v): %s", val, err)
			}

			response, err = sv.BaseService.DeleteTagType(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getTagType(
	ctx context.Context,
	id string,
) (*models.TagType, error) {
	tagTypeResponse, err := sv.ReadService.GetTagType(
		ctx,
		&services.GetTagTypeRequest{
			ID: id,
		},
	)

	return tagTypeResponse.GetTagType(), err
}
