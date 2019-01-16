package format

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//InterfaceToInt makes an int from interface
func InterfaceToInt(i interface{}) int {
	if i == nil {
		return 0
	}
	switch t := i.(type) {
	case []byte:
		i, err := strconv.Atoi(string(t))
		if err != nil {
			logrus.WithError(err).Debugf("Could not convert %#v to int", t)
		}
		return i
	case int:
		return t
	case int64:
		return int(t)
	case float64:
		return int(t)
	default:
		logrus.Warnf("Could not convert %#v to int", i)
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
		i, err := strconv.ParseInt(string(t), 10, 64)
		if err != nil {
			logrus.WithError(err).Debugf("Could not convert %#v to int64", t)
		}
		return i
	case int:
		return int64(t)
	case int64:
		return t
	case float64:
		return int64(t)
	case json.Number:
		i, err := t.Int64()
		if err != nil {
			logrus.WithError(err).Debugf("Could not convert %#v to int64", t)
		}
		return i
	default:
		logrus.Warnf("Could not convert %#v to int64", i)
	}
	return 0
}

//InterfaceToUint64 makes an uint64 from interface
// nolint: gocyclo
func InterfaceToUint64(i interface{}) uint64 {
	if i == nil {
		return 0
	}
	switch t := i.(type) {
	case []byte:
		i, err := strconv.ParseUint(string(t), 10, 64)
		if err != nil {
			logrus.WithError(err).Debugf("Could not convert %#v to uint64", t)
		}
		return i
	case int:
		return uint64(t)
	case int64:
		return uint64(t)
	case uint:
		return uint64(t)
	case uint64:
		return t
	case float64:
		return uint64(t)
	case json.Number:
		i, err := strconv.ParseUint(t.String(), 10, 64)
		if err != nil {
			logrus.WithError(err).Debugf("Could not convert %#v to uint64", t)
		}
		return i
	default:
		logrus.Warnf("Could not convert %#v to uint64", i)
	}
	return 0
}

//InterfaceToBool makes a bool from interface
func InterfaceToBool(i interface{}) bool {
	switch t := i.(type) {
	case []byte:
		b, err := strconv.ParseBool(string(t))
		if err != nil {
			logrus.WithError(err).Debugf("Could not convert %#v to bool", t)
		}
		return b
	case bool:
		return t
	case int64:
		return t == 1
	case float64:
		return t == 1
	default:
		logrus.Warnf("Could not convert %#v to bool", i)
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
	case nil:
		return ""
	default:
		logrus.Warnf("Could not convert %#v to string", i)
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
	default:
		logrus.Warnf("Could not convert %#v to []string", i)
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
	default:
		logrus.Warnf("Could not convert %#v to []int64", i)
	}
	return nil
}

//InterfaceToInterfaceList makes a interface list from interface
func InterfaceToInterfaceList(i interface{}) []interface{} {
	t, _ := i.([]interface{}) //nolint: errcheck
	return t
}

//InterfaceToStringMap makes a string map.
func InterfaceToStringMap(i interface{}) map[string]string {
	t, _ := i.(map[string]string) //nolint: errcheck
	return t
}

//InterfaceToInterfaceMap makes a interface map.
func InterfaceToInterfaceMap(i interface{}) map[string]interface{} {
	t, _ := i.(map[string]interface{}) //nolint: errcheck
	return t
}

//InterfaceToFloat makes a float.
func InterfaceToFloat(i interface{}) float64 {
	t, _ := i.(float64) //nolint: errcheck
	switch t := i.(type) {
	case []byte:
		f, err := strconv.ParseFloat(string(t), 64)
		if err != nil {
			logrus.WithError(err).Debugf("Could not convert %#v to float ", t)
		}
		return f
	case int:
		return float64(t)
	case int64:
		return float64(t)
	case float64:
		return t
	default:
		logrus.Warnf("Could not convert %#v to float", i)
	}
	return t
}

//InterfaceToBytes makes a bytes from interface
func InterfaceToBytes(i interface{}) []byte {
	switch t := i.(type) { //nolint: errcheck
	case []byte:
		return t
	case string:
		return []byte(t)
	default:
		logrus.Warnf("Could not convert %#v to []byte", i)
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
		return "", fmt.Errorf("value UUID should be string instead of %T", uuid)
	}
	return uuid, nil
}
