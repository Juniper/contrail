
package models
// LoadbalancerPool



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancerPool_uuid int = iota
    propLoadbalancerPool_parent_uuid int = iota
    propLoadbalancerPool_parent_type int = iota
    propLoadbalancerPool_fq_name int = iota
    propLoadbalancerPool_loadbalancer_pool_custom_attributes int = iota
    propLoadbalancerPool_loadbalancer_pool_provider int = iota
    propLoadbalancerPool_annotations int = iota
    propLoadbalancerPool_perms2 int = iota
    propLoadbalancerPool_loadbalancer_pool_properties int = iota
    propLoadbalancerPool_id_perms int = iota
    propLoadbalancerPool_display_name int = iota
)

// LoadbalancerPool 
type LoadbalancerPool struct {

    FQName []string `json:"fq_name,omitempty"`
    LoadbalancerPoolCustomAttributes *KeyValuePairs `json:"loadbalancer_pool_custom_attributes,omitempty"`
    LoadbalancerPoolProvider string `json:"loadbalancer_pool_provider,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    LoadbalancerPoolProperties *LoadbalancerPoolType `json:"loadbalancer_pool_properties,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`

    ServiceApplianceSetRefs []*LoadbalancerPoolServiceApplianceSetRef `json:"service_appliance_set_refs,omitempty"`
    VirtualMachineInterfaceRefs []*LoadbalancerPoolVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
    LoadbalancerListenerRefs []*LoadbalancerPoolLoadbalancerListenerRef `json:"loadbalancer_listener_refs,omitempty"`
    ServiceInstanceRefs []*LoadbalancerPoolServiceInstanceRef `json:"service_instance_refs,omitempty"`
    LoadbalancerHealthmonitorRefs []*LoadbalancerPoolLoadbalancerHealthmonitorRef `json:"loadbalancer_healthmonitor_refs,omitempty"`

    LoadbalancerMembers []*LoadbalancerMember `json:"loadbalancer_members,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// LoadbalancerPoolLoadbalancerListenerRef references each other
type LoadbalancerPoolLoadbalancerListenerRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LoadbalancerPoolServiceInstanceRef references each other
type LoadbalancerPoolServiceInstanceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LoadbalancerPoolLoadbalancerHealthmonitorRef references each other
type LoadbalancerPoolLoadbalancerHealthmonitorRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LoadbalancerPoolServiceApplianceSetRef references each other
type LoadbalancerPoolServiceApplianceSetRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LoadbalancerPoolVirtualMachineInterfaceRef references each other
type LoadbalancerPoolVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *LoadbalancerPool) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancerPool makes LoadbalancerPool
func MakeLoadbalancerPool() *LoadbalancerPool{
    return &LoadbalancerPool{
    //TODO(nati): Apply default
    LoadbalancerPoolProperties: MakeLoadbalancerPoolType(),
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        FQName: []string{},
        LoadbalancerPoolCustomAttributes: MakeKeyValuePairs(),
        LoadbalancerPoolProvider: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerPoolSlice makes a slice of LoadbalancerPool
func MakeLoadbalancerPoolSlice() []*LoadbalancerPool {
    return []*LoadbalancerPool{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerPool) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc42024b2c0])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *LoadbalancerPool) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *LoadbalancerPool) GetDefaultName() string {
    return strings.Replace("default-loadbalancer_pool", "_", "-", -1)
}

func (model *LoadbalancerPool) GetType() string {
    return strings.Replace("loadbalancer_pool", "_", "-", -1)
}

func (model *LoadbalancerPool) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerPool) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerPool) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerPool) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerPool) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerPool) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerPool) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerPool) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerPool) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancerPool_loadbalancer_pool_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerPoolProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerPoolProperties as loadbalancer_pool_properties")
        }
        msg["loadbalancer_pool_properties"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_loadbalancer_pool_custom_attributes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerPoolCustomAttributes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerPoolCustomAttributes as loadbalancer_pool_custom_attributes")
        }
        msg["loadbalancer_pool_custom_attributes"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_loadbalancer_pool_provider) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerPoolProvider); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerPoolProvider as loadbalancer_pool_provider")
        }
        msg["loadbalancer_pool_provider"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerPool_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LoadbalancerPool) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerPool) UpdateReferences() error {
    return nil
}


