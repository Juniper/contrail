package models

// MACAgingTime

type MACAgingTime int

// MakeMACAgingTime makes MACAgingTime
func MakeMACAgingTime() MACAgingTime {
	var data MACAgingTime
	return data
}

// InterfaceToMACAgingTime makes MACAgingTime from interface
func InterfaceToMACAgingTime(data interface{}) MACAgingTime {
	return data.(MACAgingTime)
}

// InterfaceToMACAgingTimeSlice makes a slice of MACAgingTime from interface
func InterfaceToMACAgingTimeSlice(data interface{}) []MACAgingTime {
	list := data.([]interface{})
	result := MakeMACAgingTimeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACAgingTime(item))
	}
	return result
}

// MakeMACAgingTimeSlice() makes a slice of MACAgingTime
func MakeMACAgingTimeSlice() []MACAgingTime {
	return []MACAgingTime{}
}
