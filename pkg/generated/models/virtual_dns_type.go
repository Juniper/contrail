package models

// VirtualDnsType

import "encoding/json"

// VirtualDnsType
type VirtualDnsType struct {
	FloatingIPRecord         FloatingIpDnsNotation `json:"floating_ip_record"`
	DomainName               string                `json:"domain_name"`
	ExternalVisible          bool                  `json:"external_visible"`
	NextVirtualDNS           string                `json:"next_virtual_DNS"`
	DynamicRecordsFromClient bool                  `json:"dynamic_records_from_client"`
	ReverseResolution        bool                  `json:"reverse_resolution"`
	DefaultTTLSeconds        int                   `json:"default_ttl_seconds"`
	RecordOrder              DnsRecordOrderType    `json:"record_order"`
}

//  parents relation object

// String returns json representation of the object
func (model *VirtualDnsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualDnsType makes VirtualDnsType
func MakeVirtualDnsType() *VirtualDnsType {
	return &VirtualDnsType{
		//TODO(nati): Apply default
		NextVirtualDNS:           "",
		DynamicRecordsFromClient: false,
		ReverseResolution:        false,
		DefaultTTLSeconds:        0,
		RecordOrder:              MakeDnsRecordOrderType(),
		FloatingIPRecord:         MakeFloatingIpDnsNotation(),
		DomainName:               "",
		ExternalVisible:          false,
	}
}

// InterfaceToVirtualDnsType makes VirtualDnsType from interface
func InterfaceToVirtualDnsType(iData interface{}) *VirtualDnsType {
	data := iData.(map[string]interface{})
	return &VirtualDnsType{
		FloatingIPRecord: InterfaceToFloatingIpDnsNotation(data["floating_ip_record"]),

		//{"Title":"","Description":"Decides how floating ip records are added","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dashed-ip","dashed-ip-tenant-name","vm-name","vm-name-tenant-name"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FloatingIpDnsNotation","CollectionType":"","Column":"","Item":null,"GoName":"FloatingIPRecord","GoType":"FloatingIpDnsNotation","GoPremitive":false}
		DomainName: data["domain_name"].(string),

		//{"Title":"","Description":"Default domain name for this virtual DNS server","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DomainName","GoType":"string","GoPremitive":true}
		ExternalVisible: data["external_visible"].(bool),

		//{"Title":"","Description":"Currently this option is not supported","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ExternalVisible","GoType":"bool","GoPremitive":true}
		NextVirtualDNS: data["next_virtual_DNS"].(string),

		//{"Title":"","Description":"Next virtual DNS server to lookup if record is not found. Default is proxy to infrastructure DNS","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NextVirtualDNS","GoType":"string","GoPremitive":true}
		DynamicRecordsFromClient: data["dynamic_records_from_client"].(bool),

		//{"Title":"","Description":"Allow automatic addition of records on VM launch, default is True","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DynamicRecordsFromClient","GoType":"bool","GoPremitive":true}
		ReverseResolution: data["reverse_resolution"].(bool),

		//{"Title":"","Description":"Allow reverse DNS resolution, ip to name mapping","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ReverseResolution","GoType":"bool","GoPremitive":true}
		DefaultTTLSeconds: data["default_ttl_seconds"].(int),

		//{"Title":"","Description":"Default Time To Live for DNS records","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DefaultTTLSeconds","GoType":"int","GoPremitive":true}
		RecordOrder: InterfaceToDnsRecordOrderType(data["record_order"]),

		//{"Title":"","Description":"Order of DNS load balancing, fixed, random, round-robin. Default is random","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["fixed","random","round-robin"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/DnsRecordOrderType","CollectionType":"","Column":"","Item":null,"GoName":"RecordOrder","GoType":"DnsRecordOrderType","GoPremitive":false}

	}
}

// InterfaceToVirtualDnsTypeSlice makes a slice of VirtualDnsType from interface
func InterfaceToVirtualDnsTypeSlice(data interface{}) []*VirtualDnsType {
	list := data.([]interface{})
	result := MakeVirtualDnsTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualDnsType(item))
	}
	return result
}

// MakeVirtualDnsTypeSlice() makes a slice of VirtualDnsType
func MakeVirtualDnsTypeSlice() []*VirtualDnsType {
	return []*VirtualDnsType{}
}
