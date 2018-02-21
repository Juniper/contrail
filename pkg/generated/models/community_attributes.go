
package models
// CommunityAttributes



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propCommunityAttributes_community_attribute int = iota
)

// CommunityAttributes 
type CommunityAttributes struct {

    CommunityAttribute CommunityAttribute `json:"community_attribute,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *CommunityAttributes) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeCommunityAttributes makes CommunityAttributes
func MakeCommunityAttributes() *CommunityAttributes{
    return &CommunityAttributes{
    //TODO(nati): Apply default
    
            
                CommunityAttribute:  MakeCommunityAttribute(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeCommunityAttributesSlice makes a slice of CommunityAttributes
func MakeCommunityAttributesSlice() []*CommunityAttributes {
    return []*CommunityAttributes{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *CommunityAttributes) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *CommunityAttributes) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *CommunityAttributes) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *CommunityAttributes) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *CommunityAttributes) GetFQName() []string {
    return model.FQName
}

func (model *CommunityAttributes) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *CommunityAttributes) GetParentType() string {
    return model.ParentType
}

func (model *CommunityAttributes) GetUuid() string {
    return model.UUID
}

func (model *CommunityAttributes) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *CommunityAttributes) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *CommunityAttributes) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *CommunityAttributes) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *CommunityAttributes) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propCommunityAttributes_community_attribute) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.CommunityAttribute); err != nil {
            return nil, errors.Wrap(err, "Marshal of: CommunityAttribute as community_attribute")
        }
        msg["community_attribute"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *CommunityAttributes) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *CommunityAttributes) UpdateReferences() error {
    return nil
}


