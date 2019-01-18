package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gogo/protobuf/types"
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
	IntPoolPath              = "int-pool"
	IntOwnerPath             = "int-owner"
	IntPoolsPath             = "int-pools"
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

// ContrailService implementation.
type ContrailService struct {
	BaseService

	MetadataGetter    baseservices.MetadataGetter
	TypeValidator     *models.TypeValidator
	InTransactionDoer InTransactionDoer
	IntPoolAllocator  IntPoolAllocator
	RefRelaxer        RefRelaxer
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

	e, err := NewEventFromRefUpdate(RefUpdateOption{
		ReferenceType: basemodels.ReferenceKind(data.Type, data.RefType),
		FromUUID:      data.UUID,
		ToUUID:        data.RefUUID,
		Operation:     data.Operation,
		AttrData:      data.Attr,
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

// Chown handles chown request.
func (service *ContrailService) Chown(ctx context.Context, request *ChownRequest) (*types.Empty, error) {
	// TODO: implement chown logic.
	return &types.Empty{}, nil
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

	data := &GetIntOwnerRequest{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	response, err := service.GetIntOwner(ctx, data)
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

	if err != nil && err != errutil.ErrorNotFound {
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
