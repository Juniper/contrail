package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	strings "strings"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

//API Path definitions.
const (
	SyncPath                 = "sync"
	RefUpdatePath            = "ref-update"
	PropCollectionUpdatePath = "prop-collection-update"
	RefRelaxForDeletePath    = "ref-relax-for-delete"
	SetTagPath               = "set-tag"
	ChownPath                = "chown"
	IntPoolPath              = "int-pool/:pool-name"
)

// Chain setup chain of services.
func Chain(services ...Service) {
	if len(services) < 2 {
		return
	}
	previous := services[0]
	for _, current := range services[1:] {
		previous.SetNext(current)
		previous = current
	}
}

// BaseService is a service that is a link in service chain and has implemented
// all Service methods as noops. Can be embedded in struct to create new service.
type BaseService struct {
	next Service
}

// Next gets next service to call in service chain.
func (service *BaseService) Next() Service {
	return service.next
}

// SetNext sets next service in service chain.
func (service *BaseService) SetNext(next Service) {
	service.next = next
}

// InTransactionDoer executes do function atomically.
type InTransactionDoer interface {
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
}

// NoTransaction executes do function non-atomically.
type NoTransaction struct{}

// DoInTransaction just runs do.
func (*NoTransaction) DoInTransaction(ctx context.Context, do func(context.Context) error) error {
	return do(ctx)
}

// IntPoolAllocator (de)allocates integers in an integer pool.
type IntPoolAllocator interface {
	AllocateInt(context.Context, string) (int64, error)
	DeallocateInt(context.Context, string, int64) error
	SetInt(context.Context, string, int64) error
}

// RefUpdateToUpdateService is a service that promotes CreateRef and DeleteRef
// methods to Update method by fetching the object and updating reference
// field with fieldmask applied.
type RefUpdateToUpdateService struct {
	BaseService

	ReadService       ReadService
	InTransactionDoer InTransactionDoer
}

//EventProcessor can handle events on generic way.
type EventProcessor interface {
	Process(ctx context.Context, event *Event) (*Event, error)
}

//EventProducerService can dispatch method call for event processor.
type EventProducerService struct {
	BaseService
	Processor EventProcessor
	Timeout   time.Duration
}

//ServiceEventProcessor dispatch event to method call.
type ServiceEventProcessor struct {
	Service Service
}

//Process processes event.
func (p *ServiceEventProcessor) Process(ctx context.Context, event *Event) (*Event, error) {
	return event.Process(ctx, p.Service)
}

// ContrailService implementation.
type ContrailService struct {
	BaseService

	MetadataGetter    baseservices.MetadataGetter
	TypeValidator     *models.TypeValidator
	InTransactionDoer InTransactionDoer
	IntPoolAllocator  IntPoolAllocator
}

// RESTSync handles Sync API request.
func (service *ContrailService) RESTSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	// TODO: Call events.Sort()

	responses, err := events.Process(c.Request().Context(), service)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, responses.Events)
}

// RefOperation is enum type for ref-update operation.
type RefOperation string

// RefOperation values.
const (
	RefOperationAdd    RefOperation = "ADD"
	RefOperationDelete RefOperation = "DELETE"
)

// RefUpdate represents ref-update input data.
type RefUpdate struct {
	Operation RefOperation    `json:"operation"`
	Type      string          `json:"type"`
	UUID      string          `json:"uuid"`
	RefType   string          `json:"ref-type"`
	RefUUID   string          `json:"ref-uuid"`
	RefFQName []string        `json:"ref-fq-name"`
	Attr      json.RawMessage `json:"attr"`
}

func (r *RefUpdate) validate() error {
	if r.UUID == "" || r.Type == "" || r.RefType == "" || r.Operation == "" {
		return errutil.ErrorBadRequestf(
			"uuid/type/ref-type/operation is null: %s, %s, %s, %s",
			r.UUID, r.Type, r.RefType, r.Operation,
		)
	}

	if r.Operation != RefOperationAdd && r.Operation != RefOperationDelete {
		return errutil.ErrorBadRequestf("operation should be ADD or DELETE, was %s", r.Operation)
	}

	return nil
}

// RESTRefUpdate handles a ref-update request.
func (service *ContrailService) RESTRefUpdate(c echo.Context) error {
	var data RefUpdate
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if err := data.validate(); err != nil {
		return errutil.ToHTTPError(err)
	}

	ctx := c.Request().Context()
	if data.RefUUID == "" {
		m, err := service.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{Type: data.RefType, FQName: data.RefFQName})
		if err != nil {
			return errutil.ToHTTPError(errutil.ErrorBadRequestf("error resolving ref-uuid using ref-fq-name: %v", err))
		}
		data.RefUUID = m.UUID
	}

	e, err := NewEventFromRefUpdate(&data)
	if err != nil {
		return errutil.ToHTTPError(errutil.ErrorBadRequest(err.Error()))
	}
	if _, err = e.Process(ctx, service); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"uuid": data.UUID})
}

// RefRelax represents ref-relax-for-delete input data.
type RefRelax struct {
	UUID    string `json:"uuid"`
	RefUUID string `json:"ref-uuid"`
}

func (r *RefRelax) validate() error {
	if r.UUID == "" || r.RefUUID == "" {
		return errutil.ErrorBadRequestf(
			"bad request: both uuid and ref-uuid should be specified: %s, %s", r.UUID, r.RefUUID)
	}

	return nil
}

// RESTRefRelaxForDelete handles a ref-relax-for-delete request.
func (service *ContrailService) RESTRefRelaxForDelete(c echo.Context) error {
	var data RefRelax
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if err := data.validate(); err != nil {
		return errutil.ToHTTPError(err)
	}

	// TODO (Kamil): implement ref-relax logic

	return c.JSON(http.StatusOK, map[string]interface{}{"uuid": data.UUID})
}

// PropCollectionUpdateRequest is input request for /prop-collection-update endpoint.
type PropCollectionUpdateRequest struct {
	UUID    string                            `json:"uuid"`
	Updates []basemodels.PropCollectionUpdate `json:"updates"`
}

func (p *PropCollectionUpdateRequest) validate() error {
	if p.UUID == "" {
		return errutil.ErrorBadRequest("prop-collection-update needs object UUID")
	}
	return nil
}

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

		updateMap, err := createUpdateMap(o, data.Updates)
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
	object basemodels.Object, updates []basemodels.PropCollectionUpdate,
) (map[string]interface{}, error) {
	updateMap := map[string]interface{}{}
	for _, update := range updates {
		updated, err := object.ApplyPropCollectionUpdate(&update)
		if err != nil {
			return nil, err
		}
		for key, value := range updated {
			updateMap[key] = value
		}
	}
	return updateMap, nil
}

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
type TagAttr struct {
	IsGlobal     bool        `json:"is_global"`
	Value        interface{} `json:"value"` // Value could be nil, string or slice of strings
	AddValues    []string    `json:"add_values"`
	DeleteValues []string    `json:"delete_values"`
}

func (t *TagAttr) isDeleteRequest() bool {
	return t.Value == nil && (len(t.AddValues) == 0 && len(t.DeleteValues) == 0)
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
	return nil
}

func (t *SetTagRequest) parseObjFields(rawJSON map[string]json.RawMessage) error {
	fields := map[string]*string{
		"obj_uuid": &t.ObjUUID,
		"obj_type": &t.ObjType,
	}

	for key, field := range fields {
		if _, ok := rawJSON[key]; ok {
			if err := json.Unmarshal(rawJSON[key], field); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid '%s' format: %v", key, err))
			}
		}
	}

	return nil
}

func (t *SetTagRequest) parseTagAttrs(rawJSON map[string]json.RawMessage) error {
	t.Tags = make(map[string]TagAttr)
	for key, val := range rawJSON {
		if key == "obj_uuid" || key == "obj_type" {
			continue
		}
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
	setTag.parseObjFields(rawJSON)
	setTag.parseTagAttrs(rawJSON)

	if err := setTag.validate(); err != nil {
		return errutil.ToHTTPError(err)
	}

	ctx := c.Request().Context()
	err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		return service.SetTag(ctx, setTag)
	})
	fmt.Println("error", err)

	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// TagLocator is an object that can identify a tag.
type TagLocator interface {
	GetUUID() string
	GetFQName() []string
	GetPerms2() *models.PermType2
	GetParentType() string
	Kind() string
}

// SetTag allows setting tags based on SetTagRequest.
func (service *ContrailService) SetTag(ctx context.Context, setTag SetTagRequest) error {
	obj, err := GetObject(ctx, service.Next(), setTag.ObjType, setTag.ObjUUID)
	if err != nil {
		return errutil.ErrorBadRequestf(
			"error: %v, while getting %v with UUID %v", err, setTag.ObjType, setTag.ObjUUID,
		)
	}

	refsPerType := models.GroupTagRefsByType(obj.GetTagReferences())
	for tagType, tagAttrs := range setTag.Tags {
		tagType = strings.ToLower(tagType)
		tagName := models.CreateTagName(tagType, tagAttrs.Value.(string))

		// address-group object can only be associated with label
		if setTag.ObjType == "address_group" && !TagTypeAuthorizedOnAddressGroup[tagType] {
			return errutil.ErrorBadRequestf(
				"invalid tag type %v for object type %v", tagType, setTag.ObjType,
			)
		}

		if tagAttrs.isDeleteRequest() {
			// Remove tag from obj.TagRefs if tagAttrs doesn't exist
			tagRef := obj.GetTagReferences().Find(func(ref basemodels.Reference) bool {
				tType, _ := models.TagTypeValueFromFQName(ref.GetTo())
				return tType == tagType
			})

			if tagRef != nil {
				event, err := setTag.tagRefEvent(tagRef.GetUUID(), RefOperationDelete)
				if err != nil {
					return errutil.ErrorBadRequest(err.Error())
				}
				event.Process(ctx, service)
				//events = append(events, event)
			}
			refsPerType[tagType] = nil
			continue
		}

		refsPerValue := models.GroupTagRefsByValue(refsPerType[tagType])
		if !TagTypeNotUniquePerObject[tagType] {
			if len(tagAttrs.AddValues) > 0 || len(tagAttrs.DeleteValues) > 0 {
				return errutil.ErrorBadRequestf(
					"tag type %v cannot be set multiple times on a same object", tagType,
				)
			}

			if _, ok := tagAttrs.Value.(string); !ok {
				return errutil.ErrorBadRequestf("no valid value provided for tag type %v", tagType)
			}

			if _, ok := refsPerValue[tagAttrs.Value.(string)]; ok {
				// don't need to update if tag type with same value already exists
				continue
			}

			for _, ref := range refsPerValue {
				rFqName := strings.Join(ref.GetTo(), "-")
				for _, tagRef := range obj.GetTagReferences() {
					tagRefFqName := strings.Join(tagRef.GetTo(), "-")
					if rFqName == tagRefFqName {
						// object already have a reference to this tag type with a different value, remove it
						tagMeta := &basemodels.Metadata{
							UUID: tagRef.GetUUID(),
							Type: models.KindTag,
						}
						event, err := setTag.tagRefEvent(tagMeta.UUID, RefOperationDelete)
						if err != nil {
							return errutil.ErrorBadRequest(err.Error())
						}
						event.Process(ctx, service)
						break
					}
				}
			}
			// finally, reference the tag type with the new value
			tagMeta, err := service.tagByFQName(ctx, obj, tagName, tagAttrs.IsGlobal)
			if err != nil {
				return err
			}
			event, err := setTag.tagRefEvent(tagMeta.UUID, RefOperationAdd)
			if err != nil {
				return errutil.ErrorBadRequest(err.Error())
			}
			event.Process(ctx, service)
		} else {
			for _, tagValue := range tagAttrs.AddValues {
				if _, ok := refsPerValue[tagValue]; ok {
					continue // already done
				}

				tagName := fmt.Sprintf("%v=%v", tagType, tagValue)
				tagMeta, err := service.tagByFQName(ctx, obj, tagName, tagAttrs.IsGlobal)
				if err != nil {
					return err
				}
				event, err := setTag.tagRefEvent(tagMeta.UUID, RefOperationAdd)
				if err != nil {
					return errutil.ErrorBadRequest(err.Error())
				}
				event.Process(ctx, service)
			}
			for _, tagValue := range tagAttrs.DeleteValues {
				tagName := models.CreateTagName(tagType, tagValue)
				tagMeta, err := service.tagByFQName(ctx, obj, tagName, tagAttrs.IsGlobal)
				if err != nil {
					return err
				}
				event, err := setTag.tagRefEvent(tagMeta.UUID, RefOperationDelete)
				if err != nil {
					return errutil.ErrorBadRequest(err.Error())
				}
				event.Process(ctx, service)
			}
		}
	}
	return nil
}

func (service *ContrailService) tagByFQName(
	ctx context.Context,
	obj basemodels.Object,
	tagName string,
	isGlobal bool,
) (*basemodels.Metadata, error) {
	var fqName []string

	if isGlobal {
		fqName = []string{tagName}
	} else {
		o := obj.(TagLocator)
		fqName = o.GetFQName()
		if obj.Kind() == "project" {
			fqName = append(fqName, tagName)
		} else if o.GetParentType() == "project" {
			fqName[len(fqName)-1] = tagName
		} else if perms2 := o.GetPerms2(); perms2 != nil {
			data, err := service.MetadataGetter.GetMetadata(
				ctx, basemodels.Metadata{UUID: perms2.Owner},
			)
			if err != nil {
				return nil, errutil.ErrorNotFoundf("cannot find %s %s owner: %v", tagName, o.GetUUID(), err)
			}
			fqName = data.FQName
		} else {
			return nil, errutil.ErrorNotFoundf("Not able to determine the scope of the tag '%s'", tagName)
		}
	}

	data, err := service.MetadataGetter.GetMetadata(
		ctx, basemodels.Metadata{FQName: fqName, Type: models.KindTag},
	)
	if err != nil {
		return nil, errutil.ErrorNotFoundf("not able to determine the scope of the tag %s: %v", tagName, err)
	}
	return data, nil
}

// Chown handles chown request.
func (service *ContrailService) Chown(ctx context.Context, request *ChownRequest) (*Empty, error) {
	// TODO: implement chown logic.
	return &Empty{}, nil
}

// RESTChown handles chown request.
func (service *ContrailService) RESTChown(c echo.Context) error {
	// TODO: bind request
	ctx := c.Request().Context()
	if _, err := service.Chown(ctx, &ChownRequest{}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// RESTIntPoolAllocate handles a POST on int-pool/:pool-name/:value and int-pool/:pool-name request.
// TODO(Michal): gRPC endpoint
func (service *ContrailService) RESTIntPoolAllocate(c echo.Context) error {
	ctx := c.Request().Context()
	auth := auth.GetAuthCTX(ctx)
	if !auth.IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}
	pool := c.Param("pool-name")
	v := c.Param("value")
	var val int64
	if v == "" {
		resp, err := service.AllocateInt(ctx, &AllocateIntRequest{Pool: pool})
		if err != nil {
			return errutil.ToHTTPError(err)
		}
		val = resp.Value
	} else {
		var err error
		val, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return errutil.ToHTTPError(errutil.ErrorBadRequestf("Invalid int to allocate (%v): %s", v, err))
		}
		if _, err = service.SetInt(ctx, &SetIntRequest{Pool: pool, Value: val}); err != nil {
			return errutil.ToHTTPError(err)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"value": val})
}

// AllocateInt allocates int in given int-pool.
func (service *ContrailService) AllocateInt(
	ctx context.Context, request *AllocateIntRequest,
) (*AllocateIntResponse, error) {
	var v int64
	if err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		var err error
		if v, err = service.IntPoolAllocator.AllocateInt(ctx, request.GetPool()); err != nil {
			return errutil.ErrorBadRequestf("Failed to allocate next int: %s", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &AllocateIntResponse{Value: v}, nil
}

// SetInt sets int in given int-pool.
func (service *ContrailService) SetInt(ctx context.Context, request *SetIntRequest) (*Empty, error) {
	if err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		if err := service.IntPoolAllocator.SetInt(ctx, request.GetPool(), request.GetValue()); err != nil {
			return errutil.ErrorBadRequestf("Failed to allocate specified int: %s", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &Empty{}, nil
}

// RESTIntPoolDeallocate handles a DELETE on int-pool/:pool-name/:value request.
// TODO(Michal): gRPC endpoint
func (service *ContrailService) RESTIntPoolDeallocate(c echo.Context) error {
	ctx := c.Request().Context()
	auth := auth.GetAuthCTX(ctx)
	if !auth.IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}
	pool := c.Param("pool-name")
	v := c.Param("value")
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return errutil.ToHTTPError(errutil.ErrorBadRequestf("Invalid int to deallocate (%v): %s", v, err))
	}

	if _, err := service.DeallocateInt(ctx, &DeallocateIntRequest{Pool: pool, Value: i}); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.NoContent(http.StatusOK)
}

// DeallocateInt deallocates int in given int-pool.
func (service *ContrailService) DeallocateInt(ctx context.Context, request *DeallocateIntRequest) (*Empty, error) {
	if err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		if err := service.IntPoolAllocator.DeallocateInt(ctx, request.GetPool(), request.GetValue()); err != nil {
			return errutil.ErrorBadRequest(err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &Empty{}, nil
}

type routeRegistry interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
