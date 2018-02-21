
package models
// Subnet



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propSubnet_subnet_ip_prefix int = iota
    propSubnet_parent_type int = iota
    propSubnet_display_name int = iota
    propSubnet_annotations int = iota
    propSubnet_perms2 int = iota
    propSubnet_fq_name int = iota
    propSubnet_id_perms int = iota
    propSubnet_uuid int = iota
    propSubnet_parent_uuid int = iota
)

// Subnet 
type Subnet struct {

    SubnetIPPrefix *SubnetType `json:"subnet_ip_prefix,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`

    VirtualMachineInterfaceRefs []*SubnetVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// SubnetVirtualMachineInterfaceRef references each other
type SubnetVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *Subnet) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSubnet makes Subnet
func MakeSubnet() *Subnet{
    return &Subnet{
    //TODO(nati): Apply default
    IDPerms: MakeIdPermsType(),
        UUID: "",
        ParentUUID: "",
        FQName: []string{},
        ParentType: "",
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        SubnetIPPrefix: MakeSubnetType(),
        
        modified: big.NewInt(0),
    }
}



// MakeSubnetSlice makes a slice of Subnet
func MakeSubnetSlice() []*Subnet {
    return []*Subnet{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Subnet) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *Subnet) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *Subnet) GetDefaultName() string {
    return strings.Replace("default-subnet", "_", "-", -1)
}

func (model *Subnet) GetType() string {
    return strings.Replace("subnet", "_", "-", -1)
}

func (model *Subnet) GetFQName() []string {
    return model.FQName
}

func (model *Subnet) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Subnet) GetParentType() string {
    return model.ParentType
}

func (model *Subnet) GetUuid() string {
    return model.UUID
}

func (model *Subnet) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Subnet) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Subnet) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Subnet) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Subnet) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propSubnet_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propSubnet_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propSubnet_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propSubnet_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propSubnet_subnet_ip_prefix) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SubnetIPPrefix); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SubnetIPPrefix as subnet_ip_prefix")
        }
        msg["subnet_ip_prefix"] = &val
    }
    
    if model.modified.Bit(propSubnet_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propSubnet_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propSubnet_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propSubnet_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Subnet) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Subnet) UpdateReferences() error {
    return nil
}


