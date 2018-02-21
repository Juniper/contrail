
package models
// AclRuleType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAclRuleType_rule_uuid int = iota
    propAclRuleType_match_condition int = iota
    propAclRuleType_direction int = iota
    propAclRuleType_action_list int = iota
)

// AclRuleType 
type AclRuleType struct {

    MatchCondition *MatchConditionType `json:"match_condition,omitempty"`
    Direction DirectionType `json:"direction,omitempty"`
    ActionList *ActionListType `json:"action_list,omitempty"`
    RuleUUID string `json:"rule_uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AclRuleType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAclRuleType makes AclRuleType
func MakeAclRuleType() *AclRuleType{
    return &AclRuleType{
    //TODO(nati): Apply default
    RuleUUID: "",
        MatchCondition: MakeMatchConditionType(),
        Direction: MakeDirectionType(),
        ActionList: MakeActionListType(),
        
        modified: big.NewInt(0),
    }
}



// MakeAclRuleTypeSlice makes a slice of AclRuleType
func MakeAclRuleTypeSlice() []*AclRuleType {
    return []*AclRuleType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AclRuleType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AclRuleType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AclRuleType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AclRuleType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AclRuleType) GetFQName() []string {
    return model.FQName
}

func (model *AclRuleType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AclRuleType) GetParentType() string {
    return model.ParentType
}

func (model *AclRuleType) GetUuid() string {
    return model.UUID
}

func (model *AclRuleType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AclRuleType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AclRuleType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AclRuleType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AclRuleType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAclRuleType_action_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ActionList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ActionList as action_list")
        }
        msg["action_list"] = &val
    }
    
    if model.modified.Bit(propAclRuleType_rule_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RuleUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RuleUUID as rule_uuid")
        }
        msg["rule_uuid"] = &val
    }
    
    if model.modified.Bit(propAclRuleType_match_condition) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MatchCondition); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MatchCondition as match_condition")
        }
        msg["match_condition"] = &val
    }
    
    if model.modified.Bit(propAclRuleType_direction) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Direction); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Direction as direction")
        }
        msg["direction"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AclRuleType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AclRuleType) UpdateReferences() error {
    return nil
}


