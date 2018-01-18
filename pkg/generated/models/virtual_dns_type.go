package models

// VirtualDnsType

import "encoding/json"

// VirtualDnsType
type VirtualDnsType struct {
	FloatingIPRecord         FloatingIpDnsNotation `json:"floating_ip_record,omitempty"`
	DomainName               string                `json:"domain_name,omitempty"`
	ExternalVisible          bool                  `json:"external_visible"`
	NextVirtualDNS           string                `json:"next_virtual_DNS,omitempty"`
	DynamicRecordsFromClient bool                  `json:"dynamic_records_from_client"`
	ReverseResolution        bool                  `json:"reverse_resolution"`
	DefaultTTLSeconds        int                   `json:"default_ttl_seconds,omitempty"`
	RecordOrder              DnsRecordOrderType    `json:"record_order,omitempty"`
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
		DomainName:               "",
		ExternalVisible:          false,
		NextVirtualDNS:           "",
		DynamicRecordsFromClient: false,
		ReverseResolution:        false,
		DefaultTTLSeconds:        0,
		RecordOrder:              MakeDnsRecordOrderType(),
		FloatingIPRecord:         MakeFloatingIpDnsNotation(),
	}
}

// MakeVirtualDnsTypeSlice() makes a slice of VirtualDnsType
func MakeVirtualDnsTypeSlice() []*VirtualDnsType {
	return []*VirtualDnsType{}
}
