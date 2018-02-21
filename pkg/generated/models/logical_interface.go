
package models
// LogicalInterface



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLogicalInterface_logical_interface_type int = iota
    propLogicalInterface_perms2 int = iota
    propLogicalInterface_uuid int = iota
    propLogicalInterface_parent_type int = iota
    propLogicalInterface_fq_name int = iota
    propLogicalInterface_id_perms int = iota
    propLogicalInterface_logical_interface_vlan_tag int = iota
    propLogicalInterface_display_name int = iota
    propLogicalInterface_annotations int = iota
    propLogicalInterface_parent_uuid int = iota
)

// LogicalInterface 
type LogicalInterface struct {

    LogicalInterfaceVlanTag int `json:"logical_interface_vlan_tag,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    LogicalInterfaceType LogicalInterfaceType `json:"logical_interface_type,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`

    VirtualMachineInterfaceRefs []*LogicalInterfaceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
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
    LogicalInterfaceType: MakeLogicalInterfaceType(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        LogicalInterfaceVlanTag: 0,
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeLogicalInterfaceSlice makes a slice of LogicalInterface
func MakeLogicalInterfaceSlice() []*LogicalInterface {
    return []*LogicalInterface{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LogicalInterface) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[physical_router:0xc42024b680 physical_interface:0xc42024b720])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *LogicalInterface) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LogicalInterface) GetDefaultName() string {
    return strings.Replace("default-logical_interface", "_", "-", -1)
}

func (model *LogicalInterface) GetType() string {
    return strings.Replace("logical_interface", "_", "-", -1)
}

func (model *LogicalInterface) GetFQName() []string {
    return model.FQName
}

func (model *LogicalInterface) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LogicalInterface) GetParentType() string {
    return model.ParentType
}

func (model *LogicalInterface) GetUuid() string {
    return model.UUID
}

func (model *LogicalInterface) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LogicalInterface) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LogicalInterface) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LogicalInterface) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LogicalInterface) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLogicalInterface_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propLogicalInterface_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propLogicalInterface_logical_interface_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LogicalInterfaceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LogicalInterfaceType as logical_interface_type")
        }
        msg["logical_interface_type"] = &val
    }
    
    if model.modified.Bit(propLogicalInterface_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propLogicalInterface_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propLogicalInterface_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propLogicalInterface_logical_interface_vlan_tag) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LogicalInterfaceVlanTag); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LogicalInterfaceVlanTag as logical_interface_vlan_tag")
        }
        msg["logical_interface_vlan_tag"] = &val
    }
    
    if model.modified.Bit(propLogicalInterface_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propLogicalInterface_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propLogicalInterface_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LogicalInterface) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LogicalInterface) UpdateReferences() error {
    return nil
}


