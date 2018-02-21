
package models
// BridgeDomain



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propBridgeDomain_isid int = iota
    propBridgeDomain_perms2 int = iota
    propBridgeDomain_parent_type int = iota
    propBridgeDomain_display_name int = iota
    propBridgeDomain_annotations int = iota
    propBridgeDomain_uuid int = iota
    propBridgeDomain_parent_uuid int = iota
    propBridgeDomain_fq_name int = iota
    propBridgeDomain_mac_aging_time int = iota
    propBridgeDomain_mac_learning_enabled int = iota
    propBridgeDomain_mac_move_control int = iota
    propBridgeDomain_mac_limit_control int = iota
    propBridgeDomain_id_perms int = iota
)

// BridgeDomain 
type BridgeDomain struct {

    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    MacAgingTime MACAgingTime `json:"mac_aging_time,omitempty"`
    MacLearningEnabled bool `json:"mac_learning_enabled"`
    MacMoveControl *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
    MacLimitControl *MACLimitControlType `json:"mac_limit_control,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Isid IsidType `json:"isid,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    DisplayName string `json:"display_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *BridgeDomain) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeBridgeDomain makes BridgeDomain
func MakeBridgeDomain() *BridgeDomain{
    return &BridgeDomain{
    //TODO(nati): Apply default
    Isid: MakeIsidType(),
        Perms2: MakePermType2(),
        ParentType: "",
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        ParentUUID: "",
        FQName: []string{},
        MacAgingTime: MakeMACAgingTime(),
        MacLearningEnabled: false,
        MacMoveControl: MakeMACMoveLimitControlType(),
        MacLimitControl: MakeMACLimitControlType(),
        IDPerms: MakeIdPermsType(),
        
        modified: big.NewInt(0),
    }
}



// MakeBridgeDomainSlice makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
    return []*BridgeDomain{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *BridgeDomain) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[virtual_network:0xc420183540])
    fqn := []string{}
    
    fqn = VirtualNetwork{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *BridgeDomain) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-virtual_network", "_", "-", -1)
}

func (model *BridgeDomain) GetDefaultName() string {
    return strings.Replace("default-bridge_domain", "_", "-", -1)
}

func (model *BridgeDomain) GetType() string {
    return strings.Replace("bridge_domain", "_", "-", -1)
}

func (model *BridgeDomain) GetFQName() []string {
    return model.FQName
}

func (model *BridgeDomain) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *BridgeDomain) GetParentType() string {
    return model.ParentType
}

func (model *BridgeDomain) GetUuid() string {
    return model.UUID
}

func (model *BridgeDomain) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *BridgeDomain) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *BridgeDomain) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *BridgeDomain) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *BridgeDomain) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propBridgeDomain_isid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Isid); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Isid as isid")
        }
        msg["isid"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_mac_aging_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacAgingTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacAgingTime as mac_aging_time")
        }
        msg["mac_aging_time"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_mac_learning_enabled) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacLearningEnabled); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacLearningEnabled as mac_learning_enabled")
        }
        msg["mac_learning_enabled"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_mac_move_control) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacMoveControl); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacMoveControl as mac_move_control")
        }
        msg["mac_move_control"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_mac_limit_control) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacLimitControl); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacLimitControl as mac_limit_control")
        }
        msg["mac_limit_control"] = &val
    }
    
    if model.modified.Bit(propBridgeDomain_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *BridgeDomain) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *BridgeDomain) UpdateReferences() error {
    return nil
}


