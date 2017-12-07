package models

// GracefulRestartTimeType

type GracefulRestartTimeType int

// MakeGracefulRestartTimeType makes GracefulRestartTimeType
func MakeGracefulRestartTimeType() GracefulRestartTimeType {
	var data GracefulRestartTimeType
	return data
}

// InterfaceToGracefulRestartTimeType makes GracefulRestartTimeType from interface
func InterfaceToGracefulRestartTimeType(data interface{}) GracefulRestartTimeType {
	return data.(GracefulRestartTimeType)
}

// InterfaceToGracefulRestartTimeTypeSlice makes a slice of GracefulRestartTimeType from interface
func InterfaceToGracefulRestartTimeTypeSlice(data interface{}) []GracefulRestartTimeType {
	list := data.([]interface{})
	result := MakeGracefulRestartTimeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToGracefulRestartTimeType(item))
	}
	return result
}

// MakeGracefulRestartTimeTypeSlice() makes a slice of GracefulRestartTimeType
func MakeGracefulRestartTimeTypeSlice() []GracefulRestartTimeType {
	return []GracefulRestartTimeType{}
}
