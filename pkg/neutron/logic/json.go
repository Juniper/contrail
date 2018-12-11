package logic

import (
	"encoding/json"
	"fmt"
)

// customJSONUnmarshaler
type customJSONUnmarshaler interface {
	unmarshalCustomJSON(map[string]json.RawMessage) error
}

//type CustomJSONMarshaler interface {
//
//}

// ParseField parses a field from json.RawMessage
func ParseField(rawJSON map[string]json.RawMessage, key string, dst interface{}) error {
	if val, ok := rawJSON[key]; ok {
		if err := json.Unmarshal(val, dst); err != nil {
			return fmt.Errorf("invalid '%s' format: %v", key, err)
		}
		delete(rawJSON, key)
	}
	return nil
}
