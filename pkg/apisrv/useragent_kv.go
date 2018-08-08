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
		if key, ok := data["key"].(string); ok && key != "" {
			return s.retrieveValue(c, key)
		}

		if keys, ok := data["key"].([]string); ok && len(keys) != 0 {
			return s.retrieveValues(c, keys)
		}

		return s.retrieveKVPs(c)
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

func (s *Server) retrieveValue(c echo.Context, key string) error {
	val, err := s.dbService.RetrieveValue(c.Request().Context(), key)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, struct {
		Value string `json:"value"`
	}{Value: val})
}

func (s *Server) retrieveValues(c echo.Context, keys []string) error {
	vals, err := s.dbService.RetrieveValues(c.Request().Context(), keys)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, struct {
		Value []string `json:"value"`
	}{Value: vals})
}

func (s *Server) retrieveKVPs(c echo.Context) error {
	kvps, err := s.dbService.RetrieveKVPs(c.Request().Context())
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, struct {
		Value []*models.KeyValuePair `json:"value"`
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
		return common.ErrorBadRequest("Key/value store API needs 'operation' parameter")
	}

	if _, ok := data["key"]; !ok {
		return common.ErrorBadRequest("Key/value store API needs 'key' parameter")
	}

	switch op := data["operation"]; op {
	case UseragentKVOperationStore, UseragentKVOperationDelete:
		return data.validateStoreOrDeleteOperation()
	case UseragentKVOperationRetrieve:
		return data.validateRetrieveOperation()
	default:
		return common.ErrorNotFoundf("Invalid Operation %v", op)
	}
}

func (data useragentKVRequest) validateRetrieveOperation() error {
	errMsg := "retrieve: 'key' must be a string or a list of strings"

	switch key := data["key"].(type) {
	case string:
	case []interface{}:
		keyStrings := make([]string, 0, len(key))
		for _, k := range key {
			if keyString, ok := k.(string); ok {
				keyStrings = append(keyStrings, keyString)
			} else {
				return common.ErrorBadRequestf(errMsg)
			}
		}
		data["key"] = keyStrings
	default:
		return common.ErrorBadRequestf(errMsg)
	}

	return nil
}

func (data useragentKVRequest) validateStoreOrDeleteOperation() error {
	if key, ok := data["key"].(string); !ok {
		return common.ErrorBadRequestf("store/delete: 'key' must be a string")
	} else if key == "" {
		return common.ErrorBadRequestf("store/delete: 'key' must be nonempty")
	}

	return nil
}
