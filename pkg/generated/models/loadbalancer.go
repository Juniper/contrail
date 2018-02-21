
package models
// Loadbalancer



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancer_loadbalancer_provider int = iota
    propLoadbalancer_parent_uuid int = iota
    propLoadbalancer_parent_type int = iota
    propLoadbalancer_display_name int = iota
    propLoadbalancer_annotations int = iota
    propLoadbalancer_loadbalancer_properties int = iota
    propLoadbalancer_perms2 int = iota
    propLoadbalancer_uuid int = iota
    propLoadbalancer_fq_name int = iota
    propLoadbalancer_id_perms int = iota
)

// Loadbalancer 
type Loadbalancer struct {

    LoadbalancerProperties *LoadbalancerType `json:"loadbalancer_properties,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    LoadbalancerProvider string `json:"loadbalancer_provider,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`

    ServiceInstanceRefs []*LoadbalancerServiceInstanceRef `json:"service_instance_refs,omitempty"`
    ServiceApplianceSetRefs []*LoadbalancerServiceApplianceSetRef `json:"service_appliance_set_refs,omitempty"`
    VirtualMachineInterfaceRefs []*LoadbalancerVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// LoadbalancerServiceApplianceSetRef references each other
type LoadbalancerServiceApplianceSetRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LoadbalancerVirtualMachineInterfaceRef references each other
type LoadbalancerVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LoadbalancerServiceInstanceRef references each other
type LoadbalancerServiceInstanceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *Loadbalancer) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancer makes Loadbalancer
func MakeLoadbalancer() *Loadbalancer{
    return &Loadbalancer{
    //TODO(nati): Apply default
    LoadbalancerProvider: "",
        ParentUUID: "",
        ParentType: "",
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        LoadbalancerProperties: MakeLoadbalancerType(),
        Perms2: MakePermType2(),
        UUID: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerSlice makes a slice of Loadbalancer
func MakeLoadbalancerSlice() []*Loadbalancer {
    return []*Loadbalancer{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Loadbalancer) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc42024b540])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *Loadbalancer) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *Loadbalancer) GetDefaultName() string {
    return strings.Replace("default-loadbalancer", "_", "-", -1)
}

func (model *Loadbalancer) GetType() string {
    return strings.Replace("loadbalancer", "_", "-", -1)
}

func (model *Loadbalancer) GetFQName() []string {
    return model.FQName
}

func (model *Loadbalancer) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Loadbalancer) GetParentType() string {
    return model.ParentType
}

func (model *Loadbalancer) GetUuid() string {
    return model.UUID
}

func (model *Loadbalancer) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Loadbalancer) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Loadbalancer) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Loadbalancer) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Loadbalancer) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancer_loadbalancer_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerProperties as loadbalancer_properties")
        }
        msg["loadbalancer_properties"] = &val
    }
    
    if model.modified.Bit(propLoadbalancer_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propLoadbalancer_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancer_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancer_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propLoadbalancer_loadbalancer_provider) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerProvider); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerProvider as loadbalancer_provider")
        }
        msg["loadbalancer_provider"] = &val
    }
    
    if model.modified.Bit(propLoadbalancer_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancer_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propLoadbalancer_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancer_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Loadbalancer) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Loadbalancer) UpdateReferences() error {
    return nil
}


