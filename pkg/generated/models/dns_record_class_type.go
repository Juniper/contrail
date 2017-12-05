package models

// DnsRecordClassType

type DnsRecordClassType string

func MakeDnsRecordClassType() DnsRecordClassType {
	var data DnsRecordClassType
	return data
}

func InterfaceToDnsRecordClassType(data interface{}) DnsRecordClassType {
	return data.(DnsRecordClassType)
}

func InterfaceToDnsRecordClassTypeSlice(data interface{}) []DnsRecordClassType {
	list := data.([]interface{})
	result := MakeDnsRecordClassTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDnsRecordClassType(item))
	}
	return result
}

func MakeDnsRecordClassTypeSlice() []DnsRecordClassType {
	return []DnsRecordClassType{}
}
