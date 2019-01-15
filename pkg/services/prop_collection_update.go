package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models/basemodels"
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

func (p *restPropCollectionUpdateRequest) toPropCollectionUpdateRequest(obj interface{}) (PropCollectionUpdateRequest, error) {
	for _, u := range p.Updates {
		c := PropCollectionChange{
			Field:     u.Field,
			Operation: u.Operation,
		}
		if u.Position != nil {
			c.Position = &PropCollectionChange_PositionString{
				PositionString: *u.Position,
			}

			if i, err := strconv.ParseInt(*u.Position, 10, 64); err == nil {
				c.Position = &PropCollectionChange_PositionInt{
					PositionInt: i,
				}
			}
		}

		item, err := newCollectionItem(obj, u.Field)
		if err != nil {
			return PropCollectionUpdateRequest{}, err
		}

		err = json.Unmarshal(u.Value, item)
		if err != nil {
			return PropCollectionUpdateRequest{}, err
		}

		c.SetValue(item)
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

	if err := service.updatePropCollection(c.Request().Context(), &data); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (service *ContrailService) updatePropCollection(
	ctx context.Context,
	data *restPropCollectionUpdateRequest,
) error {
	err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		m, err := service.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{UUID: data.UUID})
		if err != nil {
			return errors.Wrap(err, "error getting metadata for provided UUID: %v")
		}

		o, err := GetObject(ctx, service.Next(), m.Type, data.UUID)
		if err != nil {
			return errors.Wrapf(err, "error getting %v with UUID = %v", m.Type, data.UUID)
		}

		p, err := data.toPropCollectionUpdateRequest(o)
		if err != nil {
			return err
		}

		updateMap, err := createUpdateMap(o, p.Updates)
		if err != nil {
			return errutil.ErrorBadRequest(err.Error())
		}

		e, err := NewEvent(&EventOption{
			Data:      updateMap,
			Kind:      m.Type,
			UUID:      data.UUID,
			Operation: OperationUpdate,
		})
		if err != nil {
			return err
		}

		_, err = e.Process(ctx, service)
		return err
	})
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return nil
}

func createUpdateMap(
	object basemodels.Object, updates []*PropCollectionChange,
) (map[string]interface{}, error) {
	updateMap := map[string]interface{}{}
	for _, update := range updates {
		_ = update
		//updated, err := object.ApplyPropCollectionUpdate(update)
		//if err != nil {
		//return nil, err
		//}
		//for key, value := range updated {
		//updateMap[key] = value
		//}
	}
	return updateMap, nil
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
	if !isCollectionType(innerField.Type) {
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

func isCollectionType(t reflect.Type) bool {
	if t == nil {
		return false
	}
	k := t.Kind()
	return k == reflect.Map || k == reflect.Slice

}
