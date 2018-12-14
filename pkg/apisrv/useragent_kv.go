package apisrv

import (
	"fmt"
	"net/http"

	"github.com/gogo/protobuf/types"
	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/services"
)

// UserAgent key value store operations.
const (
	UserAgentKVOperationStore    = "STORE"
	UserAgentKVOperationRetrieve = "RETRIEVE"
	UserAgentKVOperationDelete   = "DELETE"
)

type userAgentKVRequest map[string]interface{}

func (s *Server) userAgentKVHandler(c echo.Context) error {
	var data userAgentKVRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	if err := data.validateKVRequest(); err != nil {
		return errutil.ToHTTPError(err)
	}

	switch op := data["operation"]; op {
	case UserAgentKVOperationStore:
		return s.storeKeyValue(c, data["key"].(string), data["value"].(string))
	case UserAgentKVOperationRetrieve:
		if key, ok := data["key"].(string); ok && key != "" {
			return s.retrieveValue(c, key)
		}

		if keys, ok := data["key"].([]string); ok && len(keys) != 0 {
			return s.retrieveValues(c, keys)
		}

		return s.retrieveKVPs(c)
	case UserAgentKVOperationDelete:
		return s.deleteKey(c, data["key"].(string))
	}

	return nil
}

func (s *Server) storeKeyValue(c echo.Context, key string, value string) error {
	if _, err := s.UserAgentKVServer.StoreKeyValue(c.Request().Context(), &services.StoreKeyValueRequest{
		Key:   key,
		Value: value,
	}); err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.NoContent(http.StatusOK)
}

func (s *Server) retrieveValue(c echo.Context, key string) error {
	kv, err := s.UserAgentKVServer.RetrieveValues(c.Request().Context(), &services.RetrieveValuesRequest{Keys: []string{key}})
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	if len(kv.Values) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No user agent key: %v", key))
	}

	return c.JSON(http.StatusOK, &services.RetrieveValueResponse{Value: kv.Values[0]})
}

func (s *Server) retrieveValues(c echo.Context, keys []string) error {
	response, err := s.UserAgentKVServer.RetrieveValues(c.Request().Context(), &services.RetrieveValuesRequest{Keys: keys})
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) retrieveKVPs(c echo.Context) error {
	response, err := s.UserAgentKVServer.RetrieveKVPs(c.Request().Context(), &types.Empty{})
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) deleteKey(c echo.Context, key string) error {
	if _, err := s.UserAgentKVServer.DeleteKey(c.Request().Context(), &services.DeleteKeyRequest{Key: key}); err != nil {
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
