package types

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	tagTypeIDPoolKey = "tag_type_id"
)

// CreateTag validates if there is no TagID in request and allocates new one,
// finally creates new Tag.
func (sv *ContrailTypeLogicService) CreateTag(
	ctx context.Context,
	request *services.CreateTagRequest,
) (*services.CreateTagResponse, error) {
	var response *services.CreateTagResponse
	tag := request.GetTag()
	if err := sv.validateTag(tag); err != nil {
		return response, err
	}

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			id, err := sv.IntPoolAllocator.AllocateInt(ctx, tagTypeIDPoolKey)
			if err != nil {
				return err
			}
			tag.TagID = fmt.Sprintf("%d", id)
			response, err = sv.Next().CreateTag(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) validateTag(tag *models.Tag) error {
	if tag.TagID != "" {
		return errutil.ErrorBadRequest("Tag ID is not settable")
	}

	if len(tag.TagTypeRefs) > 0 {
		return errutil.ErrorBadRequest("Tag Type reference is not settable")
	}

	if tag.TagTypeName == "" || tag.TagValue == "" {
		return errutil.ErrorBadRequestf(
			"Tag must be created with a type name and a value but got: type '%v', value '%v'",
			tag.TagTypeName, tag.TagValue,
		)
	}

	return nil
}
