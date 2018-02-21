
package models
// FloatingIP



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFloatingIP_floating_ip_address_family int = iota
    propFloatingIP_floating_ip_address int = iota
    propFloatingIP_parent_uuid int = iota
    propFloatingIP_parent_type int = iota
    propFloatingIP_id_perms int = iota
    propFloatingIP_floating_ip_is_virtual_ip int = iota
    propFloatingIP_floating_ip_port_mappings_enable int = iota
    propFloatingIP_floating_ip_traffic_direction int = iota
    propFloatingIP_perms2 int = iota
    propFloatingIP_floating_ip_fixed_ip_address int = iota
    propFloatingIP_annotations int = iota
    propFloatingIP_uuid int = iota
    propFloatingIP_fq_name int = iota
    propFloatingIP_floating_ip_port_mappings int = iota
    propFloatingIP_display_name int = iota
)

// FloatingIP 
type FloatingIP struct {

    FloatingIPPortMappings *PortMappings `json:"floating_ip_port_mappings,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    FloatingIPAddressFamily IpAddressFamilyType `json:"floating_ip_address_family,omitempty"`
    FloatingIPAddress IpAddressType `json:"floating_ip_address,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    FloatingIPIsVirtualIP bool `json:"floating_ip_is_virtual_ip"`
    FloatingIPPortMappingsEnable bool `json:"floating_ip_port_mappings_enable"`
    FloatingIPTrafficDirection TrafficDirectionType `json:"floating_ip_traffic_direction,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    FloatingIPFixedIPAddress IpAddressType `json:"floating_ip_fixed_ip_address,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`

    ProjectRefs []*FloatingIPProjectRef `json:"project_refs,omitempty"`
    VirtualMachineInterfaceRefs []*FloatingIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// FloatingIPVirtualMachineInterfaceRef references each other
type FloatingIPVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// FloatingIPProjectRef references each other
type FloatingIPProjectRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *FloatingIP) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFloatingIP makes FloatingIP
func MakeFloatingIP() *FloatingIP{
    return &FloatingIP{
    //TODO(nati): Apply default
    FloatingIPPortMappings: MakePortMappings(),
        DisplayName: "",
        ParentUUID: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        FloatingIPAddressFamily: MakeIpAddressFamilyType(),
        FloatingIPAddress: MakeIpAddressType(),
        FloatingIPTrafficDirection: MakeTrafficDirectionType(),
        Perms2: MakePermType2(),
        FloatingIPIsVirtualIP: false,
        FloatingIPPortMappingsEnable: false,
        UUID: "",
        FQName: []string{},
        FloatingIPFixedIPAddress: MakeIpAddressType(),
        Annotations: MakeKeyValuePairs(),
        
        modified: big.NewInt(0),
    }
}



// MakeFloatingIPSlice makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
    return []*FloatingIP{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FloatingIP) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[instance_ip:0xc42024a3c0 floating_ip_pool:0xc42024a460])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *FloatingIP) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FloatingIP) GetDefaultName() string {
    return strings.Replace("default-floating_ip", "_", "-", -1)
}

func (model *FloatingIP) GetType() string {
    return strings.Replace("floating_ip", "_", "-", -1)
}

func (model *FloatingIP) GetFQName() []string {
    return model.FQName
}

func (model *FloatingIP) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FloatingIP) GetParentType() string {
    return model.ParentType
}

func (model *FloatingIP) GetUuid() string {
    return model.UUID
}

func (model *FloatingIP) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FloatingIP) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FloatingIP) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FloatingIP) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FloatingIP) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFloatingIP_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_floating_ip_fixed_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPFixedIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPFixedIPAddress as floating_ip_fixed_ip_address")
        }
        msg["floating_ip_fixed_ip_address"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_floating_ip_port_mappings) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPPortMappings); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPPortMappings as floating_ip_port_mappings")
        }
        msg["floating_ip_port_mappings"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_floating_ip_address_family) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPAddressFamily); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPAddressFamily as floating_ip_address_family")
        }
        msg["floating_ip_address_family"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_floating_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPAddress as floating_ip_address")
        }
        msg["floating_ip_address"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_floating_ip_traffic_direction) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPTrafficDirection); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPTrafficDirection as floating_ip_traffic_direction")
        }
        msg["floating_ip_traffic_direction"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_floating_ip_is_virtual_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPIsVirtualIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPIsVirtualIP as floating_ip_is_virtual_ip")
        }
        msg["floating_ip_is_virtual_ip"] = &val
    }
    
    if model.modified.Bit(propFloatingIP_floating_ip_port_mappings_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPPortMappingsEnable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPPortMappingsEnable as floating_ip_port_mappings_enable")
        }
        msg["floating_ip_port_mappings_enable"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FloatingIP) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FloatingIP) UpdateReferences() error {
    return nil
}


