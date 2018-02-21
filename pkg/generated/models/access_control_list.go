
package models
// AccessControlList



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAccessControlList_access_control_list_entries int = iota
    propAccessControlList_parent_type int = iota
    propAccessControlList_fq_name int = iota
    propAccessControlList_display_name int = iota
    propAccessControlList_id_perms int = iota
    propAccessControlList_access_control_list_hash int = iota
    propAccessControlList_annotations int = iota
    propAccessControlList_perms2 int = iota
    propAccessControlList_uuid int = iota
    propAccessControlList_parent_uuid int = iota
)

// AccessControlList 
type AccessControlList struct {

    ParentUUID string `json:"parent_uuid,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    AccessControlListHash map[string]interface{} `json:"access_control_list_hash,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    AccessControlListEntries *AclEntriesType `json:"access_control_list_entries,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AccessControlList) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAccessControlList makes AccessControlList
func MakeAccessControlList() *AccessControlList{
    return &AccessControlList{
    //TODO(nati): Apply default
    DisplayName: "",
        AccessControlListEntries: MakeAclEntriesType(),
        ParentType: "",
        FQName: []string{},
        UUID: "",
        ParentUUID: "",
        IDPerms: MakeIdPermsType(),
        AccessControlListHash: map[string]interface{}{},
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeAccessControlListSlice makes a slice of AccessControlList
func MakeAccessControlListSlice() []*AccessControlList {
    return []*AccessControlList{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AccessControlList) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[virtual_network:0xc420182820 security_group:0xc420182780])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *AccessControlList) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AccessControlList) GetDefaultName() string {
    return strings.Replace("default-access_control_list", "_", "-", -1)
}

func (model *AccessControlList) GetType() string {
    return strings.Replace("access_control_list", "_", "-", -1)
}

func (model *AccessControlList) GetFQName() []string {
    return model.FQName
}

func (model *AccessControlList) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AccessControlList) GetParentType() string {
    return model.ParentType
}

func (model *AccessControlList) GetUuid() string {
    return model.UUID
}

func (model *AccessControlList) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AccessControlList) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AccessControlList) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AccessControlList) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AccessControlList) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAccessControlList_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propAccessControlList_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propAccessControlList_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propAccessControlList_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propAccessControlList_access_control_list_hash) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AccessControlListHash); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AccessControlListHash as access_control_list_hash")
        }
        msg["access_control_list_hash"] = &val
    }
    
    if model.modified.Bit(propAccessControlList_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propAccessControlList_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propAccessControlList_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propAccessControlList_access_control_list_entries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AccessControlListEntries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AccessControlListEntries as access_control_list_entries")
        }
        msg["access_control_list_entries"] = &val
    }
    
    if model.modified.Bit(propAccessControlList_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AccessControlList) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AccessControlList) UpdateReferences() error {
    return nil
}


