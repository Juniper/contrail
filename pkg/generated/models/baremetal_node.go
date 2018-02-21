
package models
// BaremetalNode



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propBaremetalNode_cpu_count int = iota
    propBaremetalNode_memory_mb int = iota
    propBaremetalNode_deploy_kernel int = iota
    propBaremetalNode_deploy_ramdisk int = iota
    propBaremetalNode_fq_name int = iota
    propBaremetalNode_id_perms int = iota
    propBaremetalNode_display_name int = iota
    propBaremetalNode_annotations int = iota
    propBaremetalNode_ipmi_address int = iota
    propBaremetalNode_ipmi_username int = iota
    propBaremetalNode_uuid int = iota
    propBaremetalNode_perms2 int = iota
    propBaremetalNode_name int = iota
    propBaremetalNode_cpu_arch int = iota
    propBaremetalNode_parent_uuid int = iota
    propBaremetalNode_parent_type int = iota
    propBaremetalNode_ipmi_password int = iota
    propBaremetalNode_disk_gb int = iota
)

// BaremetalNode 
type BaremetalNode struct {

    Name string `json:"name,omitempty"`
    CPUArch string `json:"cpu_arch,omitempty"`
    UUID string `json:"uuid,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    IpmiPassword string `json:"ipmi_password,omitempty"`
    DiskGB int `json:"disk_gb,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    CPUCount int `json:"cpu_count,omitempty"`
    MemoryMB int `json:"memory_mb,omitempty"`
    IpmiAddress string `json:"ipmi_address,omitempty"`
    IpmiUsername string `json:"ipmi_username,omitempty"`
    DeployKernel string `json:"deploy_kernel,omitempty"`
    DeployRamdisk string `json:"deploy_ramdisk,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *BaremetalNode) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeBaremetalNode makes BaremetalNode
func MakeBaremetalNode() *BaremetalNode{
    return &BaremetalNode{
    //TODO(nati): Apply default
    Name: "",
        CPUArch: "",
        UUID: "",
        Perms2: MakePermType2(),
        IpmiPassword: "",
        DiskGB: 0,
        ParentUUID: "",
        ParentType: "",
        CPUCount: 0,
        MemoryMB: 0,
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        IpmiAddress: "",
        IpmiUsername: "",
        DeployKernel: "",
        DeployRamdisk: "",
        
        modified: big.NewInt(0),
    }
}



// MakeBaremetalNodeSlice makes a slice of BaremetalNode
func MakeBaremetalNodeSlice() []*BaremetalNode {
    return []*BaremetalNode{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *BaremetalNode) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *BaremetalNode) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *BaremetalNode) GetDefaultName() string {
    return strings.Replace("default-baremetal_node", "_", "-", -1)
}

func (model *BaremetalNode) GetType() string {
    return strings.Replace("baremetal_node", "_", "-", -1)
}

func (model *BaremetalNode) GetFQName() []string {
    return model.FQName
}

func (model *BaremetalNode) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *BaremetalNode) GetParentType() string {
    return model.ParentType
}

func (model *BaremetalNode) GetUuid() string {
    return model.UUID
}

func (model *BaremetalNode) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *BaremetalNode) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *BaremetalNode) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *BaremetalNode) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *BaremetalNode) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propBaremetalNode_ipmi_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IpmiAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IpmiAddress as ipmi_address")
        }
        msg["ipmi_address"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_ipmi_username) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IpmiUsername); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IpmiUsername as ipmi_username")
        }
        msg["ipmi_username"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_deploy_kernel) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DeployKernel); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DeployKernel as deploy_kernel")
        }
        msg["deploy_kernel"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_deploy_ramdisk) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DeployRamdisk); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DeployRamdisk as deploy_ramdisk")
        }
        msg["deploy_ramdisk"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Name); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Name as name")
        }
        msg["name"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_cpu_arch) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.CPUArch); err != nil {
            return nil, errors.Wrap(err, "Marshal of: CPUArch as cpu_arch")
        }
        msg["cpu_arch"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_ipmi_password) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IpmiPassword); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IpmiPassword as ipmi_password")
        }
        msg["ipmi_password"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_disk_gb) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DiskGB); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DiskGB as disk_gb")
        }
        msg["disk_gb"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_cpu_count) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.CPUCount); err != nil {
            return nil, errors.Wrap(err, "Marshal of: CPUCount as cpu_count")
        }
        msg["cpu_count"] = &val
    }
    
    if model.modified.Bit(propBaremetalNode_memory_mb) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MemoryMB); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MemoryMB as memory_mb")
        }
        msg["memory_mb"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *BaremetalNode) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *BaremetalNode) UpdateReferences() error {
    return nil
}


