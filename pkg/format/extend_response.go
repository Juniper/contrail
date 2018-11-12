package format

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// ExtendResponseWithHref extends @response with href  bases on @collectionURL and
// value of @keyName field in object
func ExtendResponseWithHref(response interface{}, collectionURL, keyName string) (interface{}, error) {
	url, err := url.Parse(collectionURL)
	if err != nil {
		return nil, err
	}

	res, err := extendWithKind(reflect.ValueOf(response), url, keyName)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func structToMap(v reflect.Value) map[string]interface{} {
	res := make(map[string]interface{})
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		name := t.Field(i).Name
		if alias, ok := t.Field(i).Tag.Lookup("json"); ok {
			aliases := strings.Split(alias, ",")
			name = aliases[0]
			for j := 1; j < len(aliases); j++ {
				if aliases[j] == "-" || (aliases[j] == "omitempty" && isEmptyValue(v.Field(i))) {
					name = ""
				}
			}
		}
		if name != "" {
			res[name] = v.Field(i).Interface()
		}
	}

	return res
}

// ToMap serialize structs and ponter to struct to map[string]interface{}
func ToMap(v interface{}) interface{} {
	return toMap(reflect.ValueOf(v))
}
func toMap(v reflect.Value) interface{} {
	switch v.Type().Kind() {
	case reflect.Ptr:
		return toMap(reflect.Indirect(v))
	case reflect.Struct:
		return structToMap(v)
	default:
		return v.Interface()
	}
}

func extendWithKind(v reflect.Value, collectionURL *url.URL, keyName string) (interface{}, error) {
	switch v.Type().Kind() {
	case reflect.Struct:
		return extendStruct(v, collectionURL, keyName)
	case reflect.Ptr:
		return extendWithKind(reflect.Indirect(v), collectionURL, keyName)
	case reflect.Map:
		return extendMap(v, collectionURL, keyName)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return extendArray(v, collectionURL, keyName)
	case reflect.Interface:
		return extendWithKind(v.Elem(), collectionURL, keyName)
	default:
		return nil, fmt.Errorf("invalid type: %s", v.Type().Kind())
	}
}

func extendArray(v reflect.Value, collectionURL *url.URL, keyName string) (interface{}, error) {
	res := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		obj, err := extendWithKind(v.Index(i), collectionURL, keyName)
		if err != nil {
			return nil, err
		}
		res[i] = obj
	}

	return res, nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func extendStruct(v reflect.Value, collectionURL *url.URL, keyName string) (interface{}, error) {
	res := structToMap(v)
	key, ok := res[keyName].(string)
	if !ok {
		return nil, fmt.Errorf("failed to cast key value to string. kind: %s", reflect.TypeOf(key).Kind())
	}

	if _, ok := res["href"]; !ok && key != "" {
		keyURL, err := url.Parse(key)
		if err != nil {
			return nil, err
		}
		res["href"] = collectionURL.ResolveReference(keyURL).String()
	}

	return res, nil
}

func extendMap(v reflect.Value, collectionURL *url.URL, keyName string) (interface{}, error) {
	res := make(map[string]interface{})

	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		processed, err := extendWithKind(value, collectionURL, keyName)
		if err != nil {
			res[key.String()] = value.Interface()
		} else {
			res[key.String()] = processed
		}
	}

	return res, nil
}
