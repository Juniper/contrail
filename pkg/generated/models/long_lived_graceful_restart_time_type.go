package models

// LongLivedGracefulRestartTimeType

type LongLivedGracefulRestartTimeType int

// MakeLongLivedGracefulRestartTimeType makes LongLivedGracefulRestartTimeType
func MakeLongLivedGracefulRestartTimeType() LongLivedGracefulRestartTimeType {
	var data LongLivedGracefulRestartTimeType
	return data
}

// InterfaceToLongLivedGracefulRestartTimeType makes LongLivedGracefulRestartTimeType from interface
func InterfaceToLongLivedGracefulRestartTimeType(data interface{}) LongLivedGracefulRestartTimeType {
	return data.(LongLivedGracefulRestartTimeType)
}

// InterfaceToLongLivedGracefulRestartTimeTypeSlice makes a slice of LongLivedGracefulRestartTimeType from interface
func InterfaceToLongLivedGracefulRestartTimeTypeSlice(data interface{}) []LongLivedGracefulRestartTimeType {
	list := data.([]interface{})
	result := MakeLongLivedGracefulRestartTimeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLongLivedGracefulRestartTimeType(item))
	}
	return result
}

// MakeLongLivedGracefulRestartTimeTypeSlice() makes a slice of LongLivedGracefulRestartTimeType
func MakeLongLivedGracefulRestartTimeTypeSlice() []LongLivedGracefulRestartTimeType {
	return []LongLivedGracefulRestartTimeType{}
}
