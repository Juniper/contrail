package models

// DnsRecordOrderType

type DnsRecordOrderType string

func MakeDnsRecordOrderType() DnsRecordOrderType {
	var data DnsRecordOrderType
	return data
}

func InterfaceToDnsRecordOrderType(data interface{}) DnsRecordOrderType {
	return data.(DnsRecordOrderType)
}

func InterfaceToDnsRecordOrderTypeSlice(data interface{}) []DnsRecordOrderType {
	list := data.([]interface{})
	result := MakeDnsRecordOrderTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDnsRecordOrderType(item))
	}
	return result
}

func MakeDnsRecordOrderTypeSlice() []DnsRecordOrderType {
	return []DnsRecordOrderType{}
}
