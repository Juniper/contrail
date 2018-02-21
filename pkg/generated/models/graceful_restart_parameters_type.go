
package models
// GracefulRestartParametersType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propGracefulRestartParametersType_restart_time int = iota
    propGracefulRestartParametersType_long_lived_restart_time int = iota
    propGracefulRestartParametersType_enable int = iota
    propGracefulRestartParametersType_end_of_rib_timeout int = iota
    propGracefulRestartParametersType_bgp_helper_enable int = iota
    propGracefulRestartParametersType_xmpp_helper_enable int = iota
)

// GracefulRestartParametersType 
type GracefulRestartParametersType struct {

    RestartTime GracefulRestartTimeType `json:"restart_time,omitempty"`
    LongLivedRestartTime LongLivedGracefulRestartTimeType `json:"long_lived_restart_time,omitempty"`
    Enable bool `json:"enable"`
    EndOfRibTimeout EndOfRibTimeType `json:"end_of_rib_timeout,omitempty"`
    BGPHelperEnable bool `json:"bgp_helper_enable"`
    XMPPHelperEnable bool `json:"xmpp_helper_enable"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *GracefulRestartParametersType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeGracefulRestartParametersType makes GracefulRestartParametersType
func MakeGracefulRestartParametersType() *GracefulRestartParametersType{
    return &GracefulRestartParametersType{
    //TODO(nati): Apply default
    BGPHelperEnable: false,
        XMPPHelperEnable: false,
        RestartTime: MakeGracefulRestartTimeType(),
        LongLivedRestartTime: MakeLongLivedGracefulRestartTimeType(),
        Enable: false,
        EndOfRibTimeout: MakeEndOfRibTimeType(),
        
        modified: big.NewInt(0),
    }
}



// MakeGracefulRestartParametersTypeSlice makes a slice of GracefulRestartParametersType
func MakeGracefulRestartParametersTypeSlice() []*GracefulRestartParametersType {
    return []*GracefulRestartParametersType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *GracefulRestartParametersType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *GracefulRestartParametersType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *GracefulRestartParametersType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *GracefulRestartParametersType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *GracefulRestartParametersType) GetFQName() []string {
    return model.FQName
}

func (model *GracefulRestartParametersType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *GracefulRestartParametersType) GetParentType() string {
    return model.ParentType
}

func (model *GracefulRestartParametersType) GetUuid() string {
    return model.UUID
}

func (model *GracefulRestartParametersType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *GracefulRestartParametersType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *GracefulRestartParametersType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *GracefulRestartParametersType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *GracefulRestartParametersType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propGracefulRestartParametersType_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Enable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Enable as enable")
        }
        msg["enable"] = &val
    }
    
    if model.modified.Bit(propGracefulRestartParametersType_end_of_rib_timeout) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EndOfRibTimeout); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EndOfRibTimeout as end_of_rib_timeout")
        }
        msg["end_of_rib_timeout"] = &val
    }
    
    if model.modified.Bit(propGracefulRestartParametersType_bgp_helper_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BGPHelperEnable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BGPHelperEnable as bgp_helper_enable")
        }
        msg["bgp_helper_enable"] = &val
    }
    
    if model.modified.Bit(propGracefulRestartParametersType_xmpp_helper_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.XMPPHelperEnable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: XMPPHelperEnable as xmpp_helper_enable")
        }
        msg["xmpp_helper_enable"] = &val
    }
    
    if model.modified.Bit(propGracefulRestartParametersType_restart_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RestartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RestartTime as restart_time")
        }
        msg["restart_time"] = &val
    }
    
    if model.modified.Bit(propGracefulRestartParametersType_long_lived_restart_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LongLivedRestartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LongLivedRestartTime as long_lived_restart_time")
        }
        msg["long_lived_restart_time"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *GracefulRestartParametersType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *GracefulRestartParametersType) UpdateReferences() error {
    return nil
}


