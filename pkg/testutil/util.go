package testutil

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

//CheckDiff checks diff
// nolint: gocyclo
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
		if float64(t) != common.InterfaceToFloat(actual) {
			return fmt.Errorf("ffff expected %d but actually we got %f for path %s", t, actual, path)
		}
	default:
		if t != actual {
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
	}
	return nil
}

func logDiff(expected, actual interface{}) {
	log.Debug("expected")
	out, err := yaml.Marshal(expected)
	log.Debug(string(out), err)
	log.Debug("actual")
	out, err = yaml.Marshal(actual)
	log.Debug(string(out), err)
}

//AssertEqual test if it is correct
func AssertEqual(t *testing.T, expected, actual interface{}, message string) bool {
	expected = common.YAMLtoJSONCompat(expected)
	actual = common.YAMLtoJSONCompat(actual)
	err := CheckDiff("", expected, actual)
	if err != nil {
		logDiff(expected, actual)
	}
	return assert.NoError(t, err, message)
}
