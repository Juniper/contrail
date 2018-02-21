
package models
// Widget



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propWidget_fq_name int = iota
    propWidget_content_config int = iota
    propWidget_parent_uuid int = iota
    propWidget_parent_type int = iota
    propWidget_id_perms int = iota
    propWidget_display_name int = iota
    propWidget_annotations int = iota
    propWidget_perms2 int = iota
    propWidget_uuid int = iota
    propWidget_container_config int = iota
    propWidget_layout_config int = iota
)

// Widget 
type Widget struct {

    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ContainerConfig string `json:"container_config,omitempty"`
    LayoutConfig string `json:"layout_config,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    ContentConfig string `json:"content_config,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *Widget) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeWidget makes Widget
func MakeWidget() *Widget{
    return &Widget{
    //TODO(nati): Apply default
    UUID: "",
        ContainerConfig: "",
        LayoutConfig: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ContentConfig: "",
        ParentUUID: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeWidgetSlice makes a slice of Widget
func MakeWidgetSlice() []*Widget {
    return []*Widget{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Widget) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *Widget) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *Widget) GetDefaultName() string {
    return strings.Replace("default-widget", "_", "-", -1)
}

func (model *Widget) GetType() string {
    return strings.Replace("widget", "_", "-", -1)
}

func (model *Widget) GetFQName() []string {
    return model.FQName
}

func (model *Widget) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Widget) GetParentType() string {
    return model.ParentType
}

func (model *Widget) GetUuid() string {
    return model.UUID
}

func (model *Widget) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Widget) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Widget) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Widget) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Widget) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propWidget_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propWidget_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propWidget_container_config) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ContainerConfig); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ContainerConfig as container_config")
        }
        msg["container_config"] = &val
    }
    
    if model.modified.Bit(propWidget_layout_config) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LayoutConfig); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LayoutConfig as layout_config")
        }
        msg["layout_config"] = &val
    }
    
    if model.modified.Bit(propWidget_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propWidget_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propWidget_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propWidget_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propWidget_content_config) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ContentConfig); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ContentConfig as content_config")
        }
        msg["content_config"] = &val
    }
    
    if model.modified.Bit(propWidget_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propWidget_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Widget) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Widget) UpdateReferences() error {
    return nil
}


