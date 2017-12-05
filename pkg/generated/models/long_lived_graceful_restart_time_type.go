package models

// LongLivedGracefulRestartTimeType

type LongLivedGracefulRestartTimeType int

func MakeLongLivedGracefulRestartTimeType() LongLivedGracefulRestartTimeType {
	var data LongLivedGracefulRestartTimeType
	return data
}

func InterfaceToLongLivedGracefulRestartTimeType(data interface{}) LongLivedGracefulRestartTimeType {
	return data.(LongLivedGracefulRestartTimeType)
}

func InterfaceToLongLivedGracefulRestartTimeTypeSlice(data interface{}) []LongLivedGracefulRestartTimeType {
	list := data.([]interface{})
	result := MakeLongLivedGracefulRestartTimeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLongLivedGracefulRestartTimeType(item))
	}
	return result
}

func MakeLongLivedGracefulRestartTimeTypeSlice() []LongLivedGracefulRestartTimeType {
	return []LongLivedGracefulRestartTimeType{}
}
