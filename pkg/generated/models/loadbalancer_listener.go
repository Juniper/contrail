
package models
// LoadbalancerListener



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancerListener_uuid int = iota
    propLoadbalancerListener_parent_uuid int = iota
    propLoadbalancerListener_parent_type int = iota
    propLoadbalancerListener_fq_name int = iota
    propLoadbalancerListener_loadbalancer_listener_properties int = iota
    propLoadbalancerListener_annotations int = iota
    propLoadbalancerListener_perms2 int = iota
    propLoadbalancerListener_id_perms int = iota
    propLoadbalancerListener_display_name int = iota
)

// LoadbalancerListener 
type LoadbalancerListener struct {

    Perms2 *PermType2 `json:"perms2,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    LoadbalancerListenerProperties *LoadbalancerListenerType `json:"loadbalancer_listener_properties,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`

    LoadbalancerRefs []*LoadbalancerListenerLoadbalancerRef `json:"loadbalancer_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// LoadbalancerListenerLoadbalancerRef references each other
type LoadbalancerListenerLoadbalancerRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *LoadbalancerListener) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancerListener makes LoadbalancerListener
func MakeLoadbalancerListener() *LoadbalancerListener{
    return &LoadbalancerListener{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        LoadbalancerListenerProperties: MakeLoadbalancerListenerType(),
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerListenerSlice makes a slice of LoadbalancerListener
func MakeLoadbalancerListenerSlice() []*LoadbalancerListener {
    return []*LoadbalancerListener{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerListener) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc42024ae60])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *LoadbalancerListener) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *LoadbalancerListener) GetDefaultName() string {
    return strings.Replace("default-loadbalancer_listener", "_", "-", -1)
}

func (model *LoadbalancerListener) GetType() string {
    return strings.Replace("loadbalancer_listener", "_", "-", -1)
}

func (model *LoadbalancerListener) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerListener) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerListener) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerListener) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerListener) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerListener) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerListener) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerListener) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerListener) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancerListener_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListener_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListener_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListener_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListener_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListener_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListener_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListener_loadbalancer_listener_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerListenerProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerListenerProperties as loadbalancer_listener_properties")
        }
        msg["loadbalancer_listener_properties"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListener_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LoadbalancerListener) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerListener) UpdateReferences() error {
    return nil
}


