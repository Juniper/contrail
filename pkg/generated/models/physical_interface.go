package models

// PhysicalInterface

import "encoding/json"

// PhysicalInterface
type PhysicalInterface struct {
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	EthernetSegmentIdentifier string         `json:"ethernet_segment_identifier,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`

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
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ParentUUID:                "",
		IDPerms:                   MakeIdPermsType(),
		EthernetSegmentIdentifier: "",
		ParentType:                "",
		FQName:                    []string{},
		Perms2:                    MakePermType2(),
	}
}

// MakePhysicalInterfaceSlice() makes a slice of PhysicalInterface
func MakePhysicalInterfaceSlice() []*PhysicalInterface {
	return []*PhysicalInterface{}
}
