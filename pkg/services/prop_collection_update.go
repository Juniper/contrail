package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/types"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
)

type restPropCollectionUpdateRequest struct {
	PropCollectionUpdateRequest
	Updates []struct {
		Field     string          `json:"field"`
		Operation string          `json:"operation"`
		Position  *string         `json:"position"`
		Value     json.RawMessage `json:"value"`
	} `json:"updates"`
}

func (p *restPropCollectionUpdateRequest) validate() error {
	if p.UUID == "" {
		return errutil.ErrorBadRequest("prop-collection-update needs object UUID")
	}
	return nil
}

func (p *restPropCollectionUpdateRequest) toPropCollectionUpdateRequest(
	obj interface{},
) (PropCollectionUpdateRequest, error) {
	for _, u := range p.Updates {
		c := PropCollectionChange{
			Field:     u.Field,
			Operation: u.Operation,
		}

		if pos := u.Position; pos != nil {
			if i, err := strconv.ParseInt(*pos, 10, 64); err == nil {
				c.Position = &PropCollectionChange_PositionInt{
					PositionInt: int32(i),
				}
			} else {
				c.Position = &PropCollectionChange_PositionString{
					PositionString: *pos,
				}
			}
		}

		if len(u.Value) > 0 {
			item, err := newCollectionItem(obj, u.Field)
			if err != nil {
				return PropCollectionUpdateRequest{}, err
			}

			err = json.Unmarshal(u.Value, item)
			if err != nil {
				return PropCollectionUpdateRequest{}, err
			}
			c.SetValue(item)
		}

		p.PropCollectionUpdateRequest.Updates = append(p.PropCollectionUpdateRequest.Updates, &c)
	}
	return p.PropCollectionUpdateRequest, nil
}

// RESTPropCollectionUpdate handles a prop-collection-update request.
func (service *ContrailService) RESTPropCollectionUpdate(c echo.Context) error {
	var data restPropCollectionUpdateRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if err := data.validate(); err != nil {
		return errutil.ToHTTPError(err)
	}

	if err := service.InTransactionDoer.DoInTransaction(c.Request().Context(), func(ctx context.Context) error {
		obj, objType, err := service.getObjectAndType(ctx, data.UUID)
		if err != nil {
			return err
		}

		p, err := data.toPropCollectionUpdateRequest(obj)
		if err != nil {
			return errutil.ErrorBadRequestf("error resolving request: %v", err)
		}

		return service.updatePropCollection(ctx, &p, obj, objType)
	}); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.NoContent(http.StatusOK)
}

// PropCollectionUpdate handles a prop-collection-update grpc request.
func (service *ContrailService) PropCollectionUpdate(
	ctx context.Context, request *PropCollectionUpdateRequest,
) (*types.Empty, error) {
	err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		obj, objType, err := service.getObjectAndType(ctx, request.UUID)
		if err != nil {
			return err
		}

		return service.updatePropCollection(ctx, request, obj, objType)
	})
	return &types.Empty{}, err
}

func (service *ContrailService) getObjectAndType(
	ctx context.Context,
	uuid string,
) (basemodels.Object, string, error) {
	idResp, err := service.IDToFQName(ctx, &IDToFQNameRequest{UUID: uuid})
	if err != nil {
		return nil, "", errors.Wrapf(err, "error getting type for provided UUID: %v", uuid)
	}

	o, err := GetObject(ctx, service.Next(), idResp.Type, uuid)
	if err != nil {
		return nil, "", errors.Wrapf(err, "error getting %v with UUID = %v", idResp.Type, uuid)
	}
	return o, idResp.Type, nil
}

func (service *ContrailService) updatePropCollection(
	ctx context.Context,
	request *PropCollectionUpdateRequest,
	obj basemodels.Object,
	objType string,
) error {
	updateMap, err := createUpdateMap(obj, request.Updates)
	if err != nil {
		return errutil.ErrorBadRequest(err.Error())
	}

	e, err := NewEvent(EventOption{
		Data:      updateMap,
		Kind:      objType,
		UUID:      request.UUID,
		Operation: OperationUpdate,
	})
	if err != nil {
		return err
	}

	_, err = e.Process(ctx, service)
	return err
}

func createUpdateMap(
	object basemodels.Object, updates []*PropCollectionChange,
) (map[string]interface{}, error) {
	updateMap := map[string]interface{}{}
	for _, update := range updates {
		updated, err := object.ApplyPropCollectionUpdate(&basemodels.PropCollectionUpdate{
			Field:     update.Field,
			Operation: update.Operation,
			Value:     update.ValueAsInterface(),
			Position:  getPosition(update.Position),
		})
		if err != nil {
			return nil, err
		}
		for key, value := range updated {
			updateMap[key] = value
		}
	}
	return updateMap, nil
}

func getPosition(pos isPropCollectionChange_Position) interface{} {
	switch p := pos.(type) {
	case *PropCollectionChange_PositionInt:
		return p.PositionInt
	case *PropCollectionChange_PositionString:
		return p.PositionString
	default:
		return nil
	}
}

func newCollectionItem(obj interface{}, field string) (interface{}, error) {
	objType := reflect.TypeOf(obj)
	objType = indirect(objType)
	if objType.Kind() != reflect.Struct {
		return nil, errutil.ErrorBadRequest("obj must be a struct")
	}

	wrapperField, ok := fieldByTag(objType, "json", field)
	if !ok {
		return nil, errutil.ErrorBadRequestf("obj has no field with json tag: %s", field)
	}

	wrapperType := indirect(wrapperField.Type)

	if wrapperType.Kind() != reflect.Struct {
		return nil, errutil.ErrorBadRequestf("field '%s' must be a struct or struct pointer type", field)
	}

	innerField := wrapperType.Field(0)
	if !isPropCollectionType(innerField.Type) {
		return nil, errutil.ErrorBadRequestf("provided field '%s' is not valid collection type", field)
	}

	itemType := innerField.Type.Elem()

	return reflect.New(indirect(itemType)).Interface(), nil
}

func fieldByTag(t reflect.Type, key, value string) (reflect.StructField, bool) {
	if t == nil {
		return reflect.StructField{}, false
	}
	t = indirect(t)
	if t.Kind() != reflect.Struct {
		return reflect.StructField{}, false
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag, ok := field.Tag.Lookup(key); ok {
			tagName := strings.SplitN(tag, ",", 2)[0]
			if tagName == value {
				return field, true
			}
		}
	}
	return reflect.StructField{}, false
}

func indirect(t reflect.Type) reflect.Type {
	for t != nil && t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func isPropCollectionType(t reflect.Type) bool {
	if t == nil {
		return false
	}
	k := t.Kind()
	return k == reflect.Map || k == reflect.Slice
}
