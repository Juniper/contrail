package testutil

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/common"
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
	"null": func(path string, args, actual interface{}) error {
		if actual != nil {
			return fmt.Errorf("expecetd null but got %s on path %s", actual, path)
		}
		return nil
	},
	"number": func(path string, args, actual interface{}) error {
		switch actual.(type) {
		case int64, int, float64:
			return nil
		}
		return fmt.Errorf("expecetd integer but got %s on path %s", actual, path)
	},
}

// AssertEqual asserts that expected and actual objects are equal, performing comparison recursively.
// For lists and maps, it iterates over expected values, ignoring additional values in actual object.
func AssertEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	expected = common.YAMLtoJSONCompat(expected)
	actual = common.YAMLtoJSONCompat(actual)

	err := checkDiff("", expected, actual)

	return assert.NoError(
		t,
		err,
		append(
			msgAndArgs,
			fmt.Sprintf("objects not equal:\nexpected: %+v\nactual: %+v",
				format.MustYAML(expected),
				format.MustYAML(actual)),
		)...,
	)
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
	case map[string]interface{}:
		actualMap, ok := actual.(map[string]interface{})
		if !ok {
			return errorWithFields(t, actual, path)
		}
		for key, value := range t {
			err := checkDiff(path+"."+key, value, actualMap[key])
			if err != nil {
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
			found := false
			for _, actualValue := range actualList {
				err := checkDiff(path+"."+strconv.Itoa(i), value, actualValue)
				if err == nil {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("%s not found", path+"."+strconv.Itoa(i))
			}
		}
	case int:
		if float64(t) != common.InterfaceToFloat(actual) {
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
	return fmt.Errorf("expected:\n%v\nactual:\n%v\npath: %s",
		format.MustYAML(expected),
		format.MustYAML(actual),
		path)
}
