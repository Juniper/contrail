package plugin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/gogo/protobuf/types"
	"github.com/labstack/echo"
)

const (
	UserAgentKVPath              = "useragent-kv"
	UserAgentKVOperationStore    = "STORE"
	UserAgentKVOperationRetrieve = "RETRIEVE"
	UserAgentKVOperationDelete   = "DELETE"
)

type UserAgentKVPlugin struct {
	UserAgentKVService UserAgentKVService
}

// UserAgentKVService is a service which manages operations on key-value store
type UserAgentKVService interface {
	StoreKeyValue(ctx context.Context, key string, value string) error
	RetrieveValues(ctx context.Context, keys []string) (vals []string, err error)
	DeleteKey(ctx context.Context, key string) error
	RetrieveKVPs(ctx context.Context) (kvps []*KeyValuePair, err error)
}

var _ UserAgentKVServer = &UserAgentKVPlugin{}
var _ apiserver.APIPlugin = &UserAgentKVPlugin{}

// RESTGetObjPerms handles GET operation of obj-perms request.
func (p *UserAgentKVPlugin) RESTGetObjPerms(c echo.Context) error {
	return c.JSON(http.StatusOK, auth.GetIdentity(c.Request().Context()).GetObjPerms())
}

// RegisterHTTPAPI registers HTTP services.
func (p *UserAgentKVPlugin) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.POST(UserAgentKVPath, p.RESTUserAgentKV)
}

// RegisterGRPCAPI registers GRPC services.
func (p *UserAgentKVPlugin) RegisterGRPCAPI(r apiserver.GRPCRouter) {
	r.RegisterService(&_UserAgentKV_serviceDesc, p)
}

type userAgentKVRequest map[string]interface{}

// RESTUserAgentKV is a REST handler for UserAgentKV requests.
func (p *UserAgentKVPlugin) RESTUserAgentKV(c echo.Context) error {
	var data userAgentKVRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	if err := data.validateKVRequest(); err != nil {
		return errutil.ToHTTPError(err)
	}

	switch op := data["operation"]; op {
	case UserAgentKVOperationStore:
		return p.storeKeyValue(c, data["key"].(string), data["value"].(string))
	case UserAgentKVOperationRetrieve:
		if key, ok := data["key"].(string); ok && key != "" {
			return p.retrieveValue(c, key)
		}

		if keys, ok := data["key"].([]string); ok && len(keys) != 0 {
			return p.retrieveValues(c, keys)
		}

		return p.retrieveKVPs(c)
	case UserAgentKVOperationDelete:
		return p.deleteKey(c, data["key"].(string))
	}

	return nil
}

func (p *UserAgentKVPlugin) storeKeyValue(c echo.Context, key string, value string) error {
	if _, err := p.StoreKeyValue(c.Request().Context(), &StoreKeyValueRequest{
		Key:   key,
		Value: value,
	}); err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.NoContent(http.StatusOK)
}

func (p *UserAgentKVPlugin) retrieveValue(c echo.Context, key string) error {
	kv, err := p.RetrieveValues(
		c.Request().Context(),
		&RetrieveValuesRequest{Keys: []string{key}},
	)
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	if len(kv.Values) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No user agent key: %v", key))
	}

	return c.JSON(http.StatusOK, map[string]string{"value": kv.Values[0]})
}

func (p *UserAgentKVPlugin) retrieveValues(c echo.Context, keys []string) error {
	response, err := p.RetrieveValues(c.Request().Context(), &RetrieveValuesRequest{Keys: keys})
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

func (p *UserAgentKVPlugin) retrieveKVPs(c echo.Context) error {
	response, err := p.RetrieveKVPs(c.Request().Context(), &types.Empty{})
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

func (p *UserAgentKVPlugin) deleteKey(c echo.Context, key string) error {
	if _, err := p.DeleteKey(c.Request().Context(), &DeleteKeyRequest{Key: key}); err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.NoContent(http.StatusOK)
}

func (data userAgentKVRequest) validateKVRequest() error {
	if _, ok := data["operation"]; !ok {
		return errutil.ErrorBadRequest("Key/value store API needs 'operation' parameter")
	}

	if _, ok := data["key"]; !ok {
		return errutil.ErrorBadRequest("Key/value store API needs 'key' parameter")
	}

	switch op := data["operation"]; op {
	case UserAgentKVOperationStore, UserAgentKVOperationDelete:
		return data.validateStoreOrDeleteOperation()
	case UserAgentKVOperationRetrieve:
		return data.validateRetrieveOperation()
	default:
		return errutil.ErrorNotFoundf("Invalid Operation %v", op)
	}
}

func (data userAgentKVRequest) validateRetrieveOperation() error {
	errMsg := "retrieve: 'key' must be a string or a list of strings"

	switch key := data["key"].(type) {
	case string:
	case []interface{}:
		keyStrings := make([]string, 0, len(key))
		for _, k := range key {
			if keyString, ok := k.(string); ok {
				keyStrings = append(keyStrings, keyString)
			} else {
				return errutil.ErrorBadRequestf(errMsg)
			}
		}
		data["key"] = keyStrings
	default:
		return errutil.ErrorBadRequestf(errMsg)
	}

	return nil
}

func (data userAgentKVRequest) validateStoreOrDeleteOperation() error {
	if key, ok := data["key"].(string); !ok {
		return errutil.ErrorBadRequestf("store/delete: 'key' must be a string")
	} else if key == "" {
		return errutil.ErrorBadRequestf("store/delete: 'key' must be nonempty")
	}

	return nil
}

// StoreKeyValue stores a value under given key.
// Updates the value if key is already present.
func (p *UserAgentKVPlugin) StoreKeyValue(
	ctx context.Context,
	request *StoreKeyValueRequest,
) (*types.Empty, error) {
	return &types.Empty{}, p.UserAgentKVService.StoreKeyValue(ctx, request.Key, request.Value)
}

// RetrieveValues retrieves values corresponding to the given list of keys.
// The values are returned in an arbitrary order. Keys not present in the store are ignored.
func (p *UserAgentKVPlugin) RetrieveValues(
	ctx context.Context,
	request *RetrieveValuesRequest,
) (res *RetrieveValuesResponse, err error) {
	var values []string
	values, err = p.UserAgentKVService.RetrieveValues(ctx, request.Keys)
	if err == nil {
		res = &RetrieveValuesResponse{Values: values}
	}
	return res, err
}

// DeleteKey deletes the value under the given key.
// Nothing happens if the key is not present.
func (p *UserAgentKVPlugin) DeleteKey(
	ctx context.Context,
	request *DeleteKeyRequest,
) (*types.Empty, error) {
	return &types.Empty{}, p.UserAgentKVService.DeleteKey(ctx, request.Key)
}

// RetrieveKVPs returns the entire store as a list of (key, value) pairs.
func (p *UserAgentKVPlugin) RetrieveKVPs(
	ctx context.Context,
	request *types.Empty,
) (res *RetrieveKVPsResponse, err error) {
	var kvps []*KeyValuePair
	kvps, err = p.UserAgentKVService.RetrieveKVPs(ctx)
	if err == nil {
		res = &RetrieveKVPsResponse{KeyValuePairs: kvps}
	}
	return res, err
}
