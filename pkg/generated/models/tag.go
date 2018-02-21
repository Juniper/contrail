
package models
// Tag



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propTag_tag_id int = iota
    propTag_tag_value int = iota
    propTag_display_name int = iota
    propTag_uuid int = iota
    propTag_id_perms int = iota
    propTag_tag_type_name int = iota
    propTag_annotations int = iota
    propTag_perms2 int = iota
    propTag_parent_uuid int = iota
    propTag_parent_type int = iota
    propTag_fq_name int = iota
)

// Tag 
type Tag struct {

    TagTypeName string `json:"tag_type_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    TagID U32BitHexInt `json:"tag_id,omitempty"`
    TagValue string `json:"tag_value,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    UUID string `json:"uuid,omitempty"`

    TagTypeRefs []*TagTagTypeRef `json:"tag_type_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// TagTagTypeRef references each other
type TagTagTypeRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *Tag) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeTag makes Tag
func MakeTag() *Tag{
    return &Tag{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        TagTypeName: "",
        Annotations: MakeKeyValuePairs(),
        DisplayName: "",
        UUID: "",
        TagID: MakeU32BitHexInt(),
        TagValue: "",
        
        modified: big.NewInt(0),
    }
}



// MakeTagSlice makes a slice of Tag
func MakeTagSlice() []*Tag {
    return []*Tag{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Tag) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4202e5d60 config_root:0xc4202e5e00])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *Tag) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *Tag) GetDefaultName() string {
    return strings.Replace("default-tag", "_", "-", -1)
}

func (model *Tag) GetType() string {
    return strings.Replace("tag", "_", "-", -1)
}

func (model *Tag) GetFQName() []string {
    return model.FQName
}

func (model *Tag) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Tag) GetParentType() string {
    return model.ParentType
}

func (model *Tag) GetUuid() string {
    return model.UUID
}

func (model *Tag) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Tag) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Tag) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Tag) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Tag) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propTag_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propTag_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propTag_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propTag_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propTag_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propTag_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propTag_tag_type_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TagTypeName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TagTypeName as tag_type_name")
        }
        msg["tag_type_name"] = &val
    }
    
    if model.modified.Bit(propTag_tag_value) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TagValue); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TagValue as tag_value")
        }
        msg["tag_value"] = &val
    }
    
    if model.modified.Bit(propTag_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propTag_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propTag_tag_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TagID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TagID as tag_id")
        }
        msg["tag_id"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Tag) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Tag) UpdateReferences() error {
    return nil
}


