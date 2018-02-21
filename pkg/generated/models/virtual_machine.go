
package models
// VirtualMachine



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualMachine_annotations int = iota
    propVirtualMachine_perms2 int = iota
    propVirtualMachine_uuid int = iota
    propVirtualMachine_parent_uuid int = iota
    propVirtualMachine_parent_type int = iota
    propVirtualMachine_fq_name int = iota
    propVirtualMachine_id_perms int = iota
    propVirtualMachine_display_name int = iota
)

// VirtualMachine 
type VirtualMachine struct {

    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`

    ServiceInstanceRefs []*VirtualMachineServiceInstanceRef `json:"service_instance_refs,omitempty"`

    VirtualMachineInterfaces []*VirtualMachineInterface `json:"virtual_machine_interfaces,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// VirtualMachineServiceInstanceRef references each other
type VirtualMachineServiceInstanceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *VirtualMachine) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualMachine makes VirtualMachine
func MakeVirtualMachine() *VirtualMachine{
    return &VirtualMachine{
    //TODO(nati): Apply default
    FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualMachineSlice makes a slice of VirtualMachine
func MakeVirtualMachineSlice() []*VirtualMachine {
    return []*VirtualMachine{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualMachine) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *VirtualMachine) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualMachine) GetDefaultName() string {
    return strings.Replace("default-virtual_machine", "_", "-", -1)
}

func (model *VirtualMachine) GetType() string {
    return strings.Replace("virtual_machine", "_", "-", -1)
}

func (model *VirtualMachine) GetFQName() []string {
    return model.FQName
}

func (model *VirtualMachine) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualMachine) GetParentType() string {
    return model.ParentType
}

func (model *VirtualMachine) GetUuid() string {
    return model.UUID
}

func (model *VirtualMachine) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualMachine) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualMachine) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualMachine) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualMachine) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualMachine_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propVirtualMachine_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propVirtualMachine_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propVirtualMachine_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualMachine_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualMachine_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propVirtualMachine_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propVirtualMachine_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualMachine) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualMachine) UpdateReferences() error {
    return nil
}


