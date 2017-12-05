package models

// AlarmOperation

type AlarmOperation string

func MakeAlarmOperation() AlarmOperation {
	var data AlarmOperation
	return data
}

func InterfaceToAlarmOperation(data interface{}) AlarmOperation {
	return data.(AlarmOperation)
}

func InterfaceToAlarmOperationSlice(data interface{}) []AlarmOperation {
	list := data.([]interface{})
	result := MakeAlarmOperationSlice()
	for _, item := range list {
		result = append(result, InterfaceToAlarmOperation(item))
	}
	return result
}

func MakeAlarmOperationSlice() []AlarmOperation {
	return []AlarmOperation{}
}
