package common

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

//InterfaceToInt makes an int from interface
func InterfaceToInt(i interface{}) int {
	if i == nil {
		return 0
	}
	switch t := i.(type) {
	case []byte:
		i, _ := strconv.Atoi(string(t)) // nolint: gas
		return i
	case int:
		return t
	case int64:
		return int(t)
	case float64:
		return int(t)
	}
	return 0
}

//InterfaceToInt64 makes an int64 from interface
func InterfaceToFloat64(i interface{}) float64 {
	if i == nil {
		return 0
	}
	switch t := i.(type) {
	case []byte:
		i, _ := strconv.ParseFloat(string(t), 64) // nolint: gas
		return i
	case int:
		return float64(t)
	case int64:
		return float64(t)
	case float64:
		return t
	}
	return 0
}

//InterfaceToInt64 makes an int64 from interface
func InterfaceToInt64(i interface{}) int64 {
	if i == nil {
		return 0
	}
	switch t := i.(type) {
	case []byte:
		i, _ := strconv.ParseInt(string(t), 10, 64) // nolint: gas
		return i
	case int:
		return int64(t)
	case int64:
		return t
	case float64:
		return int64(t)
	}
	return 0
}

//InterfaceToBool makes a bool from interface
func InterfaceToBool(i interface{}) bool {
	switch t := i.(type) {
	case []byte:
		b, _ := strconv.ParseBool(string(t)) // nolint: gas
		return b
	case bool:
		return t
	case int64:
		return t == 1
	case float64:
		return t == 1
	}
	return false
}

//InterfaceToString makes a string from interface
func InterfaceToString(i interface{}) string {
	switch t := i.(type) {
	case []byte:
		return string(t)
	case string:
		return t
	}
	return ""
}

//InterfaceToStringList makes a string list from interface
func InterfaceToStringList(i interface{}) []string {
	switch t := i.(type) {
	case []string:
		return t
	case []interface{}:
		result := []string{}
		for _, s := range t {
			result = append(result, InterfaceToString(s))
		}
		return result
	}
	return nil
}

//InterfaceToInt64List makes a int64 list from interface
func InterfaceToInt64List(i interface{}) []int64 {
	switch t := i.(type) {
	case []int64:
		return t
	case []interface{}:
		result := []int64{}
		for _, s := range t {
			result = append(result, InterfaceToInt64(s))
		}
		return result
	}
	return nil
}

//InterfaceToFloat64List makes a float64 list from interface
func InterfaceToFloat64List(i interface{}) []float64 {
	switch t := i.(type) {
	case []float64:
		return t
	case []interface{}:
		result := []float64{}
		for _, s := range t {
			result = append(result, InterfaceToFloat64(s))
		}
		return result
	}
	return nil
}

//InterfaceToInterfaceList makes a interface list from interface
func InterfaceToInterfaceList(i interface{}) []interface{} {
	t, _ := i.([]interface{})
	return t
}

//InterfaceToStringMap makes a string map.
func InterfaceToStringMap(i interface{}) map[string]string {
	t, _ := i.(map[string]string)
	return t
}

//InterfaceToInterfaceMap makes a interface map.
func InterfaceToInterfaceMap(i interface{}) map[string]interface{} {
	t, _ := i.(map[string]interface{})
	return t
}

//InterfaceToFloat makes a float.
func InterfaceToFloat(i interface{}) float64 {
	t, _ := i.(float64)
	switch t := i.(type) {
	case []byte:
		f, _ := strconv.ParseFloat(string(t), 64) // nolint: gas
		return f
	case int:
		return float64(t)
	case int64:
		return float64(t)
	case float64:
		return t
	}
	return t
}

//InterfaceToBytes makes a bytes from interface
func InterfaceToBytes(i interface{}) []byte {
	switch t := i.(type) {
	case []byte:
		return t
	case string:
		return []byte(t)
	}
	return []byte{}
}

//GetUUIDFromInterface get a UUID from an interface.
func GetUUIDFromInterface(rawProperties interface{}) (string, error) {
	properties, ok := rawProperties.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid data format: no properties mapping")
	}

	rawUUID, ok := properties["uuid"]
	if !ok {
		return "", errors.New("data does not contain required UUID property")
	}

	uuid, ok := rawUUID.(string)
	if !ok {
		return "", fmt.Errorf("UUID should be string instead of %T", uuid)
	}
	return uuid, nil
}
