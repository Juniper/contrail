
package models
// LoadbalancerHealthmonitor



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancerHealthmonitor_uuid int = iota
    propLoadbalancerHealthmonitor_annotations int = iota
    propLoadbalancerHealthmonitor_loadbalancer_healthmonitor_properties int = iota
    propLoadbalancerHealthmonitor_perms2 int = iota
    propLoadbalancerHealthmonitor_fq_name int = iota
    propLoadbalancerHealthmonitor_id_perms int = iota
    propLoadbalancerHealthmonitor_display_name int = iota
    propLoadbalancerHealthmonitor_parent_uuid int = iota
    propLoadbalancerHealthmonitor_parent_type int = iota
)

// LoadbalancerHealthmonitor 
type LoadbalancerHealthmonitor struct {

    UUID string `json:"uuid,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    LoadbalancerHealthmonitorProperties *LoadbalancerHealthmonitorType `json:"loadbalancer_healthmonitor_properties,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *LoadbalancerHealthmonitor) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancerHealthmonitor makes LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitor() *LoadbalancerHealthmonitor{
    return &LoadbalancerHealthmonitor{
    //TODO(nati): Apply default
    LoadbalancerHealthmonitorProperties: MakeLoadbalancerHealthmonitorType(),
        Perms2: MakePermType2(),
        UUID: "",
        Annotations: MakeKeyValuePairs(),
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerHealthmonitorSlice makes a slice of LoadbalancerHealthmonitor
func MakeLoadbalancerHealthmonitorSlice() []*LoadbalancerHealthmonitor {
    return []*LoadbalancerHealthmonitor{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerHealthmonitor) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc42024ad20])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *LoadbalancerHealthmonitor) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *LoadbalancerHealthmonitor) GetDefaultName() string {
    return strings.Replace("default-loadbalancer_healthmonitor", "_", "-", -1)
}

func (model *LoadbalancerHealthmonitor) GetType() string {
    return strings.Replace("loadbalancer_healthmonitor", "_", "-", -1)
}

func (model *LoadbalancerHealthmonitor) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerHealthmonitor) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerHealthmonitor) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerHealthmonitor) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerHealthmonitor) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerHealthmonitor) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerHealthmonitor) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerHealthmonitor) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerHealthmonitor) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancerHealthmonitor_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitor_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitor_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitor_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitor_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitor_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitor_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitor_loadbalancer_healthmonitor_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerHealthmonitorProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerHealthmonitorProperties as loadbalancer_healthmonitor_properties")
        }
        msg["loadbalancer_healthmonitor_properties"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerHealthmonitor_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LoadbalancerHealthmonitor) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerHealthmonitor) UpdateReferences() error {
    return nil
}


