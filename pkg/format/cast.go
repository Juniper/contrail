package format

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//InterfaceToInt makes an int from interface with error reporting.
func InterfaceToInt(i interface{}) int {
	var ret, _ = InterfaceToIntE(i)
	return ret
}

//InterfaceToIntE makes an int from interface with error reporting with error reporting.
// nolint: gocyclo
func InterfaceToIntE(i interface{}) (int, error) {
	var err error
	var n int
	var n64 int64
	if i == nil {
		return 0, err
	}
	switch t := i.(type) {
	case []byte:
		n, err = strconv.Atoi(string(t))
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to int", t)
		}
		return n, err
	case string:
		n, err = strconv.Atoi(t)
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to int", t)
		}
		return n, err
	case int:
		return t, err
	case int64:
		return int(t), err
	case float64:
		return int(t), err
	case json.Number:
		n64, err = t.Int64()
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to int", t)
		}
		return int(n64), err
	case nil:
	default:
		err = fmt.Errorf("could not convert %#v to int", i)
		logrus.Debug(err.Error())
	}
	return 0, err
}

//InterfaceToInt64 makes an int64 from interface.
func InterfaceToInt64(i interface{}) int64 {
	var ret, _ = InterfaceToInt64E(i)
	return ret
}

//InterfaceToInt64E makes an int64 from interface with error reporting with error reporting.
// nolint: gocyclo
func InterfaceToInt64E(i interface{}) (int64, error) {
	var err error
	var i64 int64
	if i == nil {
		return 0, err
	}
	switch t := i.(type) {
	case []byte:
		i64, err = strconv.ParseInt(string(t), 10, 64)
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to int64", t)
		}
		return i64, err
	case string:
		i64, err = strconv.ParseInt(t, 10, 64)
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to int64", t)
		}
		return i64, err
	case int:
		return int64(t), err
	case int32:
		return int64(t), err
	case int64:
		return t, err
	case float64:
		return int64(t), err
	case json.Number:
		i64, err = t.Int64()
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to int64", t)
		}
		return i64, err
	case nil:
	default:
		err = fmt.Errorf("could not convert (%T) %#v to int64", i, i)
		logrus.Debug(err.Error())
	}
	return 0, err
}

//InterfaceToUint64 makes a uint64 from interface.
func InterfaceToUint64(i interface{}) uint64 {
	var ret, _ = InterfaceToUint64E(i)
	return ret
}

//InterfaceToUint64E makes a uint64 from interface with error reporting.
// nolint: gocyclo
func InterfaceToUint64E(i interface{}) (uint64, error) {
	var err error
	var u64 uint64
	if i == nil {
		return 0, err
	}
	switch t := i.(type) {
	case []byte:
		u64, err = strconv.ParseUint(string(t), 10, 64)
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to uint64", t)
		}
		return u64, err
	case string:
		u64, err = strconv.ParseUint(t, 10, 64)
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to uint64", t)
		}
		return u64, err
	case int:
		return uint64(t), err
	case int64:
		return uint64(t), err
	case uint:
		return uint64(t), err
	case uint64:
		return t, err
	case float64:
		return uint64(t), err
	case json.Number:
		u64, err = strconv.ParseUint(t.String(), 10, 64)
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to uint64", t)
		}
		return u64, err
	case nil:
	default:
		err = fmt.Errorf("could not convert %#v to uint64", i)
		logrus.Debug(err.Error())
	}
	return 0, err
}

//InterfaceToBool makes a bool from interface.
func InterfaceToBool(i interface{}) bool {
	var ret, _ = InterfaceToBoolE(i)
	return ret
}

//InterfaceToBoolE makes a bool from interface with error reporting.
func InterfaceToBoolE(i interface{}) (bool, error) {
	var err error
	var b bool
	switch t := i.(type) {
	case []byte:
		b, err = strconv.ParseBool(string(t))
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to bool", t)
		}
		return b, err
	case string:
		b, err = strconv.ParseBool(t)
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to bool", t)
		}
		return b, err
	case bool:
		return t, err
	case int64:
		return t == 1, err
	case float64:
		return t == 1, err
	case nil:
	default:
		err = fmt.Errorf("could not convert %#v to bool", i)
		logrus.Debug(err.Error())
	}
	return false, err
}

//InterfaceToString makes a string from interface.
func InterfaceToString(i interface{}) string {
	var ret, _ = InterfaceToStringE(i)
	return ret
}

//InterfaceToStringE makes a string from interface with error reporting.
func InterfaceToStringE(i interface{}) (string, error) {
	var err error
	switch t := i.(type) {
	case []byte:
		return string(t), err
	case string:
		return t, err
	case nil:
	default:
		err = fmt.Errorf("could not convert %#v to string", i)
		logrus.Debug(err.Error())
	}
	return "", err
}

//InterfaceToStringList makes a string list from interface.
func InterfaceToStringList(i interface{}) []string {
	var ret, _ = InterfaceToStringListE(i)
	return ret
}

//InterfaceToStringListE makes a string list from interface with error reporting.
func InterfaceToStringListE(i interface{}) ([]string, error) {
	var err error
	switch t := i.(type) {
	case []string:
		return t, err
	case []interface{}:
		result := []string{}
		for _, s := range t {
			var str string
			var lastError error
			str, lastError = InterfaceToStringE(s)
			if lastError != nil {
				logrus.WithError(err).Debugf("could not convert %#v to []string", i)
				err = lastError
			}
			result = append(result, str)
		}
		return result, err
	case nil:
		return nil, err
	default:
		err = fmt.Errorf("could not convert %#v to []string", i)
		logrus.Debug(err.Error())
	}
	return nil, err
}

//InterfaceToInt64List makes a int64 list from interface.
func InterfaceToInt64List(i interface{}) []int64 {
	var ret, _ = InterfaceToInt64ListE(i)
	return ret
}

//InterfaceToInt64ListE makes a int64 list from interface with error reporting.
func InterfaceToInt64ListE(i interface{}) ([]int64, error) {
	var err error
	switch t := i.(type) {
	case []int64:
		return t, err
	case []interface{}:
		result := []int64{}
		for _, s := range t {
			var i64, lastError = InterfaceToInt64E(s)
			if lastError != nil {
				logrus.WithError(err).Debugf("could not convert %#v to []int64", i)
				err = lastError
			}
			result = append(result, i64)
		}
		return result, err
	case nil:
	default:
		err = fmt.Errorf("could not convert %#v to []int64", i)
		logrus.Debug(err.Error())
	}
	return nil, err
}

//InterfaceToInterfaceList makes a interface list from interface.
func InterfaceToInterfaceList(i interface{}) []interface{} {
	var ret, _ = InterfaceToInterfaceListE(i)
	return ret
}

//InterfaceToInterfaceListE makes a interface list from interface with error reporting.
func InterfaceToInterfaceListE(i interface{}) ([]interface{}, error) {
	var err error
	var ret bool
	t, ret := i.([]interface{})
	if !ret {
		err = fmt.Errorf("could not convert %#v to []interface{}", i)
	}
	return t, err
}

//InterfaceToStringMap makes a string map.
func InterfaceToStringMap(i interface{}) map[string]string {
	var ret, _ = InterfaceToStringMapE(i)
	return ret
}

//InterfaceToStringMapE makes a string map with error reporting.
func InterfaceToStringMapE(i interface{}) (map[string]string, error) {
	var err error
	var ret bool
	t, ret := i.(map[string]string)
	if !ret {
		err = fmt.Errorf("could not convert %#v to map[string]string", i)
	}
	return t, err
}

//InterfaceToInt64Map makes a string map.
func InterfaceToInt64Map(i interface{}) map[string]int64 {
	var ret, _ = InterfaceToInt64MapE(i)
	return ret
}

//InterfaceToInt64MapE makes a string map with error reporting.
func InterfaceToInt64MapE(i interface{}) (map[string]int64, error) {
	var err error
	var ret bool
	t, ret := i.(map[string]int64)
	if !ret {
		err = fmt.Errorf("could not convert %#v to map[string]int64", i)
	}
	return t, err
}

//InterfaceToInterfaceMap makes a interface map.
func InterfaceToInterfaceMap(i interface{}) map[string]interface{} {
	var ret, _ = InterfaceToInterfaceMapE(i)
	return ret
}

//InterfaceToInterfaceMapE makes a interface map with error reporting.
func InterfaceToInterfaceMapE(i interface{}) (map[string]interface{}, error) {
	var err error
	var ret bool
	t, ret := i.(map[string]interface{})
	if !ret {
		err = fmt.Errorf("could not convert %#v to map[string]interface{}", i)
	}
	return t, err
}

//InterfaceToFloat makes a float.
func InterfaceToFloat(i interface{}) float64 {
	var ret, _ = InterfaceToFloatE(i)
	return ret
}

//InterfaceToFloatE makes a float with error reporting.
// nolint: gocyclo
func InterfaceToFloatE(i interface{}) (float64, error) {
	var err error
	var ret bool
	var f float64
	t, ret := i.(float64)
	if !ret {
		return t, fmt.Errorf("could not convert %#v to float64", i)
	}
	switch t := i.(type) {
	case []byte:
		f, err = strconv.ParseFloat(string(t), 64)
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to float", t)
		}
		return f, err
	case string:
		f, err = strconv.ParseFloat(t, 64)
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to float", t)
		}
		return f, err
	case int:
		return float64(t), err
	case int64:
		return float64(t), err
	case float64:
		return t, err
	case nil:
		return 0, err
	case json.Number:
		f, err = t.Float64()
		if err != nil {
			logrus.WithError(err).Debugf("could not convert %#v to float64", t)
		}
		return f, err
	default:
		err = fmt.Errorf("could not convert %#v to float64", i)
		logrus.Debug(err.Error())
	}
	return t, err
}

//InterfaceToBytes makes a bytes from interface.
func InterfaceToBytes(i interface{}) []byte {
	var ret, _ = InterfaceToBytesE(i)
	return ret
}

//InterfaceToBytesE makes a bytes from interface with error reporting.
func InterfaceToBytesE(i interface{}) ([]byte, error) {
	var err error
	switch t := i.(type) { //nolint: errcheck
	case []byte:
		return t, err
	case string:
		return []byte(t), err
	default:
		err = fmt.Errorf("could not convert %#v to []byte", i)
		logrus.Debug(err.Error())
	}
	return []byte{}, err
}

//GetUUIDFromInterface get a UUID from an interface with error reporting.
func GetUUIDFromInterface(rawProperties interface{}) string {
	var ret, _ = GetUUIDFromInterfaceE(rawProperties)
	return ret
}

//GetUUIDFromInterfaceE get a UUID from an interface with error reporting.
func GetUUIDFromInterfaceE(rawProperties interface{}) (string, error) {
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
