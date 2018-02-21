
package models
// FirewallServiceType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFirewallServiceType_dst_ports int = iota
    propFirewallServiceType_src_ports int = iota
    propFirewallServiceType_protocol_id int = iota
    propFirewallServiceType_protocol int = iota
)

// FirewallServiceType 
type FirewallServiceType struct {

    Protocol string `json:"protocol,omitempty"`
    DSTPorts *PortType `json:"dst_ports,omitempty"`
    SRCPorts *PortType `json:"src_ports,omitempty"`
    ProtocolID int `json:"protocol_id,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FirewallServiceType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallServiceType makes FirewallServiceType
func MakeFirewallServiceType() *FirewallServiceType{
    return &FirewallServiceType{
    //TODO(nati): Apply default
    Protocol: "",
        DSTPorts: MakePortType(),
        SRCPorts: MakePortType(),
        ProtocolID: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeFirewallServiceTypeSlice makes a slice of FirewallServiceType
func MakeFirewallServiceTypeSlice() []*FirewallServiceType {
    return []*FirewallServiceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FirewallServiceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FirewallServiceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallServiceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FirewallServiceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallServiceType) GetFQName() []string {
    return model.FQName
}

func (model *FirewallServiceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FirewallServiceType) GetParentType() string {
    return model.ParentType
}

func (model *FirewallServiceType) GetUuid() string {
    return model.UUID
}

func (model *FirewallServiceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FirewallServiceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FirewallServiceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FirewallServiceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FirewallServiceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFirewallServiceType_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propFirewallServiceType_dst_ports) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DSTPorts); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DSTPorts as dst_ports")
        }
        msg["dst_ports"] = &val
    }
    
    if model.modified.Bit(propFirewallServiceType_src_ports) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SRCPorts); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SRCPorts as src_ports")
        }
        msg["src_ports"] = &val
    }
    
    if model.modified.Bit(propFirewallServiceType_protocol_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProtocolID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProtocolID as protocol_id")
        }
        msg["protocol_id"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FirewallServiceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FirewallServiceType) UpdateReferences() error {
    return nil
}


