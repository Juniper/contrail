package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testJSONStruct struct {
	UUID   string  `json:"uuid"`
	Value  string  `json:"value"`
	Value2 *string `json:"value2,omitempty"`
}

func TestExtendStructWithHref(t *testing.T) {
	data := testJSONStruct{
		UUID:  "vn_blue",
		Value: "net",
	}

	res, err := ExtendResponseWithHref(data, "/virtual_networks/", "uuid")
	assert.NoError(t, err)
	assert.IsType(t, map[string]interface{}{}, res)
	resMap, castOk := res.(map[string]interface{})
	assert.True(t, castOk)
	assert.Equal(t, "vn_blue", resMap["uuid"])
	assert.Equal(t, "net", resMap["value"])
	assert.Equal(t, "/virtual_networks/vn_blue", resMap["href"])
	if val, ok := resMap["value2"]; ok {
		t.Errorf("value2 is exist in map: %v", val)
	}

	res, err = ExtendResponseWithHref(&data, "http://127.0.0.1:8080/virtual_networks/34234", "uuid")
	assert.NoError(t, err)
	assert.IsType(t, map[string]interface{}{}, res)
	resMap, castOk = res.(map[string]interface{})
	assert.True(t, castOk)
	assert.Equal(t, "vn_blue", resMap["uuid"])
	assert.Equal(t, "net", resMap["value"])
	assert.Equal(t, "http://127.0.0.1:8080/virtual_networks/vn_blue", resMap["href"])
	if val, ok := resMap["value2"]; ok {
		t.Errorf("value2 is exist in map: %v", val)
	}
}

func TestExtendStructsSliceWithHref(t *testing.T) {
	value2 := "val"
	data := []testJSONStruct{
		{
			UUID:   "vn_blue",
			Value:  "net",
			Value2: &value2,
		},
		{
			UUID:  "vn_red",
			Value: "net2",
		},
	}

	res, err := ExtendResponseWithHref(data, "/virtual_networks/", "uuid")
	assert.NoError(t, err)
	assert.IsType(t, []interface{}{}, res)
	resArray, castOk := res.([]interface{})
	assert.True(t, castOk)

	assert.IsType(t, map[string]interface{}{}, resArray[0])
	resMap, castOk := resArray[0].(map[string]interface{})
	assert.True(t, castOk)
	assert.Equal(t, "vn_blue", resMap["uuid"])
	assert.Equal(t, "net", resMap["value"])
	assert.Equal(t, "val", *(resMap["value2"].(*string)))
	assert.Equal(t, "/virtual_networks/vn_blue", resMap["href"])

	assert.IsType(t, map[string]interface{}{}, resArray[1])
	resMap, castOk = resArray[1].(map[string]interface{})
	assert.True(t, castOk)
	assert.Equal(t, "vn_red", resMap["uuid"])
	assert.Equal(t, "net2", resMap["value"])
	assert.Equal(t, "/virtual_networks/vn_red", resMap["href"])
	if val, ok := resMap["value2"]; ok {
		t.Errorf("value2 is exist in map: %v", val)
	}

	res, err = ExtendResponseWithHref(&data, "/virtual_networks/", "uuid")
	assert.NoError(t, err)
	assert.IsType(t, []interface{}{}, res)
	resArray, castOk = res.([]interface{})
	assert.True(t, castOk)

	assert.IsType(t, map[string]interface{}{}, resArray[0])
	resMap, castOk = resArray[0].(map[string]interface{})
	assert.True(t, castOk)
	assert.Equal(t, "vn_blue", resMap["uuid"])
	assert.Equal(t, "net", resMap["value"])
	assert.Equal(t, "val", *(resMap["value2"].(*string)))
	assert.Equal(t, "/virtual_networks/vn_blue", resMap["href"])

	assert.IsType(t, map[string]interface{}{}, resArray[1])
	resMap, castOk = resArray[1].(map[string]interface{})
	assert.True(t, castOk)
	assert.Equal(t, "vn_red", resMap["uuid"])
	assert.Equal(t, "net2", resMap["value"])
	assert.Equal(t, "/virtual_networks/vn_red", resMap["href"])
	if val, ok := resMap["value2"]; ok {
		t.Errorf("value2 is exist in map: %v", val)
	}
}

type responseTest2 struct {
	Test *testJSONStruct `json:"test,omitempty"`
}

func TestExtendStructFieldWithHref(t *testing.T) {
	data := testJSONStruct{
		UUID:  "vn_green",
		Value: "nGreen",
	}

	data2 := responseTest2{
		Test: &data,
	}

	res, err := ExtendResponseWithHref(ToMap(data2), "/virtual_networks/", "uuid")
	assert.NoError(t, err)
	assert.IsType(t, map[string]interface{}{}, res)

	resMap, castOk := res.(map[string]interface{})
	assert.True(t, castOk)
	if _, ok := resMap["test"]; !ok {
		t.Error("test is not exist in map")
	}
	dataValue := resMap["test"]
	assert.IsType(t, map[string]interface{}{}, dataValue)
	dataMap, castOk := dataValue.(map[string]interface{})
	assert.True(t, castOk)

	assert.Equal(t, "vn_green", dataMap["uuid"])
	assert.Equal(t, "nGreen", dataMap["value"])
	assert.Equal(t, "/virtual_networks/vn_green", dataMap["href"])
	if val, ok := dataMap["value2"]; ok {
		t.Errorf("value2 is exist in map: %v", val)
	}

	res, err = ExtendResponseWithHref(ToMap(&data2), "/virtual_networks/", "uuid")
	assert.NoError(t, err)
	assert.IsType(t, map[string]interface{}{}, res)

	resMap, castOk = res.(map[string]interface{})
	assert.True(t, castOk)
	if _, ok := resMap["test"]; !ok {
		t.Error("test is not exist in map")
	}
	dataValue = resMap["test"]
	assert.IsType(t, map[string]interface{}{}, dataValue)
	dataMap, castOk = dataValue.(map[string]interface{})
	assert.True(t, castOk)

	assert.Equal(t, "vn_green", dataMap["uuid"])
	assert.Equal(t, "nGreen", dataMap["value"])
	assert.Equal(t, "/virtual_networks/vn_green", dataMap["href"])
	if val, ok := dataMap["value2"]; ok {
		t.Errorf("value2 is exist in map: %v", val)
	}
}
