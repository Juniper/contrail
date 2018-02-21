
package models
// UserDefinedLogStat



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propUserDefinedLogStat_pattern int = iota
    propUserDefinedLogStat_name int = iota
)

// UserDefinedLogStat 
type UserDefinedLogStat struct {

    Pattern string `json:"pattern,omitempty"`
    Name string `json:"name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *UserDefinedLogStat) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeUserDefinedLogStat makes UserDefinedLogStat
func MakeUserDefinedLogStat() *UserDefinedLogStat{
    return &UserDefinedLogStat{
    //TODO(nati): Apply default
    Pattern: "",
        Name: "",
        
        modified: big.NewInt(0),
    }
}



// MakeUserDefinedLogStatSlice makes a slice of UserDefinedLogStat
func MakeUserDefinedLogStatSlice() []*UserDefinedLogStat {
    return []*UserDefinedLogStat{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *UserDefinedLogStat) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *UserDefinedLogStat) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *UserDefinedLogStat) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *UserDefinedLogStat) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *UserDefinedLogStat) GetFQName() []string {
    return model.FQName
}

func (model *UserDefinedLogStat) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *UserDefinedLogStat) GetParentType() string {
    return model.ParentType
}

func (model *UserDefinedLogStat) GetUuid() string {
    return model.UUID
}

func (model *UserDefinedLogStat) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *UserDefinedLogStat) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *UserDefinedLogStat) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *UserDefinedLogStat) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *UserDefinedLogStat) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propUserDefinedLogStat_pattern) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Pattern); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Pattern as pattern")
        }
        msg["pattern"] = &val
    }
    
    if model.modified.Bit(propUserDefinedLogStat_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Name); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Name as name")
        }
        msg["name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *UserDefinedLogStat) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *UserDefinedLogStat) UpdateReferences() error {
    return nil
}


