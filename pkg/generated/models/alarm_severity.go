package models

// AlarmSeverity

type AlarmSeverity int

func MakeAlarmSeverity() AlarmSeverity {
	var data AlarmSeverity
	return data
}

func InterfaceToAlarmSeverity(data interface{}) AlarmSeverity {
	return data.(AlarmSeverity)
}

func InterfaceToAlarmSeveritySlice(data interface{}) []AlarmSeverity {
	list := data.([]interface{})
	result := MakeAlarmSeveritySlice()
	for _, item := range list {
		result = append(result, InterfaceToAlarmSeverity(item))
	}
	return result
}

func MakeAlarmSeveritySlice() []AlarmSeverity {
	return []AlarmSeverity{}
}
