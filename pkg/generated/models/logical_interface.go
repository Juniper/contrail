package models
// LogicalInterface



import "encoding/json"

// LogicalInterface 
//proteus:generate
type LogicalInterface struct {

    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    LogicalInterfaceVlanTag int `json:"logical_interface_vlan_tag,omitempty"`
    LogicalInterfaceType LogicalInterfaceType `json:"logical_interface_type,omitempty"`

    VirtualMachineInterfaceRefs []*LogicalInterfaceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`

}


// LogicalInterfaceVirtualMachineInterfaceRef references each other
type LogicalInterfaceVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *LogicalInterface) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLogicalInterface makes LogicalInterface
func MakeLogicalInterface() *LogicalInterface{
    return &LogicalInterface{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        LogicalInterfaceVlanTag: 0,
        LogicalInterfaceType: MakeLogicalInterfaceType(),
        
    }
}



// MakeLogicalInterfaceSlice() makes a slice of LogicalInterface
func MakeLogicalInterfaceSlice() []*LogicalInterface {
    return []*LogicalInterface{}
}
