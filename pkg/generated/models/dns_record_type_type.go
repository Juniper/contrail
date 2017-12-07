package models

// DnsRecordTypeType

type DnsRecordTypeType string

// MakeDnsRecordTypeType makes DnsRecordTypeType
func MakeDnsRecordTypeType() DnsRecordTypeType {
	var data DnsRecordTypeType
	return data
}

// InterfaceToDnsRecordTypeType makes DnsRecordTypeType from interface
func InterfaceToDnsRecordTypeType(data interface{}) DnsRecordTypeType {
	return data.(DnsRecordTypeType)
}

// InterfaceToDnsRecordTypeTypeSlice makes a slice of DnsRecordTypeType from interface
func InterfaceToDnsRecordTypeTypeSlice(data interface{}) []DnsRecordTypeType {
	list := data.([]interface{})
	result := MakeDnsRecordTypeTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDnsRecordTypeType(item))
	}
	return result
}

// MakeDnsRecordTypeTypeSlice() makes a slice of DnsRecordTypeType
func MakeDnsRecordTypeTypeSlice() []DnsRecordTypeType {
	return []DnsRecordTypeType{}
}
