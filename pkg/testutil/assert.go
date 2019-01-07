package testutil

import (
	"fmt"
	"github.com/Juniper/contrail/pkg/constants"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/format"
)

const (
	funcPrefix = "$"
)

// AssertContainsStrings asserts that object contains all expected strings.
func AssertContainsStrings(t *testing.T, expected []string, actual interface{}) {
	for _, e := range expected {
		assert.Contains(t, actual, e)
	}
}

// NotNil matches any non-nil value.
func NotNil() gomock.Matcher {
	return gomock.Not(gomock.Nil())
}

// assertFunction for macros in diff.
type assertFunction func(path string, args, actual interface{}) error

// assertFunctions for additional test logic.
var assertFunctions = map[string]assertFunction{
	"any": func(path string, args, actual interface{}) error {
		return nil
	},
	"null": func(path string, _, actual interface{}) error {
		if actual != nil {
			return errors.Errorf("expected null but got %s on path %s", actual, path)
		}
		return nil
	},
	"number": func(path string, _, actual interface{}) error {
		switch actual.(type) {
		case int64, int, float64:
			return nil
		}
		return errors.Errorf("expected number but got %s on path %s", actual, path)
	},
	"uuid": func(path string, _, actual interface{}) error {
		if val, ok := actual.(string); ok {
			if _, err := uuid.Parse(val); err != nil {
				return errors.Errorf("expected uuid but got %s on path %s (error: %s)", actual, path, err)
			}
			return nil
		}
		return errors.Errorf("expected uuid string but got %s on path %s", actual, path)
	},
	"mac_address": func(path string, _, actual interface{}) error {
		if val, ok := actual.(string); ok {
			if _, err := net.ParseMAC(val); err != nil {
				return errors.Errorf("expected mac address but got %s on path %s (error: %s)", actual, path, err)
			}
			return nil
		}
		return errors.Errorf("expected mac address string string but got %s on path %s", actual, path)
	},
	"ip_address": func(path string, _, actual interface{}) error {
		if val, ok := actual.(string); ok {
			if ip := net.ParseIP(val); ip == nil {
				return errors.Errorf("expected ip address but got %s on path %s", actual, path)
			}
			return nil
		}
		return errors.Errorf("expected ip address string but got %s on path %s", actual, path)
	},
	"datetime_iso": func(path string, _, actual interface{}) error {
		if val, ok := actual.(string); ok {
			if _, err := time.Parse(constants.ISO8601TimeFormat, val); err != nil {
				return errors.Errorf("expected datetime stamp RFC3339 but got %s on path %s", actual, path)
			}
			return nil
		}
		return errors.Errorf("expected datetime stamp string but got %s on path %s", actual, path)
	},
}

// AssertEqual asserts that expected and actual objects are equal, performing comparison recursively.
// For lists and maps, it iterates over expected values, ignoring additional values in actual object.
func AssertEqual(t *testing.T, expected, actual interface{}, msg ...string) bool {
	return assert.NoError(
		t,
		IsObjectSubsetOf(expected, actual),
		fmt.Sprintf(
			"%s: objects not equal:\nexpected: %+v\nactual: %+v",
			strings.Join(msg, ", "),
			format.MustYAML(expected),
			format.MustYAML(actual),
		),
	)
}

// IsObjectSubsetOf verifies if "subset" structure contains all fields described
// in "of" structure and throws an error in case if it doesn't.
func IsObjectSubsetOf(subset, of interface{}) error {
	subset = fileutil.YAMLtoJSONCompat(subset)
	of = fileutil.YAMLtoJSONCompat(of)
	return checkDiff("", subset, of)
}

// nolint: gocyclo
func checkDiff(path string, expected, actual interface{}) error {
	if expected == nil {
		return nil
	}
	if isFunction(expected) {
		return runFunction(path, expected, actual)
	}
	switch t := expected.(type) {
	case map[interface{}]interface{}:
		actualMap, ok := actual.(map[string]interface{})
		if !ok {
			return errorWithFields(t, actual, path)
		}
		for keyI, value := range t {
			key := fmt.Sprint(keyI)
			if err := checkDiff(fmt.Sprintf("%s.%s", path, key), value, actualMap[key]); err != nil {
				return err
			}
		}
	case map[string]interface{}:
		actualMap, ok := actual.(map[string]interface{})
		if !ok {
			return errorWithFields(t, actual, path)
		}
		for key, value := range t {
			if err := checkDiff(fmt.Sprintf("%s.%s", path, key), value, actualMap[key]); err != nil {
				return err
			}
		}
	case []interface{}:
		actualList, ok := actual.([]interface{})
		if !ok {
			return errorWithFields(t, actual, path)
		}
		if len(t) != len(actualList) {
			return errorWithFields(t, actual, path)
		}
		for i, value := range t {
			var mErr errutil.MultiError
			found := false
			for _, actualValue := range actualList {
				err := checkDiff(path+"."+strconv.Itoa(i), value, actualValue)
				if err == nil {
					found = true
					break
				}
				mErr = append(mErr, err)
			}
			if !found {
				return fmt.Errorf("%s not found, last err: %v", path+"."+strconv.Itoa(i), mErr)
			}
		}
	case int:
		if float64(t) != format.InterfaceToFloat(actual) {
			return errorWithFields(t, actual, path)
		}
	default:
		if t != actual {
			return errorWithFields(t, actual, path)
		}
	}
	return nil
}

func isFunction(expected interface{}) bool {
	switch t := expected.(type) {
	case map[string]interface{}:
		for key := range t {
			if isStringFunction(key) {
				return true
			}
		}
	case string:
		return isStringFunction(t)
	}
	return false
}

func isStringFunction(key string) bool {
	return strings.HasPrefix(key, funcPrefix)
}

// nolint: gocyclo
func runFunction(path string, expected, actual interface{}) (err error) {
	switch t := expected.(type) {
	case map[string]interface{}:
		for key := range t {
			if isStringFunction(key) {
				for key, value := range t {
					assert, err := getAssertFunction(key)
					if err != nil {
						return err
					}
					err = assert(path, value, actual)
					if err != nil {
						return err
					}
				}
			}
		}
	case string:
		if isStringFunction(t) {
			assert, err := getAssertFunction(t)
			if err != nil {
				return err
			}
			err = assert(path, nil, actual)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getAssertFunction(key string) (assertFunction, error) {
	assertName := strings.TrimPrefix(key, funcPrefix)
	assert, ok := assertFunctions[assertName]
	if !ok {
		return nil, fmt.Errorf("assert function %s not found", assertName)
	}
	return assert, nil
}

func errorWithFields(expected, actual interface{}, path string) error {
	return fmt.Errorf("expected(%T):\n%v\nactual(%T):\n%v\npath: %s",
		expected,
		format.MustYAML(expected),
		actual,
		format.MustYAML(actual),
		path)
}
