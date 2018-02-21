
package models
// TagType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propTagType_tag_type_id int = iota
    propTagType_parent_type int = iota
    propTagType_fq_name int = iota
    propTagType_annotations int = iota
    propTagType_uuid int = iota
    propTagType_parent_uuid int = iota
    propTagType_id_perms int = iota
    propTagType_display_name int = iota
    propTagType_perms2 int = iota
)

// TagType 
type TagType struct {

    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    TagTypeID U16BitHexInt `json:"tag_type_id,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *TagType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeTagType makes TagType
func MakeTagType() *TagType{
    return &TagType{
    //TODO(nati): Apply default
    Annotations: MakeKeyValuePairs(),
        UUID: "",
        ParentUUID: "",
        TagTypeID: MakeU16BitHexInt(),
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeTagTypeSlice makes a slice of TagType
func MakeTagTypeSlice() []*TagType {
    return []*TagType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *TagType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *TagType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *TagType) GetDefaultName() string {
    return strings.Replace("default-tag_type", "_", "-", -1)
}

func (model *TagType) GetType() string {
    return strings.Replace("tag_type", "_", "-", -1)
}

func (model *TagType) GetFQName() []string {
    return model.FQName
}

func (model *TagType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *TagType) GetParentType() string {
    return model.ParentType
}

func (model *TagType) GetUuid() string {
    return model.UUID
}

func (model *TagType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *TagType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *TagType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *TagType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *TagType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propTagType_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propTagType_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propTagType_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propTagType_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propTagType_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propTagType_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propTagType_tag_type_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TagTypeID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TagTypeID as tag_type_id")
        }
        msg["tag_type_id"] = &val
    }
    
    if model.modified.Bit(propTagType_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propTagType_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *TagType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *TagType) UpdateReferences() error {
    return nil
}


