package apisrv

import (
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/labstack/echo"
)

const (
	UseragentKVOperationStore    = "STORE"
	UseragentKVOperationRetrieve = "RETRIEVE"
	UseragentKVOperationDelete   = "DELETE"
)

func (s *Server) useragentKVHandler(c echo.Context) error {
	var data map[string]interface{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	if err := validateKVRequest(data); err != nil {
		return common.ToHTTPError(err)
	}

	switch op := data["operation"]; op {
	case UseragentKVOperationStore:
		if err := s.dbService.StoreKV(c, data["key"].(string), data["value"].(string)); err != nil {
			return common.ToHTTPError(err)
		}
		return c.NoContent(http.StatusOK)
	case UseragentKVOperationRetrieve:
		if key, ok := data["key"]; ok {
			switch key := key.(type) {
			case string:
				val, err := s.dbService.RetrieveValue(c, key)
				if err != nil {
					return common.ToHTTPError(err)
				}
				return c.JSON(http.StatusOK, struct {
					value string `json:"value"`
				}{value: val})
			case []string:
				vals, err := s.dbService.RetrieveValues(c, key)
				if err != nil {
					return common.ToHTTPError(err)
				}
				return c.JSON(http.StatusOK, struct {
					value string `json:"value"`
				}{value: vals})
			}
		} else {
			// Empty key provided: retrieve the entire store
			kvps, err := s.dbService.RetrieveKVPs(c)
			if err != nil {
				return common.ToHTTPError(err)
			}
			return c.JSON(http.StatusOK, kvps)
		}
	case UseragentKVOperationDelete:
		err := s.dbService.DeleteKV(c, data["key"].(string))
		if err != nil {
			return common.ToHTTPError(err)
		}
		return c.NoContent(http.StatusOK)
	}

	return nil
}

func validateKVRequest(data map[string]interface{}) err {
	if _, ok := data["operation"]; !ok {
		return common.ErrorBadRequest("Error: Key/value store API needs 'operation' parameter")
	}

	if _, ok := data["key"]; !ok {
		return common.ErrorBadRequest("Error: Key/value store API needs 'key' parameter")
	}

	// TODO: handle general JSON data for key and value?
	switch op := data["operation"]; op {
	case UseragentKVOperationStore:
		if _, ok := data["key"].(string); !ok {
			return common.ErrorBadRequestf("Error: store: 'key' must be a string")
		}
	case UseragentKVOperationRetrieve:
		errMsg := "Error: retrieve: 'key' must be a string or a list of strings"

		if key, ok := data["key"].(string); ok {
			if key == "" {
				delete(data, "key")
			}

			break
		}

		if keys, ok := data["key"].([]interface{}); ok {
			var keyStrings []string
			for key := range keys {
				if keyString, ok := key.(string); ok {
					keyStrings = append(keyStrings, keyString)
				} else {
					return common.ErrorBadRequestf(errMsg)
				}
			}
			data["key"] = keyStrings

			break
		}

		return common.ErrorBadRequestf(errMsg)
	case UseragentKVOperationDelete:
		if _, ok := data["key"].(string); !ok {
			return common.ErrorBadRequestf("Error: delete: 'key' must be a string")
		}
	default:
		return common.ErrorNotFoundf("Invalid Operation %v", op)
	}

	return nil
}
