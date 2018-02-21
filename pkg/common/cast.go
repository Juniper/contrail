package common

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
