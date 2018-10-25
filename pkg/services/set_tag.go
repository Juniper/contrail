package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	strings "strings"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

var (
	// TagTypeNotUniquePerObject contains not unique tag-types per object
	TagTypeNotUniquePerObject = map[string]bool{
		"label": true,
	}
	// TagTypeAuthorizedOnAddressGroup contains authorized on address group tag-types
	TagTypeAuthorizedOnAddressGroup = map[string]bool{
		"label": true,
	}
)

func isTagTypeUniquePerObject(tagType string) bool {
	return !TagTypeNotUniquePerObject[tagType]
}

//TagAttr is a part of set-tag input data.
type TagAttr struct {
	IsGlobal     bool        `json:"is_global"`
	Value        interface{} `json:"value"` // Value could be nil, string or slice of strings
	AddValues    []string    `json:"add_values"`
	DeleteValues []string    `json:"delete_values"`
}

func (t *TagAttr) isDeleteRequest() bool {
	return t.Value == nil && (len(t.AddValues) == 0 && len(t.DeleteValues) == 0)
}

func (t *TagAttr) hasAddValues() bool {
	return len(t.AddValues) > 0
}
func (t *TagAttr) hasDeleteValues() bool {
	return len(t.DeleteValues) > 0
}

// SetTagRequest represents set-tag input data.
type SetTagRequest struct {
	ObjUUID string `json:"obj_uuid"`
	ObjType string `json:"obj_type"`
	Tags    map[string]TagAttr
}

func (t *SetTagRequest) validate() error {
	if t.ObjUUID == "" || t.ObjType == "" {
		return errutil.ErrorBadRequestf(
			"both obj_uuid and obj_type should be specified but got uuid: '%s' and type: '%s",
			t.ObjUUID, t.ObjType,
		)
	}
	for tagType, tagAttr := range t.Tags {
		if err := t.validateTagAttr(tagType, tagAttr); err != nil {
			return err
		}
	}
	return nil
}

func (t *SetTagRequest) validateTagAttr(tagType string, tagAttr TagAttr) error {
	tagType = strings.ToLower(tagType)

	// address-group object can only be associated with label
	if t.ObjType == "address_group" && !TagTypeAuthorizedOnAddressGroup[tagType] {
		return errutil.ErrorBadRequestf(
			"invalid tag type %v for object type %v", tagType, t.ObjType,
		)
	}
	if isTagTypeUniquePerObject(tagType) {
		if len(tagAttr.AddValues) > 0 || len(tagAttr.DeleteValues) > 0 {
			return errutil.ErrorBadRequestf(
				"tag type %v cannot be set multiple times on a same object", tagType,
			)
		}

		if _, ok := tagAttr.Value.(string); !ok && !tagAttr.isDeleteRequest() {
			return errutil.ErrorBadRequestf("no valid value provided for tag type %v", tagType)
		}
	}
	return nil
}

func (t *SetTagRequest) parseObjFields(rawJSON map[string]json.RawMessage) error {
	if err := parseField(rawJSON, "obj_uuid", &t.ObjUUID); err != nil {
		return err
	}
	if err := parseField(rawJSON, "obj_type", &t.ObjType); err != nil {
		return err
	}

	return nil
}

func parseField(rawJSON map[string]json.RawMessage, key string, dst interface{}) error {
	if val, ok := rawJSON[key]; ok {
		if err := json.Unmarshal(val, dst); err != nil {
			return errutil.ErrorBadRequestf("invalid '%s' format: %v", key, err)
		}
		delete(rawJSON, key)
	}
	return nil
}

func (t *SetTagRequest) parseTagAttrs(rawJSON map[string]json.RawMessage) error {
	t.Tags = make(map[string]TagAttr)
	for key, val := range rawJSON {
		var tagAttr TagAttr
		if err := json.Unmarshal(val, &tagAttr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid '%v' format: %v", key, err))
		}
		t.Tags[key] = tagAttr
	}
	return nil
}

func (t *SetTagRequest) tagRefEvent(tagUUID string, operation RefOperation) (*Event, error) {
	return NewEventFromRefUpdate(&RefUpdate{
		Operation: operation,
		Type:      t.ObjType,
		UUID:      t.ObjUUID,
		RefType:   models.KindTag,
		RefUUID:   tagUUID,
	})
}

// RESTSetTag handles set-tag request.
func (service *ContrailService) RESTSetTag(c echo.Context) error {
	var rawJSON map[string]json.RawMessage
	if err := c.Bind(&rawJSON); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	setTag := SetTagRequest{}
	if err := setTag.parseObjFields(rawJSON); err != nil {
		return errutil.ToHTTPError(err)
	}
	if err := setTag.parseTagAttrs(rawJSON); err != nil {
		return errutil.ToHTTPError(err)
	}

	ctx := c.Request().Context()
	err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		return service.SetTag(ctx, setTag)
	})

	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// SetTag allows setting tags based on SetTagRequest.
func (service *ContrailService) SetTag(ctx context.Context, setTag SetTagRequest) error {
	if err := setTag.validate(); err != nil {
		return err
	}

	obj, err := GetObject(ctx, service.Next(), setTag.ObjType, setTag.ObjUUID)
	if err != nil {
		return errutil.ErrorBadRequestf(
			"error: %v, while getting %v with UUID %v", err, setTag.ObjType, setTag.ObjUUID,
		)
	}

	references := obj.GetTagReferences()

	for tagType, tagAttr := range setTag.Tags {
		tagType = strings.ToLower(tagType)
		if references, err = service.handleTagAttr(ctx, tagAttr, tagType, obj, references); err != nil {
			return err
		}
	}
	e, err := NewEvent(&EventOption{
		Data:      map[string]interface{}{"tag_refs": references.Unique()},
		Kind:      obj.Kind(),
		UUID:      obj.GetUUID(),
		Operation: OperationUpdate,
	})
	if err != nil {
		return err
	}

	_, err = e.Process(ctx, service)

	return err
}

func (service *ContrailService) handleTagAttr(
	ctx context.Context, tagAttr TagAttr, tagType string, obj basemodels.Object, refs basemodels.References,
) (basemodels.References, error) {
	switch {
	case tagAttr.isDeleteRequest():
		return removeTagsOfType(refs, tagType), nil
	case isTagTypeUniquePerObject(tagType):
		refs = removeTagsOfType(refs, tagType)

		tagValue := format.InterfaceToString(tagAttr.Value)
		uuid, err := service.getTagUUIDInScope(ctx, tagType, tagValue, tagAttr.IsGlobal, obj)

		return append(refs, basemodels.NewReference(uuid, models.KindTag)), err
	case tagAttr.hasAddValues():
		for _, tagValue := range tagAttr.AddValues {
			uuid, err := service.getTagUUIDInScope(ctx, tagType, tagValue, tagAttr.IsGlobal, obj)
			if err != nil {
				return nil, err
			}

			refs = append(refs, basemodels.NewReference(uuid, models.KindTag))
		}
		return refs, nil
	case tagAttr.hasDeleteValues():
		toDelete := map[string]bool{}
		for _, tagValue := range tagAttr.DeleteValues {
			uuid, err := service.getTagUUIDInScope(ctx, tagType, tagValue, tagAttr.IsGlobal, obj)
			if err != nil {
				return nil, err
			}

			toDelete[uuid] = true
		}
		return refs.Filter(func(r basemodels.Reference) bool {
			return !toDelete[r.GetUUID()]
		}), nil
	default:
		return refs, nil
	}
}

func removeTagsOfType(r basemodels.References, tagType string) basemodels.References {
	return r.Filter(func(ref basemodels.Reference) bool {
		tType, _ := models.TagTypeValueFromFQName(ref.GetTo())
		return tType != tagType
	})
}

// TagLocator is an object that references a tag and helps determining tag scope.
type TagLocator interface {
	GetUUID() string
	GetFQName() []string
	GetPerms2() *models.PermType2
	GetParentType() string
	Kind() string
}

func cannotDetermineTagScopeError(tagName string) error {
	return errutil.ErrorNotFoundf("Not able to determine the scope of the tag '%s'", tagName)
}

func (service *ContrailService) getTagFQNameInScope(
	ctx context.Context, tagName string, isGlobal bool, obj basemodels.Object,
) ([]string, error) {
	tl, ok := obj.(TagLocator)
	if !ok {
		return nil, cannotDetermineTagScopeError(tagName)
	}

	switch {
	case isGlobal:
		return []string{tagName}, nil
	case tl.Kind() == "project":
		return basemodels.ChildFQName(tl.GetFQName(), tagName), nil
	case tl.GetParentType() == "project" && len(tl.GetFQName()) > 1:
		fqName := tl.GetFQName()
		fqName[len(fqName)-1] = tagName
		return fqName, nil
	case tl.GetPerms2() != nil:
		data, err := service.MetadataGetter.GetMetadata(
			ctx, basemodels.Metadata{UUID: tl.GetPerms2().GetOwner()},
		)
		if err != nil {
			return nil, errutil.ErrorNotFoundf("cannot find %s %s owner: %v", tagName, tl.GetUUID(), err)
		}
		return basemodels.ChildFQName(data.FQName, tagName), nil
	default:
		return nil, cannotDetermineTagScopeError(tagName)
	}
}

func (service *ContrailService) getTagUUIDInScope(
	ctx context.Context, tagType, tagValue string, isGlobal bool, obj basemodels.Object,
) (string, error) {
	tagName := models.CreateTagName(tagType, tagValue)

	fqName, err := service.getTagFQNameInScope(ctx, tagName, isGlobal, obj)
	if err != nil {
		return "", err
	}

	m, err := service.MetadataGetter.GetMetadata(
		ctx, basemodels.Metadata{FQName: fqName, Type: models.KindTag},
	)
	if err != nil {
		return "", errutil.ErrorNotFoundf("not able to determine the scope of the tag %s: %v", tagName, err)
	}
	return m.UUID, nil
}
