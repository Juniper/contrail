
package models
// JunosServicePorts



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propJunosServicePorts_service_port int = iota
)

// JunosServicePorts 
type JunosServicePorts struct {

    ServicePort []string `json:"service_port,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *JunosServicePorts) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeJunosServicePorts makes JunosServicePorts
func MakeJunosServicePorts() *JunosServicePorts{
    return &JunosServicePorts{
    //TODO(nati): Apply default
    ServicePort: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeJunosServicePortsSlice makes a slice of JunosServicePorts
func MakeJunosServicePortsSlice() []*JunosServicePorts {
    return []*JunosServicePorts{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *JunosServicePorts) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *JunosServicePorts) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *JunosServicePorts) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *JunosServicePorts) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *JunosServicePorts) GetFQName() []string {
    return model.FQName
}

func (model *JunosServicePorts) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *JunosServicePorts) GetParentType() string {
    return model.ParentType
}

func (model *JunosServicePorts) GetUuid() string {
    return model.UUID
}

func (model *JunosServicePorts) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *JunosServicePorts) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *JunosServicePorts) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *JunosServicePorts) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *JunosServicePorts) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propJunosServicePorts_service_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServicePort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServicePort as service_port")
        }
        msg["service_port"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *JunosServicePorts) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *JunosServicePorts) UpdateReferences() error {
    return nil
}


