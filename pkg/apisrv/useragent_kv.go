package apisrv

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
)

const (
	UseragentKVOperationStore    = "STORE"
	UseragentKVOperationRetrieve = "RETRIEVE"
	UseragentKVOperationDelete   = "DELETE"
)

// UseragentKVHandler handles requests to access the useragent key value store.
func (s *Server) UseragentKVHandler(c echo.Context) error {
	var data map[string]interface{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	if err := validateKVRequest(data); err != nil {
		return common.ToHTTPError(err)
	}

	ctx := c.Request().Context()
	switch op := data["operation"]; op {
	case UseragentKVOperationStore:
		if err := s.dbService.StoreKV(ctx, data["key"].(string), data["value"].(string)); err != nil {
			return common.ToHTTPError(err)
		}
		return c.NoContent(http.StatusOK)
	case UseragentKVOperationRetrieve:
		if key, ok := data["key"]; ok {
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
		} else {
			// Empty key provided: retrieve the entire store
			kvps, err := s.dbService.RetrieveKVPs(ctx)
			if err != nil {
				return common.ToHTTPError(err)
			}
			return c.JSON(http.StatusOK, struct {
				Value []models.KeyValuePair `json:"value"`
			}{Value: kvps})
		}
	case UseragentKVOperationDelete:
		err := s.dbService.DeleteKey(ctx, data["key"].(string))
		if err != nil {
			return common.ToHTTPError(err)
		}
		return c.NoContent(http.StatusOK)
	}

	return nil
}

func validateKVRequest(data map[string]interface{}) error {
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

		if data["key"].(string) == "" {
			return common.ErrorBadRequestf("Error: store: 'key' must be nonempty")
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

			break
		}

		return common.ErrorBadRequestf(errMsg)
	case UseragentKVOperationDelete:
		if _, ok := data["key"].(string); !ok {
			return common.ErrorBadRequestf("Error: delete: 'key' must be a string")
		}

		if data["key"].(string) == "" {
			return common.ErrorBadRequestf("Error: delete: 'key' must be nonempty")
		}
	default:
		return common.ErrorNotFoundf("Invalid Operation %v", op)
	}

	return nil
}
