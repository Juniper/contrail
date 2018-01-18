package models

// PhysicalInterface

import "encoding/json"

// PhysicalInterface
type PhysicalInterface struct {
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	EthernetSegmentIdentifier string         `json:"ethernet_segment_identifier,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`

	PhysicalInterfaceRefs []*PhysicalInterfacePhysicalInterfaceRef `json:"physical_interface_refs,omitempty"`

	LogicalInterfaces []*LogicalInterface `json:"logical_interfaces,omitempty"`
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
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
	}
}

// MakePhysicalInterfaceSlice() makes a slice of PhysicalInterface
func MakePhysicalInterfaceSlice() []*PhysicalInterface {
	return []*PhysicalInterface{}
}
