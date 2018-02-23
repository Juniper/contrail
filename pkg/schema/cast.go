package schema

//InterfaceToInt makes an int from interface
func InterfaceToInt(i interface{}) int {
	if i == nil {
		return 0
	}
	switch t := i.(type) {
	case []byte:
		return int(t[0])
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
func InterfaceToInt64(i interface{}) int64 {
	if i == nil {
		return 0
	}
	switch t := i.(type) {
	case []byte:
		return int64(t[0])
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
		return len(t) == 1 && t[0] == 1
	case bool:
		return t
	case int64:
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

//InterfaceToInt64List makes a string list from interface
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
