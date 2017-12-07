package models

// DnsRecordClassType

type DnsRecordClassType string

// MakeDnsRecordClassType makes DnsRecordClassType
func MakeDnsRecordClassType() DnsRecordClassType {
	var data DnsRecordClassType
	return data
}

// InterfaceToDnsRecordClassType makes DnsRecordClassType from interface
func InterfaceToDnsRecordClassType(data interface{}) DnsRecordClassType {
	return data.(DnsRecordClassType)
}

// InterfaceToDnsRecordClassTypeSlice makes a slice of DnsRecordClassType from interface
func InterfaceToDnsRecordClassTypeSlice(data interface{}) []DnsRecordClassType {
	list := data.([]interface{})
	result := MakeDnsRecordClassTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDnsRecordClassType(item))
	}
	return result
}

// MakeDnsRecordClassTypeSlice() makes a slice of DnsRecordClassType
func MakeDnsRecordClassTypeSlice() []DnsRecordClassType {
	return []DnsRecordClassType{}
}
