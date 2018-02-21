
package models
// ActionListType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propActionListType_mirror_to int = iota
    propActionListType_simple_action int = iota
    propActionListType_apply_service int = iota
    propActionListType_gateway_name int = iota
    propActionListType_log int = iota
    propActionListType_alert int = iota
    propActionListType_qos_action int = iota
    propActionListType_assign_routing_instance int = iota
)

// ActionListType 
type ActionListType struct {

    QosAction string `json:"qos_action,omitempty"`
    AssignRoutingInstance string `json:"assign_routing_instance,omitempty"`
    MirrorTo *MirrorActionType `json:"mirror_to,omitempty"`
    SimpleAction SimpleActionType `json:"simple_action,omitempty"`
    ApplyService []string `json:"apply_service,omitempty"`
    GatewayName string `json:"gateway_name,omitempty"`
    Log bool `json:"log"`
    Alert bool `json:"alert"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ActionListType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeActionListType makes ActionListType
func MakeActionListType() *ActionListType{
    return &ActionListType{
    //TODO(nati): Apply default
    SimpleAction: MakeSimpleActionType(),
        ApplyService: []string{},
        GatewayName: "",
        Log: false,
        Alert: false,
        QosAction: "",
        AssignRoutingInstance: "",
        MirrorTo: MakeMirrorActionType(),
        
        modified: big.NewInt(0),
    }
}



// MakeActionListTypeSlice makes a slice of ActionListType
func MakeActionListTypeSlice() []*ActionListType {
    return []*ActionListType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ActionListType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ActionListType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ActionListType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ActionListType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ActionListType) GetFQName() []string {
    return model.FQName
}

func (model *ActionListType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ActionListType) GetParentType() string {
    return model.ParentType
}

func (model *ActionListType) GetUuid() string {
    return model.UUID
}

func (model *ActionListType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ActionListType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ActionListType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ActionListType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ActionListType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propActionListType_simple_action) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SimpleAction); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SimpleAction as simple_action")
        }
        msg["simple_action"] = &val
    }
    
    if model.modified.Bit(propActionListType_apply_service) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ApplyService); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ApplyService as apply_service")
        }
        msg["apply_service"] = &val
    }
    
    if model.modified.Bit(propActionListType_gateway_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GatewayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GatewayName as gateway_name")
        }
        msg["gateway_name"] = &val
    }
    
    if model.modified.Bit(propActionListType_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Log); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Log as log")
        }
        msg["log"] = &val
    }
    
    if model.modified.Bit(propActionListType_alert) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Alert); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Alert as alert")
        }
        msg["alert"] = &val
    }
    
    if model.modified.Bit(propActionListType_qos_action) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.QosAction); err != nil {
            return nil, errors.Wrap(err, "Marshal of: QosAction as qos_action")
        }
        msg["qos_action"] = &val
    }
    
    if model.modified.Bit(propActionListType_assign_routing_instance) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AssignRoutingInstance); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AssignRoutingInstance as assign_routing_instance")
        }
        msg["assign_routing_instance"] = &val
    }
    
    if model.modified.Bit(propActionListType_mirror_to) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MirrorTo); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MirrorTo as mirror_to")
        }
        msg["mirror_to"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ActionListType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ActionListType) UpdateReferences() error {
    return nil
}


