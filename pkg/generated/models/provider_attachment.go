
package models
// ProviderAttachment



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propProviderAttachment_fq_name int = iota
    propProviderAttachment_id_perms int = iota
    propProviderAttachment_display_name int = iota
    propProviderAttachment_annotations int = iota
    propProviderAttachment_perms2 int = iota
    propProviderAttachment_uuid int = iota
    propProviderAttachment_parent_uuid int = iota
    propProviderAttachment_parent_type int = iota
)

// ProviderAttachment 
type ProviderAttachment struct {

    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`

    VirtualRouterRefs []*ProviderAttachmentVirtualRouterRef `json:"virtual_router_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// ProviderAttachmentVirtualRouterRef references each other
type ProviderAttachmentVirtualRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *ProviderAttachment) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeProviderAttachment makes ProviderAttachment
func MakeProviderAttachment() *ProviderAttachment{
    return &ProviderAttachment{
    //TODO(nati): Apply default
    Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        
        modified: big.NewInt(0),
    }
}



// MakeProviderAttachmentSlice makes a slice of ProviderAttachment
func MakeProviderAttachmentSlice() []*ProviderAttachment {
    return []*ProviderAttachment{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ProviderAttachment) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *ProviderAttachment) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ProviderAttachment) GetDefaultName() string {
    return strings.Replace("default-provider_attachment", "_", "-", -1)
}

func (model *ProviderAttachment) GetType() string {
    return strings.Replace("provider_attachment", "_", "-", -1)
}

func (model *ProviderAttachment) GetFQName() []string {
    return model.FQName
}

func (model *ProviderAttachment) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ProviderAttachment) GetParentType() string {
    return model.ParentType
}

func (model *ProviderAttachment) GetUuid() string {
    return model.UUID
}

func (model *ProviderAttachment) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ProviderAttachment) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ProviderAttachment) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ProviderAttachment) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ProviderAttachment) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propProviderAttachment_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propProviderAttachment_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propProviderAttachment_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propProviderAttachment_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propProviderAttachment_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propProviderAttachment_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propProviderAttachment_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propProviderAttachment_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ProviderAttachment) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ProviderAttachment) UpdateReferences() error {
    return nil
}


