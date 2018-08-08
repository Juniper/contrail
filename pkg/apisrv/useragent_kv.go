package apisrv

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
)

// Useragent key value store operations.
const (
	UseragentKVOperationStore    = "STORE"
	UseragentKVOperationRetrieve = "RETRIEVE"
	UseragentKVOperationDelete   = "DELETE"
)

type useragentKVRequest map[string]interface{}

// UseragentKVHandler handles requests to access the useragent key value store.
func (s *Server) UseragentKVHandler(c echo.Context) error {
	var data useragentKVRequest
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	if err := data.validateKVRequest(); err != nil {
		return common.ToHTTPError(err)
	}

	switch op := data["operation"]; op {
	case UseragentKVOperationStore:
		return s.storeKeyValue(c, data["key"].(string), data["value"].(string))
	case UseragentKVOperationRetrieve:
		if key, ok := data["key"]; ok {
			return s.retrieveValue(c, key)
		} else {
			return s.retrieveKVPs(c)
		}
	case UseragentKVOperationDelete:
		return s.deleteKey(c, data["key"].(string))
	}

	return nil
}

func (s *Server) storeKeyValue(c echo.Context, key string, value string) error {
	if err := s.dbService.StoreKV(c.Request().Context(), key, value); err != nil {
		return common.ToHTTPError(err)
	}
	return c.NoContent(http.StatusOK)
}

func (s *Server) retrieveValue(c echo.Context, key interface{}) error {
	ctx := c.Request().Context()
	switch key := key.(type) {
	case string:
		val, err := s.dbService.RetrieveValue(ctx, key)
		if err != nil {
			return common.ToHTTPError(err)
		}
		return c.JSON(http.StatusOK, struct {
			Value string `json:"value"`
		}{Value: val})
	case []string:
		vals, err := s.dbService.RetrieveValues(ctx, key)
		if err != nil {
			return common.ToHTTPError(err)
		}
		return c.JSON(http.StatusOK, struct {
			Value []string `json:"value"`
		}{Value: vals})
	}

	return nil
}

func (s *Server) retrieveKVPs(c echo.Context) error {
	kvps, err := s.dbService.RetrieveKVPs(c.Request().Context())
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, struct {
		Value []models.KeyValuePair `json:"value"`
	}{Value: kvps})
}

func (s *Server) deleteKey(c echo.Context, key string) error {
	if err := s.dbService.DeleteKey(c.Request().Context(), key); err != nil {
		return common.ToHTTPError(err)
	}
	return c.NoContent(http.StatusOK)
}

func (data useragentKVRequest) validateKVRequest() error {
	if _, ok := data["operation"]; !ok {
		return common.ErrorBadRequest("Error: Key/value store API needs 'operation' parameter")
	}

	if _, ok := data["key"]; !ok {
		return common.ErrorBadRequest("Error: Key/value store API needs 'key' parameter")
	}

	switch op := data["operation"]; op {
	case UseragentKVOperationStore:
		return data.validateStoreOperation()
	case UseragentKVOperationRetrieve:
		return data.validateRetrieveOperation()
	case UseragentKVOperationDelete:
		return data.validateDeleteOperation()
	default:
		return common.ErrorNotFoundf("Invalid Operation %v", op)
	}
}

func (data useragentKVRequest) validateStoreOperation() error {
	if _, ok := data["key"].(string); !ok {
		return common.ErrorBadRequestf("Error: store: 'key' must be a string")
	}

	if data["key"].(string) == "" {
		return common.ErrorBadRequestf("Error: store: 'key' must be nonempty")
	}

	return nil
}

func (data useragentKVRequest) validateRetrieveOperation() error {
	errMsg := "Error: retrieve: 'key' must be a string or a list of strings"

	if key, ok := data["key"].(string); ok {
		if key == "" {
			delete(data, "key")
		}

		return nil
	}

	if keys, ok := data["key"].([]interface{}); ok {
		if len(keys) == 0 {
			delete(data, "key")
		} else {
			var keyStrings []string
			for _, key := range keys {
				if keyString, ok := key.(string); ok {
					keyStrings = append(keyStrings, keyString)
				} else {
					return common.ErrorBadRequestf(errMsg)
				}
			}
			data["key"] = keyStrings
		}

		return nil
	}

	return common.ErrorBadRequestf(errMsg)
}

func (data useragentKVRequest) validateDeleteOperation() error {
	if _, ok := data["key"].(string); !ok {
		return common.ErrorBadRequestf("Error: delete: 'key' must be a string")
	}

	if data["key"].(string) == "" {
		return common.ErrorBadRequestf("Error: delete: 'key' must be nonempty")
	}

	return nil
}
