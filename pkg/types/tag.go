package types

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gogo/protobuf/types"

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
		func(ctx context.Context) (err error) {
			var tagType *models.TagType
			tagType, err = sv.findTagTypeByName(ctx, tag.TagTypeName)
			if err != nil {
				return err
			}

			if tagType == nil {
				tagType, err = sv.createTagTypeFromName(ctx, tag.TagTypeName)
				if err != nil {
					return err
				}
			}

			tag.TagID, err = sv.allocateTagID(ctx, tagType.TagTypeID)
			if err != nil {
				return err
			}

			response, err = sv.Next().CreateTag(ctx, request)
			if err != nil {
				return err
			}

			tagType.AddTagRef(&models.TagTypeTagRef{
				UUID: response.GetTag().GetUUID(),
			})
			_, err = sv.WriteService.UpdateTagType(ctx, &services.UpdateTagTypeRequest{
				TagType: tagType,
				FieldMask: types.FieldMask{
					Paths: []string{models.TagTypeFieldTagRefs},
				},
			})
			if err != nil {
				return err
			}

			return nil
		})

	return response, err
}

// DeleteTag deallocate int from tag_id pool and then deletes Tag.
func (sv *ContrailTypeLogicService) DeleteTag(
	ctx context.Context,
	request *services.DeleteTagRequest,
) (*services.DeleteTagResponse, error) {
	var response *services.DeleteTagResponse
	var tagTypeName string
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) (err error) {
			tagResponse, err := sv.ReadService.GetTag(ctx, &services.GetTagRequest{ID: request.GetID()})
			// Tag doesn't exist
			if tagResponse == nil {
				return nil
			}

			// Don't de-allocate ID and remove pre-defined tag types
			if _, ok := models.TagTypeIDs[tagResponse.Tag.TagTypeName]; ok {
				response = &services.DeleteTagResponse{ID: request.ID}
				return nil
			}

			tagTypeName = tagResponse.Tag.TagTypeName
			tagType, err := sv.findTagTypeByName(ctx, tagTypeName)
			if err != nil {
				return err
			}
			if tagType == nil {
				// TODO(pawel.zadrozny) Risky edge case. DeleteTagType should take care of related tags deletion
				// TagID consist of allocated TagInt, and allocated TagTypeInt converted to hex.
				// Without TagTypeID there is no way to separate them, and safely deallocate TagInt.
				return errutil.ErrorNotFoundf(
					"Cannot delete tag because related tag_type '%s' doesn't exist",
					tagResponse.GetTag().TagTypeName,
				)
			}

			err = sv.deallocateTagID(ctx, tagResponse.Tag.TagID, tagType.TagTypeID)
			if err != nil {
				return err
			}

			tagType.RemoveTagRef(&models.TagTypeTagRef{
				UUID: tagResponse.GetTag().GetUUID(),
			})
			_, err = sv.WriteService.UpdateTagType(ctx, &services.UpdateTagTypeRequest{
				TagType: tagType,
				FieldMask: types.FieldMask{
					Paths: []string{models.TagTypeFieldTagRefs},
				},
			})
			if err != nil {
				return err
			}

			response, err = sv.Next().DeleteTag(ctx, request)
			if err != nil {
				return err
			}

			// Try to delete referenced tag-type if no references left and ignore constraint violation error.
			if len(tagType.GetReferences()) == 0 {
				_, err = sv.WriteService.DeleteTagType(ctx, &services.DeleteTagTypeRequest{ID: tagType.GetUUID()})
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
	allocInt, err := sv.IntPoolAllocator.AllocateInt(ctx, tagIDPoolKey)
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("%s%04x", tagTypeID, allocInt)
	return id, err
}

func (sv *ContrailTypeLogicService) deallocateTagID(
	ctx context.Context,
	tagIDHex string,
	tagTypeIDHex string,
) error {
	tagID, err := strconv.ParseInt(tagIDHex, 0, 64)
	if err != nil {
		return errutil.ErrorBadRequestf("Convert hex to int error for TagID (%v): %s", tagIDHex, err)
	}

	tagTypeID, err := strconv.ParseInt(tagTypeIDHex, 0, 64)
	if err != nil {
		return errutil.ErrorBadRequestf("Convert hex to int error for TagTypeID (%v): %s", tagTypeIDHex, err)
	}

	id := tagID & tagTypeID
	if err = sv.IntPoolAllocator.DeallocateInt(ctx, tagIDPoolKey, id); err != nil {
		return errutil.ErrorBadRequestf("Cannot deallocate int (%v): %s", id, err)
	}
	return nil
}

func (sv *ContrailTypeLogicService) findTagTypeByName(
	ctx context.Context,
	tagTypeName string,
) (tagType *models.TagType, err error) {
	tagTypeResponse, err := sv.ReadService.ListTagType(ctx, &services.ListTagTypeRequest{
		Spec: &baseservices.ListSpec{
			Detail: true,
			Filters: []*baseservices.Filter{
				{
					Key:    models.TagFieldTagTypeName,
					Values: []string{tagTypeName},
				},
			},
		},
	})
	if err != nil {
		return tagType, err
	}

	if tagTypeResponse.TagTypeCount == 1 {
		tagType = tagTypeResponse.TagTypes[0]
	} else if tagTypeResponse.TagTypeCount > 0 {
		for _, t := range tagTypeResponse.TagTypes {
			if t.DisplayName == tagTypeName {
				tagType = t
				break
			}
		}
	}

	return tagType, err
}

func (sv *ContrailTypeLogicService) createTagTypeFromName(
	ctx context.Context,
	tagTypeName string,
) (tagType *models.TagType, err error) {
	tagTypeCreated, err := sv.WriteService.CreateTagType(ctx, &services.CreateTagTypeRequest{
		TagType: &models.TagType{
			FQName: []string{tagTypeName},
		},
	})
	if err != nil {
		return tagType, err
	}
	tagType = tagTypeCreated.TagType
	return tagType, err
}
