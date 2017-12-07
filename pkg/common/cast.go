package common

//InterfaceToInt makes an int from interface
func InterfaceToInt(i interface{}) int {
	switch t := i.(type) {
	case []byte:
		return int(t[0])
	case int:
		return t
	case int64:
		return int(t)
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
