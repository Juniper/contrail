package models

// PhysicalInterface

import "encoding/json"

// PhysicalInterface
type PhysicalInterface struct {
	ParentType                string         `json:"parent_type"`
	FQName                    []string       `json:"fq_name"`
	DisplayName               string         `json:"display_name"`
	Perms2                    *PermType2     `json:"perms2"`
	UUID                      string         `json:"uuid"`
	ParentUUID                string         `json:"parent_uuid"`
	EthernetSegmentIdentifier string         `json:"ethernet_segment_identifier"`
	IDPerms                   *IdPermsType   `json:"id_perms"`
	Annotations               *KeyValuePairs `json:"annotations"`

	PhysicalInterfaceRefs []*PhysicalInterfacePhysicalInterfaceRef `json:"physical_interface_refs"`

	LogicalInterfaces []*LogicalInterface `json:"logical_interfaces"`
}

// PhysicalInterfacePhysicalInterfaceRef references each other
type PhysicalInterfacePhysicalInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *PhysicalInterface) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePhysicalInterface makes PhysicalInterface
func MakePhysicalInterface() *PhysicalInterface {
	return &PhysicalInterface{
		//TODO(nati): Apply default
		EthernetSegmentIdentifier: "",
		IDPerms:                   MakeIdPermsType(),
		Annotations:               MakeKeyValuePairs(),
		DisplayName:               "",
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
	}
}

// InterfaceToPhysicalInterface makes PhysicalInterface from interface
func InterfaceToPhysicalInterface(iData interface{}) *PhysicalInterface {
	data := iData.(map[string]interface{})
	return &PhysicalInterface{
		EthernetSegmentIdentifier: data["ethernet_segment_identifier"].(string),

		//{"description":"Ethernet Segment Id configured for the Physical Interface. In a multihomed environment, user should configure the peer Physical interface with the same ESI.","type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToPhysicalInterfaceSlice makes a slice of PhysicalInterface from interface
func InterfaceToPhysicalInterfaceSlice(data interface{}) []*PhysicalInterface {
	list := data.([]interface{})
	result := MakePhysicalInterfaceSlice()
	for _, item := range list {
		result = append(result, InterfaceToPhysicalInterface(item))
	}
	return result
}

// MakePhysicalInterfaceSlice() makes a slice of PhysicalInterface
func MakePhysicalInterfaceSlice() []*PhysicalInterface {
	return []*PhysicalInterface{}
}
