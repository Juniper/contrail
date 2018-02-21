
package models
// BridgeDomainMembershipType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propBridgeDomainMembershipType_vlan_tag int = iota
)

// BridgeDomainMembershipType 
type BridgeDomainMembershipType struct {

    VlanTag Dot1QTagType `json:"vlan_tag,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *BridgeDomainMembershipType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeBridgeDomainMembershipType makes BridgeDomainMembershipType
func MakeBridgeDomainMembershipType() *BridgeDomainMembershipType{
    return &BridgeDomainMembershipType{
    //TODO(nati): Apply default
    VlanTag: MakeDot1QTagType(),
        
        modified: big.NewInt(0),
    }
}



// MakeBridgeDomainMembershipTypeSlice makes a slice of BridgeDomainMembershipType
func MakeBridgeDomainMembershipTypeSlice() []*BridgeDomainMembershipType {
    return []*BridgeDomainMembershipType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *BridgeDomainMembershipType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *BridgeDomainMembershipType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *BridgeDomainMembershipType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *BridgeDomainMembershipType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *BridgeDomainMembershipType) GetFQName() []string {
    return model.FQName
}

func (model *BridgeDomainMembershipType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *BridgeDomainMembershipType) GetParentType() string {
    return model.ParentType
}

func (model *BridgeDomainMembershipType) GetUuid() string {
    return model.UUID
}

func (model *BridgeDomainMembershipType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *BridgeDomainMembershipType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *BridgeDomainMembershipType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *BridgeDomainMembershipType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *BridgeDomainMembershipType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propBridgeDomainMembershipType_vlan_tag) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VlanTag); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VlanTag as vlan_tag")
        }
        msg["vlan_tag"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *BridgeDomainMembershipType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *BridgeDomainMembershipType) UpdateReferences() error {
    return nil
}


