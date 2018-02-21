
package models
// LinklocalServiceEntryType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLinklocalServiceEntryType_linklocal_service_ip int = iota
    propLinklocalServiceEntryType_ip_fabric_service_port int = iota
    propLinklocalServiceEntryType_ip_fabric_DNS_service_name int = iota
    propLinklocalServiceEntryType_linklocal_service_port int = iota
    propLinklocalServiceEntryType_ip_fabric_service_ip int = iota
    propLinklocalServiceEntryType_linklocal_service_name int = iota
)

// LinklocalServiceEntryType 
type LinklocalServiceEntryType struct {

    IPFabricServiceIP []string `json:"ip_fabric_service_ip,omitempty"`
    LinklocalServiceName string `json:"linklocal_service_name,omitempty"`
    LinklocalServiceIP string `json:"linklocal_service_ip,omitempty"`
    IPFabricServicePort int `json:"ip_fabric_service_port,omitempty"`
    IPFabricDNSServiceName string `json:"ip_fabric_DNS_service_name,omitempty"`
    LinklocalServicePort int `json:"linklocal_service_port,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *LinklocalServiceEntryType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLinklocalServiceEntryType makes LinklocalServiceEntryType
func MakeLinklocalServiceEntryType() *LinklocalServiceEntryType{
    return &LinklocalServiceEntryType{
    //TODO(nati): Apply default
    IPFabricServiceIP: []string{},
        LinklocalServiceName: "",
        LinklocalServiceIP: "",
        IPFabricServicePort: 0,
        IPFabricDNSServiceName: "",
        LinklocalServicePort: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeLinklocalServiceEntryTypeSlice makes a slice of LinklocalServiceEntryType
func MakeLinklocalServiceEntryTypeSlice() []*LinklocalServiceEntryType {
    return []*LinklocalServiceEntryType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LinklocalServiceEntryType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LinklocalServiceEntryType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LinklocalServiceEntryType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LinklocalServiceEntryType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LinklocalServiceEntryType) GetFQName() []string {
    return model.FQName
}

func (model *LinklocalServiceEntryType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LinklocalServiceEntryType) GetParentType() string {
    return model.ParentType
}

func (model *LinklocalServiceEntryType) GetUuid() string {
    return model.UUID
}

func (model *LinklocalServiceEntryType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LinklocalServiceEntryType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LinklocalServiceEntryType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LinklocalServiceEntryType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LinklocalServiceEntryType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLinklocalServiceEntryType_ip_fabric_service_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPFabricServiceIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPFabricServiceIP as ip_fabric_service_ip")
        }
        msg["ip_fabric_service_ip"] = &val
    }
    
    if model.modified.Bit(propLinklocalServiceEntryType_linklocal_service_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LinklocalServiceName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LinklocalServiceName as linklocal_service_name")
        }
        msg["linklocal_service_name"] = &val
    }
    
    if model.modified.Bit(propLinklocalServiceEntryType_linklocal_service_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LinklocalServiceIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LinklocalServiceIP as linklocal_service_ip")
        }
        msg["linklocal_service_ip"] = &val
    }
    
    if model.modified.Bit(propLinklocalServiceEntryType_ip_fabric_service_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPFabricServicePort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPFabricServicePort as ip_fabric_service_port")
        }
        msg["ip_fabric_service_port"] = &val
    }
    
    if model.modified.Bit(propLinklocalServiceEntryType_ip_fabric_DNS_service_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPFabricDNSServiceName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPFabricDNSServiceName as ip_fabric_DNS_service_name")
        }
        msg["ip_fabric_DNS_service_name"] = &val
    }
    
    if model.modified.Bit(propLinklocalServiceEntryType_linklocal_service_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LinklocalServicePort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LinklocalServicePort as linklocal_service_port")
        }
        msg["linklocal_service_port"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LinklocalServiceEntryType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LinklocalServiceEntryType) UpdateReferences() error {
    return nil
}


