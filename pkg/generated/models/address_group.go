
package models
// AddressGroup



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAddressGroup_perms2 int = iota
    propAddressGroup_parent_uuid int = iota
    propAddressGroup_parent_type int = iota
    propAddressGroup_id_perms int = iota
    propAddressGroup_display_name int = iota
    propAddressGroup_annotations int = iota
    propAddressGroup_uuid int = iota
    propAddressGroup_fq_name int = iota
    propAddressGroup_address_group_prefix int = iota
)

// AddressGroup 
type AddressGroup struct {

    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    AddressGroupPrefix *SubnetListType `json:"address_group_prefix,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AddressGroup) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAddressGroup makes AddressGroup
func MakeAddressGroup() *AddressGroup{
    return &AddressGroup{
    //TODO(nati): Apply default
    IDPerms: MakeIdPermsType(),
        Perms2: MakePermType2(),
        ParentUUID: "",
        ParentType: "",
        AddressGroupPrefix: MakeSubnetListType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeAddressGroupSlice makes a slice of AddressGroup
func MakeAddressGroupSlice() []*AddressGroup {
    return []*AddressGroup{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AddressGroup) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4201828c0 policy_management:0xc420182960])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *AddressGroup) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AddressGroup) GetDefaultName() string {
    return strings.Replace("default-address_group", "_", "-", -1)
}

func (model *AddressGroup) GetType() string {
    return strings.Replace("address_group", "_", "-", -1)
}

func (model *AddressGroup) GetFQName() []string {
    return model.FQName
}

func (model *AddressGroup) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AddressGroup) GetParentType() string {
    return model.ParentType
}

func (model *AddressGroup) GetUuid() string {
    return model.UUID
}

func (model *AddressGroup) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AddressGroup) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AddressGroup) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AddressGroup) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AddressGroup) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAddressGroup_address_group_prefix) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AddressGroupPrefix); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AddressGroupPrefix as address_group_prefix")
        }
        msg["address_group_prefix"] = &val
    }
    
    if model.modified.Bit(propAddressGroup_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propAddressGroup_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propAddressGroup_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propAddressGroup_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propAddressGroup_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propAddressGroup_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propAddressGroup_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propAddressGroup_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AddressGroup) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AddressGroup) UpdateReferences() error {
    return nil
}


