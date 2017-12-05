package models

// DnsRecordTypeType

type DnsRecordTypeType string

func MakeDnsRecordTypeType() DnsRecordTypeType {
	var data DnsRecordTypeType
	return data
}

func InterfaceToDnsRecordTypeType(data interface{}) DnsRecordTypeType {
	return data.(DnsRecordTypeType)
}

func InterfaceToDnsRecordTypeTypeSlice(data interface{}) []DnsRecordTypeType {
	list := data.([]interface{})
	result := MakeDnsRecordTypeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDnsRecordTypeType(item))
	}
	return result
}

func MakeDnsRecordTypeTypeSlice() []DnsRecordTypeType {
	return []DnsRecordTypeType{}
}
