package intpool

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/gogo/protobuf/types"
	"github.com/labstack/echo"
)

const (
	IntPoolPath              = "int-pool"
	IntPoolsPath             = "int-pools"
)

type IntPoolPlugin struct {
	Allocator IntPoolAllocator
	InTransactionDoer InTransactionDoer
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

// InTransactionDoer executes do function atomically.
type InTransactionDoer interface {
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
}

func (p *IntPoolPlugin) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.GET(IntPoolPath, p.RESTGetIntOwner)
	r.POST(IntPoolPath, p.RESTIntPoolAllocate)
	r.DELETE(IntPoolPath, p.RESTIntPoolDeallocate)
	r.POST(IntPoolsPath, p.RESTCreateIntPool)
	r.DELETE(IntPoolsPath, p.RESTDeleteIntPool)
}

// RegisterGRPCAPI registers GRPC services.
func (p *IntPoolPlugin) RegisterGRPCAPI(r apiserver.GRPCRouter) {
	r.RegisterService(&_IPAM_serviceDesc, p)
}

// RESTCreateIntPool handles a POST on int-pools requests
func (p *IntPoolPlugin) RESTCreateIntPool(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetIdentity(ctx).IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}

	data := &CreateIntPoolRequest{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if _, err := p.CreateIntPool(ctx, data); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// RESTGetIntOwner handles a GET on int-owner requests
func (p *IntPoolPlugin) RESTGetIntOwner(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetIdentity(ctx).IsAdmin() {
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
	response, err := p.GetIntOwner(ctx, &GetIntOwnerRequest{Pool: pool, Value: int64(value)})
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, response)
}

// RESTDeleteIntPool handles a POST on int-pools requests
func (p *IntPoolPlugin) RESTDeleteIntPool(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetIdentity(ctx).IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}

	data := &DeleteIntPoolRequest{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if _, err := p.DeleteIntPool(ctx, data); err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// CreateIntPool creates empty int pool
func (p *IntPoolPlugin)CreateIntPool(
	ctx context.Context,  r *CreateIntPoolRequest,
) (*types.Empty, error) {
	if err := p.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error{
		return p.Allocator.CreateIntPool(ctx, r.Pool, r.Start, r.End)
	}); err != nil {
		return nil, err
	}

	return &types.Empty{}, nil
}

// GetIntOwner returns owner of allocated int in given int-pool.
func (p *IntPoolPlugin) GetIntOwner(
	ctx context.Context, request *GetIntOwnerRequest,
) (*GetIntOwnerResponse, error) {
	if request.GetPool() == "" {
		return nil, errutil.ErrorBadRequest("Missing pool name for getting int owner")
	}

	var err error
	response := &GetIntOwnerResponse{}
	err = p.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		var owner string
		owner, err = p.Allocator.GetIntOwner(ctx, request.GetPool(), request.GetValue())
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
func (p *IntPoolPlugin) DeleteIntPool(
	ctx context.Context, r *DeleteIntPoolRequest,
) (*types.Empty, error) {
	if err := p.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		return p.Allocator.DeleteIntPool(ctx, r.Pool)
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
func (p *IntPoolPlugin) RESTIntPoolAllocate(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetIdentity(ctx).IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}
	var allocReq IntPoolAllocationBody
	if err := c.Bind(&allocReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}
	var allocatedVal int64
	if allocReq.Value == nil {
		resp, err := p.AllocateInt(ctx, &AllocateIntRequest{Pool: allocReq.Pool, Owner: allocReq.Owner})
		if err != nil {
			return errutil.ToHTTPError(err)
		}
		allocatedVal = resp.Value
	} else {
		if _, err := p.SetInt(
			ctx, &SetIntRequest{Pool: allocReq.Pool, Value: *allocReq.Value, Owner: allocReq.Owner},
		); err != nil {
			return errutil.ToHTTPError(err)
		}
		allocatedVal = *allocReq.Value
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"value": allocatedVal})
}

// AllocateInt allocates int in given int-pool.
func (p *IntPoolPlugin) AllocateInt(
	ctx context.Context, request *AllocateIntRequest,
) (*AllocateIntResponse, error) {
	if request.GetPool() == "" {
		err := errutil.ErrorBadRequest("Missing pool name for int-pool allocation")
		return nil, err
	}
	var v int64
	if err := p.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		var err error
		if v, err = p.Allocator.AllocateInt(ctx, request.GetPool(), request.GetOwner()); err != nil{
			return errutil.ErrorBadRequestf("Failed to allocate next int: %s", err)
		}
		return nil
	});err != nil {
		return nil, err
	}
	return &AllocateIntResponse{Value: v}, nil
}

// SetInt sets int in given int-pool.
func (p *IntPoolPlugin) SetInt(
	ctx context.Context, request *SetIntRequest,
) (*types.Empty, error) {
	if request.GetPool() == "" {
		err := errutil.ErrorBadRequest("Missing pool name for int-pool allocation")
		return nil, err
	}
	if err := p.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		if err := p.Allocator.SetInt(
			ctx, request.GetPool(), request.GetValue(), request.GetOwner(),
		); err != nil{
			return errutil.ErrorBadRequestf("Failed to allocate specified int: %s", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}

// RESTIntPoolDeallocate handles a DELETE request on int-pool.
func (p *IntPoolPlugin) RESTIntPoolDeallocate(c echo.Context) error {
	ctx := c.Request().Context()
	if !auth.GetIdentity(ctx).IsAdmin() {
		return errutil.ToHTTPError(errutil.ErrorPermissionDenied)
	}
	var allocReq IntPoolAllocationBody
	if err := c.Bind(&allocReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}
	if allocReq.Value == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "missing value for int-pool deallocation")
	}
	if _, err := p.DeallocateInt(
		ctx, &DeallocateIntRequest{Pool: allocReq.Pool, Value: *allocReq.Value}); err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.NoContent(http.StatusOK)
}

// DeallocateInt deallocates int in given int-pool.
func (p *IntPoolPlugin) DeallocateInt(
	ctx context.Context, request *DeallocateIntRequest,
) (*types.Empty, error) {
	if request.GetPool() == "" {
		return nil, errutil.ErrorBadRequest("missing pool name for int-pool allocation")
	}
	if err := p.InTransactionDoer.DoInTransaction(ctx, func(ctx context.Context) error {
		if err := p.Allocator.DeallocateInt(ctx, request.GetPool(), request.GetValue()); err != nil {
			return errutil.ErrorBadRequest(err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}

