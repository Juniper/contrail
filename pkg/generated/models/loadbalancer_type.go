
package models
// LoadbalancerType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancerType_admin_state int = iota
    propLoadbalancerType_vip_address int = iota
    propLoadbalancerType_vip_subnet_id int = iota
    propLoadbalancerType_operating_status int = iota
    propLoadbalancerType_status int = iota
    propLoadbalancerType_provisioning_status int = iota
)

// LoadbalancerType 
type LoadbalancerType struct {

    VipSubnetID UuidStringType `json:"vip_subnet_id,omitempty"`
    OperatingStatus string `json:"operating_status,omitempty"`
    Status string `json:"status,omitempty"`
    ProvisioningStatus string `json:"provisioning_status,omitempty"`
    AdminState bool `json:"admin_state"`
    VipAddress IpAddressType `json:"vip_address,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *LoadbalancerType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancerType makes LoadbalancerType
func MakeLoadbalancerType() *LoadbalancerType{
    return &LoadbalancerType{
    //TODO(nati): Apply default
    Status: "",
        ProvisioningStatus: "",
        AdminState: false,
        VipAddress: MakeIpAddressType(),
        VipSubnetID: MakeUuidStringType(),
        OperatingStatus: "",
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerTypeSlice makes a slice of LoadbalancerType
func MakeLoadbalancerTypeSlice() []*LoadbalancerType {
    return []*LoadbalancerType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LoadbalancerType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LoadbalancerType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerType) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerType) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerType) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancerType_status) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Status); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Status as status")
        }
        msg["status"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerType_provisioning_status) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStatus); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStatus as provisioning_status")
        }
        msg["provisioning_status"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerType_admin_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AdminState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AdminState as admin_state")
        }
        msg["admin_state"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerType_vip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VipAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VipAddress as vip_address")
        }
        msg["vip_address"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerType_vip_subnet_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VipSubnetID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VipSubnetID as vip_subnet_id")
        }
        msg["vip_subnet_id"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerType_operating_status) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OperatingStatus); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OperatingStatus as operating_status")
        }
        msg["operating_status"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LoadbalancerType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerType) UpdateReferences() error {
    return nil
}


