
package models
// APIAccessList



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAPIAccessList_api_access_list_entries int = iota
    propAPIAccessList_uuid int = iota
    propAPIAccessList_display_name int = iota
    propAPIAccessList_annotations int = iota
    propAPIAccessList_perms2 int = iota
    propAPIAccessList_parent_uuid int = iota
    propAPIAccessList_parent_type int = iota
    propAPIAccessList_fq_name int = iota
    propAPIAccessList_id_perms int = iota
)

// APIAccessList 
type APIAccessList struct {

    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    UUID string `json:"uuid,omitempty"`
    APIAccessListEntries *RbacRuleEntriesType `json:"api_access_list_entries,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *APIAccessList) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAPIAccessList makes APIAccessList
func MakeAPIAccessList() *APIAccessList{
    return &APIAccessList{
    //TODO(nati): Apply default
    APIAccessListEntries: MakeRbacRuleEntriesType(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeAPIAccessListSlice makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
    return []*APIAccessList{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *APIAccessList) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc420182e60 global_system_config:0xc420182f00 domain:0xc420182fa0])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *APIAccessList) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *APIAccessList) GetDefaultName() string {
    return strings.Replace("default-api_access_list", "_", "-", -1)
}

func (model *APIAccessList) GetType() string {
    return strings.Replace("api_access_list", "_", "-", -1)
}

func (model *APIAccessList) GetFQName() []string {
    return model.FQName
}

func (model *APIAccessList) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *APIAccessList) GetParentType() string {
    return model.ParentType
}

func (model *APIAccessList) GetUuid() string {
    return model.UUID
}

func (model *APIAccessList) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *APIAccessList) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *APIAccessList) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *APIAccessList) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *APIAccessList) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAPIAccessList_api_access_list_entries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.APIAccessListEntries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: APIAccessListEntries as api_access_list_entries")
        }
        msg["api_access_list_entries"] = &val
    }
    
    if model.modified.Bit(propAPIAccessList_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propAPIAccessList_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propAPIAccessList_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propAPIAccessList_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propAPIAccessList_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propAPIAccessList_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propAPIAccessList_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propAPIAccessList_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *APIAccessList) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *APIAccessList) UpdateReferences() error {
    return nil
}


