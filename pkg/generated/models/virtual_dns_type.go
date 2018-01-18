package models

// VirtualDnsType

import "encoding/json"

// VirtualDnsType
type VirtualDnsType struct {
	DynamicRecordsFromClient bool                  `json:"dynamic_records_from_client"`
	ReverseResolution        bool                  `json:"reverse_resolution"`
	DefaultTTLSeconds        int                   `json:"default_ttl_seconds,omitempty"`
	RecordOrder              DnsRecordOrderType    `json:"record_order,omitempty"`
	FloatingIPRecord         FloatingIpDnsNotation `json:"floating_ip_record,omitempty"`
	DomainName               string                `json:"domain_name,omitempty"`
	ExternalVisible          bool                  `json:"external_visible"`
	NextVirtualDNS           string                `json:"next_virtual_DNS,omitempty"`
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
		RecordOrder:              MakeDnsRecordOrderType(),
		FloatingIPRecord:         MakeFloatingIpDnsNotation(),
		DomainName:               "",
		ExternalVisible:          false,
		NextVirtualDNS:           "",
		DynamicRecordsFromClient: false,
		ReverseResolution:        false,
		DefaultTTLSeconds:        0,
	}
}

// MakeVirtualDnsTypeSlice() makes a slice of VirtualDnsType
func MakeVirtualDnsTypeSlice() []*VirtualDnsType {
	return []*VirtualDnsType{}
}
