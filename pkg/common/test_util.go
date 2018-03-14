package common

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//CheckDiff checks diff
func CheckDiff(path string, expected, actual interface{}) error {
	if expected == nil {
		return nil
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
		for index, value := range t {
			err := CheckDiff(path+"."+strconv.Itoa(index), value, actualList[index])
			if err != nil {
				return err
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

//AssertEqual test if it is correct
func AssertEqual(t *testing.T, expected, actual interface{}, message string) bool {
	expected = YAMLtoJSONCompat(expected)
	actual = YAMLtoJSONCompat(actual)
	err := CheckDiff("", expected, actual)
	return assert.NoError(t, err, message)
}
