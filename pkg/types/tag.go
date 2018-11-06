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
							Key:    models.TagTypeFieldDisplayName,
							Values: []string{tag.TagTypeName},
						},
					},
				},
			})
			if err != nil {
				return err
			}
			var tagType *models.TagType
			if tagTypeResponse.TagTypeCount == 1 {
				tagType = tagTypeResponse.TagTypes[0]
			} else if tagTypeResponse.TagTypeCount > 0 {
				for _, t := range tagTypeResponse.TagTypes {
					if t.DisplayName == tag.TagTypeName {
						tagType = t
						break
					}
				}
			}
			if tagType == nil {
				tagTypeCreated, werr := sv.WriteService.CreateTagType(ctx, &services.CreateTagTypeRequest{
					TagType: &models.TagType{
						FQName: []string{tag.TagTypeName},
					},
				})
				if werr != nil {
					return werr
				}
				tagType = tagTypeCreated.TagType
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

// DeleteTag deallocate int from tag_id pool and then deletes Tag.
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
			tagIDHex := strings.TrimPrefix(tagResponse.Tag.TagID, tagType.TagTypeID)
			val, err := strconv.ParseInt(tagIDHex, 10, 64)
			if err != nil {
				return errutil.ErrorBadRequestf("Convert hex to int error for (%v): %s", tagIDHex, err)
			}
			if aerr := sv.IntPoolAllocator.DeallocateInt(ctx, tagIDPoolKey, val); aerr != nil {
				return errutil.ErrorBadRequestf("Cannot deallocate int (%v): %s", val, aerr)
			}

			response, err = sv.BaseService.DeleteTag(ctx, request)
			return err
		})

	return response, err
}
