
package models
// Dashboard



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propDashboard_parent_uuid int = iota
    propDashboard_id_perms int = iota
    propDashboard_perms2 int = iota
    propDashboard_uuid int = iota
    propDashboard_display_name int = iota
    propDashboard_annotations int = iota
    propDashboard_container_config int = iota
    propDashboard_parent_type int = iota
    propDashboard_fq_name int = iota
)

// Dashboard 
type Dashboard struct {

    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ContainerConfig string `json:"container_config,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *Dashboard) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeDashboard makes Dashboard
func MakeDashboard() *Dashboard{
    return &Dashboard{
    //TODO(nati): Apply default
    ContainerConfig: "",
        ParentType: "",
        FQName: []string{},
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        IDPerms: MakeIdPermsType(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeDashboardSlice makes a slice of Dashboard
func MakeDashboardSlice() []*Dashboard {
    return []*Dashboard{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Dashboard) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *Dashboard) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *Dashboard) GetDefaultName() string {
    return strings.Replace("default-dashboard", "_", "-", -1)
}

func (model *Dashboard) GetType() string {
    return strings.Replace("dashboard", "_", "-", -1)
}

func (model *Dashboard) GetFQName() []string {
    return model.FQName
}

func (model *Dashboard) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Dashboard) GetParentType() string {
    return model.ParentType
}

func (model *Dashboard) GetUuid() string {
    return model.UUID
}

func (model *Dashboard) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Dashboard) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Dashboard) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Dashboard) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Dashboard) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propDashboard_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propDashboard_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propDashboard_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propDashboard_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propDashboard_container_config) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ContainerConfig); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ContainerConfig as container_config")
        }
        msg["container_config"] = &val
    }
    
    if model.modified.Bit(propDashboard_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propDashboard_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propDashboard_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propDashboard_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Dashboard) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Dashboard) UpdateReferences() error {
    return nil
}


