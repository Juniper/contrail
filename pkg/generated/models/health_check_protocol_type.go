package models

// HealthCheckProtocolType

type HealthCheckProtocolType string

// MakeHealthCheckProtocolType makes HealthCheckProtocolType
func MakeHealthCheckProtocolType() HealthCheckProtocolType {
	var data HealthCheckProtocolType
	return data
}

// InterfaceToHealthCheckProtocolType makes HealthCheckProtocolType from interface
func InterfaceToHealthCheckProtocolType(data interface{}) HealthCheckProtocolType {
	return data.(HealthCheckProtocolType)
}

// InterfaceToHealthCheckProtocolTypeSlice makes a slice of HealthCheckProtocolType from interface
func InterfaceToHealthCheckProtocolTypeSlice(data interface{}) []HealthCheckProtocolType {
	list := data.([]interface{})
	result := MakeHealthCheckProtocolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToHealthCheckProtocolType(item))
	}
	return result
}

// MakeHealthCheckProtocolTypeSlice() makes a slice of HealthCheckProtocolType
func MakeHealthCheckProtocolTypeSlice() []HealthCheckProtocolType {
	return []HealthCheckProtocolType{}
}
