package models

// AlarmSeverity

type AlarmSeverity int

// MakeAlarmSeverity makes AlarmSeverity
func MakeAlarmSeverity() AlarmSeverity {
	var data AlarmSeverity
	return data
}

// InterfaceToAlarmSeverity makes AlarmSeverity from interface
func InterfaceToAlarmSeverity(data interface{}) AlarmSeverity {
	return data.(AlarmSeverity)
}

// InterfaceToAlarmSeveritySlice makes a slice of AlarmSeverity from interface
func InterfaceToAlarmSeveritySlice(data interface{}) []AlarmSeverity {
	list := data.([]interface{})
	result := MakeAlarmSeveritySlice()
	for _, item := range list {
		result = append(result, InterfaceToAlarmSeverity(item))
	}
	return result
}

// MakeAlarmSeveritySlice() makes a slice of AlarmSeverity
func MakeAlarmSeveritySlice() []AlarmSeverity {
	return []AlarmSeverity{}
}
