package models

// VirtualDnsType

import "encoding/json"

// VirtualDnsType
type VirtualDnsType struct {
	RecordOrder              DnsRecordOrderType    `json:"record_order"`
	FloatingIPRecord         FloatingIpDnsNotation `json:"floating_ip_record"`
	DomainName               string                `json:"domain_name"`
	ExternalVisible          bool                  `json:"external_visible"`
	NextVirtualDNS           string                `json:"next_virtual_DNS"`
	DynamicRecordsFromClient bool                  `json:"dynamic_records_from_client"`
	ReverseResolution        bool                  `json:"reverse_resolution"`
	DefaultTTLSeconds        int                   `json:"default_ttl_seconds"`
}

// String returns json representation of the object
func (model *VirtualDnsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualDnsType makes VirtualDnsType
func MakeVirtualDnsType() *VirtualDnsType {
	return &VirtualDnsType{
		//TODO(nati): Apply default
		FloatingIPRecord:         MakeFloatingIpDnsNotation(),
		DomainName:               "",
		ExternalVisible:          false,
		NextVirtualDNS:           "",
		DynamicRecordsFromClient: false,
		ReverseResolution:        false,
		DefaultTTLSeconds:        0,
		RecordOrder:              MakeDnsRecordOrderType(),
	}
}

// InterfaceToVirtualDnsType makes VirtualDnsType from interface
func InterfaceToVirtualDnsType(iData interface{}) *VirtualDnsType {
	data := iData.(map[string]interface{})
	return &VirtualDnsType{
		DefaultTTLSeconds: data["default_ttl_seconds"].(int),

		//{"description":"Default Time To Live for DNS records","type":"integer"}
		RecordOrder: InterfaceToDnsRecordOrderType(data["record_order"]),

		//{"description":"Order of DNS load balancing, fixed, random, round-robin. Default is random","type":"string","enum":["fixed","random","round-robin"]}
		FloatingIPRecord: InterfaceToFloatingIpDnsNotation(data["floating_ip_record"]),

		//{"description":"Decides how floating ip records are added","type":"string","enum":["dashed-ip","dashed-ip-tenant-name","vm-name","vm-name-tenant-name"]}
		DomainName: data["domain_name"].(string),

		//{"description":"Default domain name for this virtual DNS server","type":"string"}
		ExternalVisible: data["external_visible"].(bool),

		//{"description":"Currently this option is not supported","type":"boolean"}
		NextVirtualDNS: data["next_virtual_DNS"].(string),

		//{"description":"Next virtual DNS server to lookup if record is not found. Default is proxy to infrastructure DNS","type":"string"}
		DynamicRecordsFromClient: data["dynamic_records_from_client"].(bool),

		//{"description":"Allow automatic addition of records on VM launch, default is True","type":"boolean"}
		ReverseResolution: data["reverse_resolution"].(bool),

		//{"description":"Allow reverse DNS resolution, ip to name mapping","type":"boolean"}

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
