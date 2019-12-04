package services

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/gogo/protobuf/types"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	// TODO(dfurman): Decouple from below packages
	//"github.com/Juniper/asf/pkg/auth"
	//"github.com/Juniper/asf/pkg/collector"

	uuid "github.com/satori/go.uuid"
)

//API Path definitions.
const (
	SyncPath                 = "sync"
	RefUpdatePath            = "ref-update"
	PropCollectionUpdatePath = "prop-collection-update"
	RefRelaxForDeletePath    = "ref-relax-for-delete"
	SetTagPath               = "set-tag"
	ChownPath                = "chown"
	IntPoolPath              = "int-pool"
	IntPoolsPath             = "int-pools"
	ObjPerms                 = "obj-perms"
)

// Reference is a generic reference that can be retrieved from ref update event.
type Reference = basemodels.Reference

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

// InternalContextWriteServiceWrapper is a WriteService that marks it requests
// with internal request context.
type InternalContextWriteServiceWrapper struct {
	WriteService
}

// InTransactionDoer executes do function atomically.
type InTransactionDoer interface {
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
}

// NoTransaction executes do function non-atomically.
var NoTransaction = noTransaction{}

type noTransaction struct{}

// DoInTransaction just runs do.
func (noTransaction) DoInTransaction(ctx context.Context, do func(context.Context) error) error {
	return do(ctx)
}

// IntPoolAllocator (de)allocates integers in an integer pool.
type IntPoolAllocator interface {
	CreateIntPool(context.Context, string, int64, int64) error
	GetIntOwner(context.Context, string, int64) (string, error)
	DeleteIntPool(context.Context, string) error
	AllocateInt(context.Context, string, string) (int64, error)
	SetInt(context.Context, string, int64, string) error
	DeallocateInt(context.Context, string, int64) error
}

// RefRelaxer makes references not prevent the referenced resource from being deleted.
type RefRelaxer interface {
	RelaxRef(ctx context.Context, request *RelaxRefRequest) error
}

// UserAgentKVService is a service which manages operations on key-value store
type UserAgentKVService interface {
	StoreKeyValue(ctx context.Context, key string, value string) error
	RetrieveValues(ctx context.Context, keys []string) (vals []string, err error)
	DeleteKey(ctx context.Context, key string) error
	RetrieveKVPs(ctx context.Context) (kvps []*models.KeyValuePair, err error)
}

// RefUpdateToUpdateService is a service that promotes CreateRef and DeleteRef
// methods to Update method by fetching the object and updating reference
// field with fieldmask applied.
type RefUpdateToUpdateService struct {
	BaseService

	ReadService       ReadService
	InTransactionDoer InTransactionDoer
}

// RBACService will RBAC check on resource opeations.
type RBACService struct {
	BaseService
	ReadService ReadService
	AAAMode     string
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

// EventListProcessor processes event lists in transaction.
type EventListProcessor struct {
	EventProcessor
	InTransactionDoer InTransactionDoer
}

// ProcessList processes list of events.
func (p *EventListProcessor) ProcessList(ctx context.Context, e *EventList) (*EventList, error) {
	var results []*Event
	if err := p.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		for _, event := range e.Events {
			r, err := p.Process(ctx, event)
			if err != nil {
				return err
			}
			results = append(results, r)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &EventList{
		Events: results,
	}, nil
}

// ContrailService implementation.
type ContrailService struct {
	BaseService

	DBService          ReadService
	MetadataGetter     baseservices.MetadataGetter
	TypeValidator      *models.TypeValidator
	InTransactionDoer  InTransactionDoer
	IntPoolAllocator   IntPoolAllocator
	RefRelaxer         RefRelaxer
	UserAgentKVService UserAgentKVService
	Collector          collector.Collector
}

// RefUpdate represents ref-update input data.
type RefUpdate struct {
	Operation RefOperation           `json:"operation"`
	Type      string                 `json:"type"`
	UUID      string                 `json:"uuid"`
	RefType   string                 `json:"ref-type"`
	RefUUID   string                 `json:"ref-uuid"`
	RefFQName []string               `json:"ref-fq-name"`
	Attr      map[string]interface{} `json:"attr,omitempty"`
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

	e, err := NewRefUpdateEvent(RefUpdateOption{
		ReferenceType: basemodels.ReferenceKind(data.Type, data.RefType),
		FromUUID:      data.UUID,
		ToUUID:        data.RefUUID,
		Operation:     data.Operation,
		Attr:          data.Attr,
	})
	if err != nil {
		return errutil.ToHTTPError(errutil.ErrorBadRequest(err.Error()))
	}
	if _, err = e.Process(ctx, service); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"uuid": data.UUID})
}

// RESTRefRelaxForDelete handles a ref-relax-for-delete request.
func (service *ContrailService) RESTRefRelaxForDelete(c echo.Context) error {
	var data RelaxRefRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if err := validateRelaxRefRequest(&data); err != nil {
		return errutil.ToHTTPError(err)
	}

	response, err := service.RelaxRef(c.Request().Context(), &data)
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, response)
}

func validateRelaxRefRequest(r *RelaxRefRequest) error {
	if r.UUID == "" || r.RefUUID == "" {
		return errutil.ErrorBadRequestf(
			"bad request: both uuid and ref-uuid should be specified: %s, %s", r.UUID, r.RefUUID)
	}

	return nil
}

// RelaxRef makes a reference not prevent the referenced resource from being deleted.
func (service *ContrailService) RelaxRef(ctx context.Context, request *RelaxRefRequest) (*RelaxRefResponse, error) {
	err := service.RefRelaxer.RelaxRef(ctx, request)
	if err != nil {
		return nil, err
	}
	return &RelaxRefResponse{UUID: request.UUID}, nil
}

// RESTChown handles chown request.
func (service *ContrailService) RESTChown(c echo.Context) error {
	var data ChownRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	ctx := c.Request().Context()
	if _, err := service.Chown(ctx, &data); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// Chown handles chown request.
func (service *ContrailService) Chown(ctx context.Context, request *ChownRequest) (*types.Empty, error) {
	if err := validateChownRequest(request); err != nil {
		return nil, err
	}

	if err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		metadata, err := service.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{UUID: request.GetUUID()})
		if err != nil {
			return errors.Wrapf(err, "failed to change the owner of the resource with UUID '%v'", request.GetUUID())
		}

		// nolint: lll
		// TODO: check permissions, see https://github.com/Juniper/asf-controller/blob/137e2a08025e1ae7084621c0f081f7b99d1b04cd/src/config/api-server/vnc_cfg_api_server/vnc_cfg_api_server.py#L2409

		var fm types.FieldMask
		basemodels.FieldMaskAppend(&fm, basemodels.CommonFieldPerms2, models.PermType2FieldOwner)

		event, err := NewEvent(EventOption{
			UUID:      request.GetUUID(),
			Kind:      metadata.Type,
			Operation: OperationUpdate,
			Data: map[string]interface{}{
				"perms2": map[string]interface{}{
					"owner": request.GetOwner(),
				},
			},
			FieldMask: &fm,
		})
		if err != nil {
			return errors.Wrapf(err, "failed to change the owner of '%v' with UUID '%v'", metadata.Type, request.GetUUID())
		}

		_, err = event.Process(ctx, service)
		return errors.Wrapf(err, "failed to change the owner of '%v' with UUID '%v'", metadata.Type, request.GetUUID())
	}); err != nil {
		return nil, err
	}

	return &types.Empty{}, nil
}

func validateChownRequest(r *ChownRequest) error {
	if r == nil || r.UUID == "" || r.Owner == "" {
		return errutil.ErrorBadRequestf(
			"bad request: both uuid and owner should be specified: %s, %s", r.GetUUID(), r.GetOwner())
	}

	if _, err := uuid.FromString(r.GetUUID()); err != nil {
		return errutil.ErrorBadRequestf(
			"bad request: invalid uuid format (not UUID): %s", r.GetUUID())
	}
	if _, err := uuid.FromString(r.GetOwner()); err != nil {
		return errutil.ErrorBadRequestf(
			"bad request: invalid owner format (not UUID): %s", r.GetOwner())
	}

	return nil
}

// RESTCreateIntPool handles a POST on int-pools requests
func (service *ContrailService) RESTCreateIntPool(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetAuthCTX(ctx).IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}

	data := &CreateIntPoolRequest{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if _, err := service.CreateIntPool(ctx, data); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// RESTGetIntOwner handles a GET on int-owner requests
func (service *ContrailService) RESTGetIntOwner(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetAuthCTX(ctx).IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}
	aValue := c.QueryParam("value")
	if aValue == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request: missing value for getting int owner")
	}
	value, err := strconv.Atoi(aValue)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid request: invalid value (%v) "+
			"for getting int owner: %v", aValue, err))
	}
	if value < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid request: invalid value (%v) "+
			"for getting int owner", value))
	}
	pool := c.QueryParam("pool")
	if pool == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request: missing pool name for getting int owner")
	}
	response, err := service.GetIntOwner(ctx, &GetIntOwnerRequest{Pool: pool, Value: int64(value)})
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, response)
}

// RESTDeleteIntPool handles a POST on int-pools requests
func (service *ContrailService) RESTDeleteIntPool(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetAuthCTX(ctx).IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}

	data := &DeleteIntPoolRequest{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if _, err := service.DeleteIntPool(ctx, data); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// CreateIntPool creates empty int pool
func (service *ContrailService) CreateIntPool(
	ctx context.Context, r *CreateIntPoolRequest,
) (*types.Empty, error) {
	if err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		return service.IntPoolAllocator.CreateIntPool(ctx, r.Pool, r.Start, r.End)
	}); err != nil {
		return nil, err
	}

	return &types.Empty{}, nil
}

// GetIntOwner returns owner of allocated int in given int-pool.
func (service *ContrailService) GetIntOwner(
	ctx context.Context, request *GetIntOwnerRequest,
) (*GetIntOwnerResponse, error) {
	if request.GetPool() == "" {
		return nil, errutil.ErrorBadRequest("Missing pool name for getting int owner")
	}

	var err error
	response := &GetIntOwnerResponse{}
	err = service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		var owner string
		owner, err = service.IntPoolAllocator.GetIntOwner(ctx, request.GetPool(), request.GetValue())
		if err != nil {
			return err
		}
		response.Owner = owner
		return nil
	})

	if err != nil && !errutil.IsNotFound(err) {
		return nil, errutil.ErrorBadRequestf("Failed to fetch int owner: %s", err)
	}
	return response, nil
}

// DeleteIntPool deletes int pool
func (service *ContrailService) DeleteIntPool(
	ctx context.Context, r *DeleteIntPoolRequest,
) (*types.Empty, error) {
	if err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		return service.IntPoolAllocator.DeleteIntPool(ctx, r.Pool)
	}); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}

// IntPoolAllocationBody represents int-pool input data.
type IntPoolAllocationBody struct {
	Pool  string `json:"pool"`
	Value *int64 `json:"value,omitempty"`
	Owner string `json:"owner,omitempty"`
}

// RESTIntPoolAllocate handles a POST request on int-pool.
func (service *ContrailService) RESTIntPoolAllocate(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetAuthCTX(ctx).IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}
	var allocReq IntPoolAllocationBody
	if err := c.Bind(&allocReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}
	var allocatedVal int64
	if allocReq.Value == nil {
		resp, err := service.AllocateInt(ctx, &AllocateIntRequest{Pool: allocReq.Pool, Owner: allocReq.Owner})
		if err != nil {
			return errutil.ToHTTPError(err)
		}
		allocatedVal = resp.Value
	} else {
		if _, err := service.SetInt(
			ctx,
			&SetIntRequest{Pool: allocReq.Pool, Value: *allocReq.Value, Owner: allocReq.Owner},
		); err != nil {
			return errutil.ToHTTPError(err)
		}
		allocatedVal = *allocReq.Value
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"value": allocatedVal})
}

// AllocateInt allocates int in given int-pool.
func (service *ContrailService) AllocateInt(
	ctx context.Context, request *AllocateIntRequest,
) (*AllocateIntResponse, error) {
	var v int64
	if request.GetPool() == "" {
		err := errutil.ErrorBadRequest("Missing pool name for int-pool allocation")
		return nil, err
	}
	if err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		var err error
		if v, err = service.IntPoolAllocator.AllocateInt(ctx, request.GetPool(), request.GetOwner()); err != nil {
			return errutil.ErrorBadRequestf("Failed to allocate next int: %s", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &AllocateIntResponse{Value: v}, nil
}

// SetInt sets int in given int-pool.
func (service *ContrailService) SetInt(ctx context.Context, request *SetIntRequest) (*types.Empty, error) {
	if request.GetPool() == "" {
		err := errutil.ErrorBadRequest("Missing pool name for int-pool allocation")
		return nil, err
	}
	if err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		if err := service.IntPoolAllocator.SetInt(
			ctx, request.GetPool(), request.GetValue(), request.GetOwner(),
		); err != nil {
			return errutil.ErrorBadRequestf("Failed to allocate specified int: %s", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}

// RESTIntPoolDeallocate handles a DELETE request on int-pool.
func (service *ContrailService) RESTIntPoolDeallocate(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetAuthCTX(ctx).IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}
	var allocReq IntPoolAllocationBody
	if err := c.Bind(&allocReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}
	if allocReq.Value == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "missing value for int-pool deallocation")
	}
	if _, err := service.DeallocateInt(
		ctx, &DeallocateIntRequest{Pool: allocReq.Pool, Value: *allocReq.Value}); err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.NoContent(http.StatusOK)
}

// DeallocateInt deallocates int in given int-pool.
func (service *ContrailService) DeallocateInt(
	ctx context.Context, request *DeallocateIntRequest,
) (*types.Empty, error) {
	if request.GetPool() == "" {
		return nil, errutil.ErrorBadRequest("missing pool name for int-pool allocation")
	}
	if err := service.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		if err := service.IntPoolAllocator.DeallocateInt(ctx, request.GetPool(), request.GetValue()); err != nil {
			return errutil.ErrorBadRequest(err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}

type routeRegistry interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// GetRequestSchema returns 'https://' for TLS based request or 'http://' otherwise
func GetRequestSchema(r *http.Request) string {
	if r.TLS != nil {
		return "https://"
	}
	return "http://"
}

var requestIDKey interface{} = "requestIDKey"

// WithRequestID assign new request_id to context if there is no one in.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	if ctx.Value(requestIDKey) != nil {
		return ctx
	}

	if requestID == "" {
		requestID = "req-" + uuid.NewV4().String()
	}

	return context.WithValue(ctx, requestIDKey, requestID)
}

// GetRequestID retrieves request id from context.
func GetRequestID(ctx context.Context) string {
	value := ctx.Value(requestIDKey)
	if value == nil {
		return "NO-REQUESTID"
	}

	requestID, ok := value.(string)
	if !ok {
		return "NO-REQUESTID"
	}

	return requestID
}
