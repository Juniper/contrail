package models

// AlarmOperation

type AlarmOperation string

// MakeAlarmOperation makes AlarmOperation
func MakeAlarmOperation() AlarmOperation {
	var data AlarmOperation
	return data
}

// InterfaceToAlarmOperation makes AlarmOperation from interface
func InterfaceToAlarmOperation(data interface{}) AlarmOperation {
	return data.(AlarmOperation)
}

// InterfaceToAlarmOperationSlice makes a slice of AlarmOperation from interface
func InterfaceToAlarmOperationSlice(data interface{}) []AlarmOperation {
	list := data.([]interface{})
	result := MakeAlarmOperationSlice()
	for _, item := range list {
		result = append(result, InterfaceToAlarmOperation(item))
	}
	return result
}

// MakeAlarmOperationSlice() makes a slice of AlarmOperation
func MakeAlarmOperationSlice() []AlarmOperation {
	return []AlarmOperation{}
}
