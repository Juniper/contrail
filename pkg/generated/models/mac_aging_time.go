package models

// MACAgingTime

type MACAgingTime int

func MakeMACAgingTime() MACAgingTime {
	var data MACAgingTime
	return data
}

func InterfaceToMACAgingTime(data interface{}) MACAgingTime {
	return data.(MACAgingTime)
}

func InterfaceToMACAgingTimeSlice(data interface{}) []MACAgingTime {
	list := data.([]interface{})
	result := MakeMACAgingTimeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACAgingTime(item))
	}
	return result
}

func MakeMACAgingTimeSlice() []MACAgingTime {
	return []MACAgingTime{}
}
