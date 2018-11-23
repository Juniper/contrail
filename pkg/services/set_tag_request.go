package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	strings "strings"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
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

//TagAttr is a part of set-tag input data.
//type TagAttr struct {
//IsGlobal     bool        `json:"is_global"`
//Value        interface{} `json:"value"` // Value could be nil, string or slice of strings
//AddValues    []string    `json:"add_values"`
//DeleteValues []string    `json:"delete_values"`
//}

func (a *SetTagAttr) isDeleteRequest() bool {
	return a.Value == nil && (len(a.AddValues) == 0 && len(a.DeleteValues) == 0)
}

func (a *SetTagAttr) hasAddValues() bool {
	return len(a.AddValues) > 0
}

func (a *SetTagAttr) hasDeleteValues() bool {
	return len(a.DeleteValues) > 0
}

func (a *SetTagAttr) hasTypeUniquePerObject() bool {
	return !TagTypeNotUniquePerObject[a.GetType()]
}

// SetTagRequest represents set-tag input data.
//type SetTagRequest struct {
//ObjUUID string `json:"obj_uuid"`
//ObjType string `json:"obj_type"`
//Tags    map[string]TagAttr
//}

func (t *SetTagRequest) validate() error {
	if t.ObjUUID == "" || t.ObjType == "" {
		return errutil.ErrorBadRequestf(
			"both obj_uuid and obj_type should be specified but got uuid: '%s' and type: '%s",
			t.ObjUUID, t.ObjType,
		)
	}
	for _, tagAttr := range t.Tags {
		if err := tagAttr.validate(t.ObjType); err != nil {
			return err
		}
	}
	return nil
}

func (a *SetTagAttr) validate(objType string) error {
	tagType := strings.ToLower(a.Type)

	// address-group object can only be associated with label
	if objType == "address_group" && !TagTypeAuthorizedOnAddressGroup[tagType] {
		return errutil.ErrorBadRequestf(
			"invalid tag type %v for object type %v", tagType, objType,
		)
	}
	if isTagTypeUniquePerObject(tagType) {
		if len(a.AddValues) > 0 || len(a.DeleteValues) > 0 {
			return errutil.ErrorBadRequestf(
				"tag type %v cannot be set multiple times on a same object", tagType,
			)
		}
		if a.Value == nil && !a.isDeleteRequest() {
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
	for key, val := range rawJSON {
		tagAttr := SetTagAttr{
			Type: strings.ToLower(tagAttr.GetType()),
		}
		if err := json.Unmarshal(val, &tagAttr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid '%v' format: %v", key, err))
		}
		t.Tags = append(t.Tags, &tagAttr)
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
