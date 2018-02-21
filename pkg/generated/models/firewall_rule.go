
package models
// FirewallRule



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFirewallRule_fq_name int = iota
    propFirewallRule_parent_type int = iota
    propFirewallRule_match_tag_types int = iota
    propFirewallRule_match_tags int = iota
    propFirewallRule_display_name int = iota
    propFirewallRule_perms2 int = iota
    propFirewallRule_endpoint_2 int = iota
    propFirewallRule_id_perms int = iota
    propFirewallRule_parent_uuid int = iota
    propFirewallRule_service int = iota
    propFirewallRule_action_list int = iota
    propFirewallRule_direction int = iota
    propFirewallRule_annotations int = iota
    propFirewallRule_uuid int = iota
    propFirewallRule_endpoint_1 int = iota
)

// FirewallRule 
type FirewallRule struct {

    Service *FirewallServiceType `json:"service,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    Endpoint1 *FirewallRuleEndpointType `json:"endpoint_1,omitempty"`
    ActionList *ActionListType `json:"action_list,omitempty"`
    Direction FirewallRuleDirectionType `json:"direction,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    MatchTagTypes *FirewallRuleMatchTagsTypeIdList `json:"match_tag_types,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    Endpoint2 *FirewallRuleEndpointType `json:"endpoint_2,omitempty"`
    MatchTags *FirewallRuleMatchTagsType `json:"match_tags,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`

    ServiceGroupRefs []*FirewallRuleServiceGroupRef `json:"service_group_refs,omitempty"`
    AddressGroupRefs []*FirewallRuleAddressGroupRef `json:"address_group_refs,omitempty"`
    SecurityLoggingObjectRefs []*FirewallRuleSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
    VirtualNetworkRefs []*FirewallRuleVirtualNetworkRef `json:"virtual_network_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// FirewallRuleServiceGroupRef references each other
type FirewallRuleServiceGroupRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// FirewallRuleAddressGroupRef references each other
type FirewallRuleAddressGroupRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// FirewallRuleSecurityLoggingObjectRef references each other
type FirewallRuleSecurityLoggingObjectRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// FirewallRuleVirtualNetworkRef references each other
type FirewallRuleVirtualNetworkRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *FirewallRule) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallRule makes FirewallRule
func MakeFirewallRule() *FirewallRule{
    return &FirewallRule{
    //TODO(nati): Apply default
    Endpoint1: MakeFirewallRuleEndpointType(),
        ActionList: MakeActionListType(),
        Direction: MakeFirewallRuleDirectionType(),
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        MatchTagTypes: MakeFirewallRuleMatchTagsTypeIdList(),
        FQName: []string{},
        ParentType: "",
        Endpoint2: MakeFirewallRuleEndpointType(),
        MatchTags: MakeFirewallRuleMatchTagsType(),
        DisplayName: "",
        Perms2: MakePermType2(),
        Service: MakeFirewallServiceType(),
        IDPerms: MakeIdPermsType(),
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeFirewallRuleSlice makes a slice of FirewallRule
func MakeFirewallRuleSlice() []*FirewallRule {
    return []*FirewallRule{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FirewallRule) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc42024a0a0 policy_management:0xc42024a140])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *FirewallRule) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallRule) GetDefaultName() string {
    return strings.Replace("default-firewall_rule", "_", "-", -1)
}

func (model *FirewallRule) GetType() string {
    return strings.Replace("firewall_rule", "_", "-", -1)
}

func (model *FirewallRule) GetFQName() []string {
    return model.FQName
}

func (model *FirewallRule) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FirewallRule) GetParentType() string {
    return model.ParentType
}

func (model *FirewallRule) GetUuid() string {
    return model.UUID
}

func (model *FirewallRule) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FirewallRule) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FirewallRule) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FirewallRule) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FirewallRule) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFirewallRule_match_tag_types) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MatchTagTypes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MatchTagTypes as match_tag_types")
        }
        msg["match_tag_types"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_endpoint_2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Endpoint2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Endpoint2 as endpoint_2")
        }
        msg["endpoint_2"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_match_tags) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MatchTags); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MatchTags as match_tags")
        }
        msg["match_tags"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_service) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Service); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Service as service")
        }
        msg["service"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_endpoint_1) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Endpoint1); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Endpoint1 as endpoint_1")
        }
        msg["endpoint_1"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_action_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ActionList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ActionList as action_list")
        }
        msg["action_list"] = &val
    }
    
    if model.modified.Bit(propFirewallRule_direction) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Direction); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Direction as direction")
        }
        msg["direction"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FirewallRule) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FirewallRule) UpdateReferences() error {
    return nil
}


