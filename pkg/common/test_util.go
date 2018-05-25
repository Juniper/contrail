package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

const (
	funcPrefix = "$"
)

//AssertFunction for macros in diff
type AssertFunction func(path string, args, actual interface{}) error

//AssertFunctions for additional test logic.
var AssertFunctions = map[string]AssertFunction{
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

func isStringFunction(key string) bool {
	return strings.HasPrefix(key, funcPrefix)
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

func getAssertFunction(key string) (AssertFunction, error) {
	assertName := strings.TrimPrefix(key, funcPrefix)
	assert, ok := AssertFunctions[assertName]
	if !ok {
		return nil, fmt.Errorf("assert function %s not found", assertName)
	}
	return assert, nil
}

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

//CheckDiff checks diff
func CheckDiff(path string, expected, actual interface{}) error {
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
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
		for key, value := range t {
			err := CheckDiff(path+"."+key, value, actualMap[key])
			if err != nil {
				return err
			}
		}
	case []interface{}:
		actualList, ok := actual.([]interface{})
		if !ok {
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
		if len(t) != len(actualList) {
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
		for i, value := range t {
			found := false
			for _, actualValue := range actualList {
				err := CheckDiff(path+"."+strconv.Itoa(i), value, actualValue)
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
		if float64(t) != InterfaceToFloat(actual) {
			return fmt.Errorf("ffff expected %d but actually we got %f for path %s", t, actual, path)
		}
	default:
		if t != actual {
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
	}
	return nil
}

func logDiff(actual, expected interface{}) {
	log.Debug("expected")
	out, err := yaml.Marshal(expected)
	fmt.Print(string(out), err)
	log.Debug("actual")
	out, err = yaml.Marshal(actual)
	fmt.Print(string(out), err)
}

//AssertEqual test if it is correct
func AssertEqual(t *testing.T, expected, actual interface{}, message string) bool {
	expected = YAMLtoJSONCompat(expected)
	actual = YAMLtoJSONCompat(actual)
	err := CheckDiff("", expected, actual)
	if err != nil {
		logDiff(expected, actual)
	}
	return assert.NoError(t, err, message)
}

// MustDecodeInt64Value tries to decode Int64 from binary form and fails test in case of an error.
func MustDecodeInt64Value(t *testing.T, encodedValue string) interface{} {
	var value int64
	err := binary.Read(bytes.NewBufferString(encodedValue), binary.BigEndian, &value)
	require.NoError(t, err)

	return value
}

// MustDecodeFloat64Value tries to decode Float64 from binary form and fails test in case of an error.
func MustDecodeFloat64Value(t *testing.T, encodedValue string) interface{} {
	var value float64
	err := binary.Read(bytes.NewBufferString(encodedValue), binary.BigEndian, &value)
	require.NoError(t, err)

	return value
}

// MustDecodeBoolValue tries to decode Bool from binary form and fails test in case of an error.
func MustDecodeBoolValue(t *testing.T, encodedValue string) interface{} {
	var value bool
	err := binary.Read(bytes.NewBufferString(encodedValue), binary.BigEndian, &value)
	require.NoError(t, err)

	return value
}
