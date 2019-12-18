package types

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	tagIDPoolKey  = "tag_id"
	tagTypeIDMask = 0x0000ffff
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
		func(ctx context.Context) (err error) {
			var tagType *models.TagType
			if tagType, err = sv.findOrCreateTagTypeByName(ctx, tag.GetTagTypeName()); err != nil {
				return err
			}

			if tag.TagID, err = sv.allocateTagID(ctx, tagType.GetTagTypeID()); err != nil {
				return err
			}

			if response, err = sv.BaseService.CreateTag(ctx, request); err != nil {
				return err
			}

			tagType.AddTagRef(&models.TagTypeTagRef{
				UUID: tag.GetUUID(),
			})
			_, err = sv.WriteService.UpdateTagType(
				ctx, &services.UpdateTagTypeRequest{
					TagType: tagType,
					FieldMask: types.FieldMask{
						Paths: []string{models.TagTypeFieldTagRefs},
					},
				},
			)

			return err
		})

	return response, err
}

// DeleteTag deallocate int from tag_id pool and then deletes Tag.
func (sv *ContrailTypeLogicService) DeleteTag(
	ctx context.Context,
	request *services.DeleteTagRequest,
) (*services.DeleteTagResponse, error) {
	var response *services.DeleteTagResponse

	id := request.GetID()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) (err error) {
			tagResponse, err := sv.ReadService.GetTag(ctx, &services.GetTagRequest{ID: id})
			if tagResponse == nil {
				// Tag don't exist, skip.
				return errutil.ErrorNotFoundf("No tag: %s", id)
			}

			tag := tagResponse.GetTag()

			tagType, err := sv.findTagTypeByName(ctx, tag.GetTagTypeName())
			if err != nil {
				return err
			}

			if err = sv.removeTagTypeTagRef(ctx, tagType, id); err != nil {
				return err
			}

			if response, err = sv.BaseService.DeleteTag(ctx, request); err != nil {
				return err
			}

			// Don't de-allocate ID and remove pre-defined tag types
			if _, ok := models.TagTypeIDs[tag.GetTagTypeName()]; ok {
				response = &services.DeleteTagResponse{ID: id}
				return nil
			}

			if err = sv.deallocateTagID(ctx, tag.GetTagID()); err != nil {
				return err
			}

			// Try to delete referenced tag-type if no references left.
			if tagType != nil && len(tagType.GetReferences()) == 0 {
				_, err = sv.WriteService.DeleteTagType(
					ctx, &services.DeleteTagTypeRequest{ID: tagType.GetUUID()},
				)
				if err != nil {
					return err
				}
			}

			return nil
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

func (sv *ContrailTypeLogicService) allocateTagID(
	ctx context.Context,
	tagTypeID string,
) (string, error) {
	allocInt, err := sv.IntPoolAllocator.AllocateInt(ctx, tagIDPoolKey, db.EmptyIntOwner)
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("%s%04x", tagTypeID, allocInt)
	return id, err
}

func (sv *ContrailTypeLogicService) deallocateTagID(
	ctx context.Context,
	tagIDHex string,
) error {
	tagID, err := strconv.ParseInt(tagIDHex, 0, 64)
	if err != nil {
		return errutil.ErrorBadRequestf("Convert hex to int error for TagID (%v): %s", tagIDHex, err)
	}

	id := tagID & tagTypeIDMask
	if err = sv.IntPoolAllocator.DeallocateInt(ctx, tagIDPoolKey, id); err != nil {
		return errutil.ErrorBadRequestf("Cannot deallocate int (%v): %s", id, err)
	}
	return nil
}

func (sv *ContrailTypeLogicService) findTagTypeByName(
	ctx context.Context,
	tagTypeName string,
) (tagType *models.TagType, err error) {
	tagTypeFqName, err := json.Marshal([]string{tagTypeName})
	if err != nil {
		return tagType, errors.Errorf("failed to parse fq name to string: %v", err)
	}

	tagTypeResponse, err := sv.ReadService.ListTagType(ctx, &services.ListTagTypeRequest{
		Spec: &baseservices.ListSpec{
			Detail: true,
			Filters: []*baseservices.Filter{
				{
					Key:    models.TagTypeFieldFQName,
					Values: []string{string(tagTypeFqName)},
				},
			},
		},
	})
	if tagTypeResponse.TagTypeCount > 0 {
		tagType = tagTypeResponse.TagTypes[0]
	}

	return tagType, err
}

func (sv *ContrailTypeLogicService) createTagTypeFromName(
	ctx context.Context,
	tagTypeName string,
) (tagType *models.TagType, err error) {
	tagTypeCreated, err := sv.WriteService.CreateTagType(
		ctx, &services.CreateTagTypeRequest{
			TagType: &models.TagType{
				FQName: []string{tagTypeName},
			},
		},
	)
	if err != nil {
		return tagType, err
	}
	tagType = tagTypeCreated.TagType
	return tagType, err
}

func (sv *ContrailTypeLogicService) findOrCreateTagTypeByName(
	ctx context.Context,
	tagTypeName string,
) (tagType *models.TagType, err error) {
	if tagType, err = sv.findTagTypeByName(ctx, tagTypeName); err != nil {
		return tagType, err
	}
	if tagType == nil {
		if tagType, err = sv.createTagTypeFromName(ctx, tagTypeName); err != nil {
			return tagType, err
		}
	}
	return tagType, err
}

func (sv *ContrailTypeLogicService) removeTagTypeTagRef(
	ctx context.Context,
	tagType *models.TagType,
	tagUUID string,
) (err error) {
	if tagType != nil {
		tagType.RemoveTagRef(&models.TagTypeTagRef{UUID: tagUUID})
		_, err = sv.WriteService.UpdateTagType(
			ctx, &services.UpdateTagTypeRequest{
				TagType: tagType,
				FieldMask: types.FieldMask{
					Paths: []string{models.TagTypeFieldTagRefs},
				},
			},
		)
	}
	return err
}
