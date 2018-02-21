
package models
// FirewallServiceGroupType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFirewallServiceGroupType_firewall_service int = iota
)

// FirewallServiceGroupType 
type FirewallServiceGroupType struct {

    FirewallService []*FirewallServiceType `json:"firewall_service,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FirewallServiceGroupType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallServiceGroupType makes FirewallServiceGroupType
func MakeFirewallServiceGroupType() *FirewallServiceGroupType{
    return &FirewallServiceGroupType{
    //TODO(nati): Apply default
    
            
                FirewallService:  MakeFirewallServiceTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeFirewallServiceGroupTypeSlice makes a slice of FirewallServiceGroupType
func MakeFirewallServiceGroupTypeSlice() []*FirewallServiceGroupType {
    return []*FirewallServiceGroupType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FirewallServiceGroupType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FirewallServiceGroupType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallServiceGroupType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FirewallServiceGroupType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallServiceGroupType) GetFQName() []string {
    return model.FQName
}

func (model *FirewallServiceGroupType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FirewallServiceGroupType) GetParentType() string {
    return model.ParentType
}

func (model *FirewallServiceGroupType) GetUuid() string {
    return model.UUID
}

func (model *FirewallServiceGroupType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FirewallServiceGroupType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FirewallServiceGroupType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FirewallServiceGroupType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FirewallServiceGroupType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFirewallServiceGroupType_firewall_service) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FirewallService); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FirewallService as firewall_service")
        }
        msg["firewall_service"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FirewallServiceGroupType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FirewallServiceGroupType) UpdateReferences() error {
    return nil
}


