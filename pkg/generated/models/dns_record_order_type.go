package models

// DnsRecordOrderType

type DnsRecordOrderType string

// MakeDnsRecordOrderType makes DnsRecordOrderType
func MakeDnsRecordOrderType() DnsRecordOrderType {
	var data DnsRecordOrderType
	return data
}

// InterfaceToDnsRecordOrderType makes DnsRecordOrderType from interface
func InterfaceToDnsRecordOrderType(data interface{}) DnsRecordOrderType {
	return data.(DnsRecordOrderType)
}

// InterfaceToDnsRecordOrderTypeSlice makes a slice of DnsRecordOrderType from interface
func InterfaceToDnsRecordOrderTypeSlice(data interface{}) []DnsRecordOrderType {
	list := data.([]interface{})
	result := MakeDnsRecordOrderTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDnsRecordOrderType(item))
	}
	return result
}

// MakeDnsRecordOrderTypeSlice() makes a slice of DnsRecordOrderType
func MakeDnsRecordOrderTypeSlice() []DnsRecordOrderType {
	return []DnsRecordOrderType{}
}
