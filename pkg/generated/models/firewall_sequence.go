
package models
// FirewallSequence



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFirewallSequence_sequence int = iota
)

// FirewallSequence 
type FirewallSequence struct {

    Sequence string `json:"sequence,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FirewallSequence) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallSequence makes FirewallSequence
func MakeFirewallSequence() *FirewallSequence{
    return &FirewallSequence{
    //TODO(nati): Apply default
    Sequence: "",
        
        modified: big.NewInt(0),
    }
}



// MakeFirewallSequenceSlice makes a slice of FirewallSequence
func MakeFirewallSequenceSlice() []*FirewallSequence {
    return []*FirewallSequence{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FirewallSequence) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FirewallSequence) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallSequence) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FirewallSequence) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallSequence) GetFQName() []string {
    return model.FQName
}

func (model *FirewallSequence) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FirewallSequence) GetParentType() string {
    return model.ParentType
}

func (model *FirewallSequence) GetUuid() string {
    return model.UUID
}

func (model *FirewallSequence) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FirewallSequence) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FirewallSequence) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FirewallSequence) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FirewallSequence) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFirewallSequence_sequence) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Sequence); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Sequence as sequence")
        }
        msg["sequence"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FirewallSequence) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FirewallSequence) UpdateReferences() error {
    return nil
}


