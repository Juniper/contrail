package services

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// RESTPropCollectionUpdate handles a prop-collection-update request.
func (service *ContrailService) RESTPropCollectionUpdate(c echo.Context) error {
	var data PropCollectionUpdateRequest
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
	data *PropCollectionUpdateRequest,
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
		_ = o
		if err = data.resolveCollectionItems(o); err != nil {
			return errors.Wrap(err, "error resolving collection fields")
		}

		updateMap, err := createUpdateMap(o, data.Updates)
		if err != nil {
			return errutil.ErrorBadRequest(err.Error())
		}
		_ = updateMap

		e, err := NewEvent(&EventOption{
			//Data:      updateMap, TODO
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

func (p *PropCollectionUpdateRequest) validate() error {
	if p.UUID == "" {
		return errutil.ErrorBadRequest("prop-collection-update needs object UUID")
	}
	return nil
}

func (p *PropCollectionUpdateRequest) resolveCollectionItems(obj interface{}) error {
	for _, u := range p.Updates {
		item, err := newCollectionItem(obj, u.Field)
		if err != nil {
			return err
		}
		u.SetValue(item)
	}
	return nil
}

func newCollectionItem(obj interface{}, field string) (interface{}, error) {
	objType := reflect.TypeOf(obj)
	objType = indirectPtr(objType)
	if objType.Kind() != reflect.Struct {
		return nil, errutil.ErrorBadRequest("obj must be a struct")
	}

	wrapperField, ok := fieldByTag(objType, "json", field)
	if !ok {
		return nil, errutil.ErrorBadRequestf("obj has no field with json tag: %s", field)
	}

	wrapperType := indirectPtr(wrapperField.Type)

	if wrapperType.Kind() != reflect.Struct {
		return nil, errutil.ErrorBadRequestf("field '%s' must be a struct or struct pointer type", field)
	}

	innerField := wrapperType.Field(0)
	if !isCollectionType(innerField.Type) {
		return nil, errutil.ErrorBadRequestf("provided field '%s' is not valid collection type", field)
	}

	itemType := innerField.Type.Elem()

	return reflect.Zero(itemType).Interface(), nil
}

func fieldByTag(t reflect.Type, key, value string) (reflect.StructField, bool) {
	if t == nil {
		return reflect.StructField{}, false
	}
	t = indirectPtr(t)
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

func indirectPtr(t reflect.Type) reflect.Type {
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
