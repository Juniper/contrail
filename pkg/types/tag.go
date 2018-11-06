package types

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
	tagIDPoolKey = "tag_id"
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
			tagTypeResponse, err := sv.ReadService.ListTagType(ctx, &services.ListTagTypeRequest{
				Spec: &baseservices.ListSpec{
					Filters: []*baseservices.Filter{
						{
							Key:    models.TagFieldTagTypeName,
							Values: []string{tag.TagTypeName},
						},
					},
				},
			})
			// TODO(pawel.zadrozny) Filter doesn't work as expected, output must be filtered manually
			// Context: tagTypeResponse contains all Tag Types even with .Name == ""
			// Remove this loop when filters are fixed
			var tagType *models.TagType
			for _, t := range tagTypeResponse.TagTypes {
				if t.Name == tag.TagTypeName {
					tagType = t
					break
				}
			}
			if tagType == nil {
				return errutil.ErrorNotFoundf("Cannot create tag because tag_type %s doesn't exist")
			}

			id, err := sv.IntPoolAllocator.AllocateInt(ctx, tagIDPoolKey)
			if err != nil {
				return err
			}
			tag.TagID = fmt.Sprintf("%s%x", tagType.TagTypeID, id)
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

func (sv *ContrailTypeLogicService) DeleteTag(
	ctx context.Context,
	request *services.DeleteTagRequest,
) (*services.DeleteTagResponse, error) {
	var response *services.DeleteTagResponse
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			tagResponse, err := sv.ReadService.GetTag(ctx, &services.GetTagRequest{ID: request.GetID()})
			if err != nil {
				errutil.ErrorNotFoundf("Cannot find tag with ID: '%s'", request.GetID())
			}
			tagTypeResponse, err := sv.ReadService.ListTagType(ctx, &services.ListTagTypeRequest{
				Spec: &baseservices.ListSpec{
					Filters: []*baseservices.Filter{
						{
							Key:    models.TagFieldTagTypeName,
							Values: []string{tagResponse.Tag.TagTypeName},
						},
					},
				},
			})
			var tagType *models.TagType
			for _, t := range tagTypeResponse.TagTypes {
				if t.Name == tagResponse.Tag.TagTypeName {
					tagType = t
					break
				}
			}
			if tagType == nil {
				return errutil.ErrorNotFoundf("Cannot create tag because tag_type '%s' doesn't exist")
			}
			tagIdHex := strings.TrimPrefix(tagResponse.Tag.TagID, tagType.TagTypeID)
			val, err := strconv.ParseInt(tagIdHex, 10, 64)
			if err != nil {
				return errutil.ErrorBadRequestf("Convert hex to int error for (%v): %s", tagIdHex, err)
			}
			if err := sv.IntPoolAllocator.DeallocateInt(ctx, tagIDPoolKey, val); err != nil {
				return errutil.ErrorBadRequestf("Cannot deallocate int (%v): %s", val, err)
			}

			response, err = sv.BaseService.DeleteTag(ctx, request)
			return err
		})

	return response, err
}
