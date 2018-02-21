
package models
// NetworkPolicy



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propNetworkPolicy_fq_name int = iota
    propNetworkPolicy_network_policy_entries int = iota
    propNetworkPolicy_display_name int = iota
    propNetworkPolicy_annotations int = iota
    propNetworkPolicy_parent_type int = iota
    propNetworkPolicy_perms2 int = iota
    propNetworkPolicy_uuid int = iota
    propNetworkPolicy_parent_uuid int = iota
    propNetworkPolicy_id_perms int = iota
)

// NetworkPolicy 
type NetworkPolicy struct {

    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    NetworkPolicyEntries *PolicyEntriesType `json:"network_policy_entries,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *NetworkPolicy) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeNetworkPolicy makes NetworkPolicy
func MakeNetworkPolicy() *NetworkPolicy{
    return &NetworkPolicy{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        IDPerms: MakeIdPermsType(),
        NetworkPolicyEntries: MakePolicyEntriesType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        ParentType: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeNetworkPolicySlice makes a slice of NetworkPolicy
func MakeNetworkPolicySlice() []*NetworkPolicy {
    return []*NetworkPolicy{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *NetworkPolicy) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc42024bf40])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *NetworkPolicy) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *NetworkPolicy) GetDefaultName() string {
    return strings.Replace("default-network_policy", "_", "-", -1)
}

func (model *NetworkPolicy) GetType() string {
    return strings.Replace("network_policy", "_", "-", -1)
}

func (model *NetworkPolicy) GetFQName() []string {
    return model.FQName
}

func (model *NetworkPolicy) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *NetworkPolicy) GetParentType() string {
    return model.ParentType
}

func (model *NetworkPolicy) GetUuid() string {
    return model.UUID
}

func (model *NetworkPolicy) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *NetworkPolicy) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *NetworkPolicy) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *NetworkPolicy) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *NetworkPolicy) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propNetworkPolicy_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propNetworkPolicy_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propNetworkPolicy_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propNetworkPolicy_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propNetworkPolicy_network_policy_entries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NetworkPolicyEntries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NetworkPolicyEntries as network_policy_entries")
        }
        msg["network_policy_entries"] = &val
    }
    
    if model.modified.Bit(propNetworkPolicy_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propNetworkPolicy_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propNetworkPolicy_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propNetworkPolicy_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *NetworkPolicy) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *NetworkPolicy) UpdateReferences() error {
    return nil
}


