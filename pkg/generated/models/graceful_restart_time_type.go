package models

// GracefulRestartTimeType

type GracefulRestartTimeType int

func MakeGracefulRestartTimeType() GracefulRestartTimeType {
	var data GracefulRestartTimeType
	return data
}

func InterfaceToGracefulRestartTimeType(data interface{}) GracefulRestartTimeType {
	return data.(GracefulRestartTimeType)
}

func InterfaceToGracefulRestartTimeTypeSlice(data interface{}) []GracefulRestartTimeType {
	list := data.([]interface{})
	result := MakeGracefulRestartTimeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToGracefulRestartTimeType(item))
	}
	return result
}

func MakeGracefulRestartTimeTypeSlice() []GracefulRestartTimeType {
	return []GracefulRestartTimeType{}
}
