package types

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/services"
	"strconv"
)

const (
	tagTypeIDPoolKey = "tag_type_id"
)

// CreateTagType validates if TagTypeID is not set, allocates one and creates new TagType.
func (sv *ContrailTypeLogicService) CreateTagType(
	ctx context.Context,
	request *services.CreateTagTypeRequest,
) (*services.CreateTagTypeResponse, error) {
	var response *services.CreateTagTypeResponse
	tagType := request.GetTagType()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			tagTypeFQ := tagType.GetFQName()
			tagTypeName := tagTypeFQ[len(tagTypeFQ) - 1]
			tagType.Name = tagTypeName
			tagType.DisplayName = tagTypeName

			if tagType.GetTagTypeID() != "" {
				return errutil.ErrorBadRequestf("Tag Type ID is not setable")
			}

			id, err := sv.IntPoolAllocator.AllocateInt(ctx, tagTypeIDPoolKey)
			if err != nil {
				return err
			}
			tagType.TagTypeID = fmt.Sprintf("0x%04x", id)

			response, err = sv.BaseService.CreateTagType(ctx, request)
			return err
		})

	return response, err
}

// DeleteTagType deallocates TagTypeID and deletes TagType.
func (sv *ContrailTypeLogicService) DeleteTagType(
	ctx context.Context,
	request *services.DeleteTagTypeRequest,
) (*services.DeleteTagTypeResponse, error) {
	var response *services.DeleteTagTypeResponse
	
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			tagTypeResponse, err := sv.ReadService.GetTagType(ctx, &services.GetTagTypeRequest{ID: request.GetID()})
			if err != nil {
				errutil.ErrorNotFoundf("Cannot find tag type with ID: '%s'", request.GetID())
			}

			val, err := strconv.ParseInt(tagTypeResponse.TagType.TagTypeID, 0, 64)
			if err != nil {
				return errutil.ErrorBadRequestf("Convert hex to int error for (%v): %s", tagTypeResponse.TagType.TagTypeID, err)
			}

			if err := sv.IntPoolAllocator.DeallocateInt(ctx, tagTypeIDPoolKey, val); err != nil {
				return errutil.ErrorBadRequestf("Cannot deallocate int (%v): %s", val, err)
			}

			response, err = sv.BaseService.DeleteTagType(ctx, request)
			return err
		})

	return response, err
}
