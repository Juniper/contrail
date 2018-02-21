
package models
// NetworkIpam



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propNetworkIpam_uuid int = iota
    propNetworkIpam_parent_type int = iota
    propNetworkIpam_display_name int = iota
    propNetworkIpam_annotations int = iota
    propNetworkIpam_network_ipam_mgmt int = iota
    propNetworkIpam_ipam_subnets int = iota
    propNetworkIpam_ipam_subnet_method int = iota
    propNetworkIpam_perms2 int = iota
    propNetworkIpam_parent_uuid int = iota
    propNetworkIpam_fq_name int = iota
    propNetworkIpam_id_perms int = iota
)

// NetworkIpam 
type NetworkIpam struct {

    IpamSubnetMethod SubnetMethodType `json:"ipam_subnet_method,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    NetworkIpamMGMT *IpamType `json:"network_ipam_mgmt,omitempty"`
    IpamSubnets *IpamSubnets `json:"ipam_subnets,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`

    VirtualDNSRefs []*NetworkIpamVirtualDNSRef `json:"virtual_DNS_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// NetworkIpamVirtualDNSRef references each other
type NetworkIpamVirtualDNSRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *NetworkIpam) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeNetworkIpam makes NetworkIpam
func MakeNetworkIpam() *NetworkIpam{
    return &NetworkIpam{
    //TODO(nati): Apply default
    UUID: "",
        ParentType: "",
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        ParentUUID: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        NetworkIpamMGMT: MakeIpamType(),
        IpamSubnets: MakeIpamSubnets(),
        IpamSubnetMethod: MakeSubnetMethodType(),
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeNetworkIpamSlice makes a slice of NetworkIpam
func MakeNetworkIpamSlice() []*NetworkIpam {
    return []*NetworkIpam{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *NetworkIpam) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc42024bea0])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *NetworkIpam) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *NetworkIpam) GetDefaultName() string {
    return strings.Replace("default-network_ipam", "_", "-", -1)
}

func (model *NetworkIpam) GetType() string {
    return strings.Replace("network_ipam", "_", "-", -1)
}

func (model *NetworkIpam) GetFQName() []string {
    return model.FQName
}

func (model *NetworkIpam) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *NetworkIpam) GetParentType() string {
    return model.ParentType
}

func (model *NetworkIpam) GetUuid() string {
    return model.UUID
}

func (model *NetworkIpam) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *NetworkIpam) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *NetworkIpam) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *NetworkIpam) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *NetworkIpam) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propNetworkIpam_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_ipam_subnet_method) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IpamSubnetMethod); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IpamSubnetMethod as ipam_subnet_method")
        }
        msg["ipam_subnet_method"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_network_ipam_mgmt) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NetworkIpamMGMT); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NetworkIpamMGMT as network_ipam_mgmt")
        }
        msg["network_ipam_mgmt"] = &val
    }
    
    if model.modified.Bit(propNetworkIpam_ipam_subnets) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IpamSubnets); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IpamSubnets as ipam_subnets")
        }
        msg["ipam_subnets"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *NetworkIpam) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *NetworkIpam) UpdateReferences() error {
    return nil
}


