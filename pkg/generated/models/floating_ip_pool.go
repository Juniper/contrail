
package models
// FloatingIPPool



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFloatingIPPool_fq_name int = iota
    propFloatingIPPool_floating_ip_pool_subnets int = iota
    propFloatingIPPool_perms2 int = iota
    propFloatingIPPool_parent_uuid int = iota
    propFloatingIPPool_parent_type int = iota
    propFloatingIPPool_display_name int = iota
    propFloatingIPPool_annotations int = iota
    propFloatingIPPool_uuid int = iota
    propFloatingIPPool_id_perms int = iota
)

// FloatingIPPool 
type FloatingIPPool struct {

    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    FloatingIPPoolSubnets *FloatingIpPoolSubnetType `json:"floating_ip_pool_subnets,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`


    FloatingIPs []*FloatingIP `json:"floating_ips,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FloatingIPPool) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFloatingIPPool makes FloatingIPPool
func MakeFloatingIPPool() *FloatingIPPool{
    return &FloatingIPPool{
    //TODO(nati): Apply default
    DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        IDPerms: MakeIdPermsType(),
        FloatingIPPoolSubnets: MakeFloatingIpPoolSubnetType(),
        Perms2: MakePermType2(),
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeFloatingIPPoolSlice makes a slice of FloatingIPPool
func MakeFloatingIPPoolSlice() []*FloatingIPPool {
    return []*FloatingIPPool{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FloatingIPPool) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[virtual_network:0xc42024a1e0])
    fqn := []string{}
    
    fqn = VirtualNetwork{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *FloatingIPPool) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-virtual_network", "_", "-", -1)
}

func (model *FloatingIPPool) GetDefaultName() string {
    return strings.Replace("default-floating_ip_pool", "_", "-", -1)
}

func (model *FloatingIPPool) GetType() string {
    return strings.Replace("floating_ip_pool", "_", "-", -1)
}

func (model *FloatingIPPool) GetFQName() []string {
    return model.FQName
}

func (model *FloatingIPPool) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FloatingIPPool) GetParentType() string {
    return model.ParentType
}

func (model *FloatingIPPool) GetUuid() string {
    return model.UUID
}

func (model *FloatingIPPool) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FloatingIPPool) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FloatingIPPool) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FloatingIPPool) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FloatingIPPool) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFloatingIPPool_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propFloatingIPPool_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propFloatingIPPool_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propFloatingIPPool_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propFloatingIPPool_floating_ip_pool_subnets) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPPoolSubnets); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPPoolSubnets as floating_ip_pool_subnets")
        }
        msg["floating_ip_pool_subnets"] = &val
    }
    
    if model.modified.Bit(propFloatingIPPool_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propFloatingIPPool_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propFloatingIPPool_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propFloatingIPPool_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FloatingIPPool) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FloatingIPPool) UpdateReferences() error {
    return nil
}


