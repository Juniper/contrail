
package models
// ApplicationPolicySet



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propApplicationPolicySet_all_applications int = iota
    propApplicationPolicySet_parent_type int = iota
    propApplicationPolicySet_id_perms int = iota
    propApplicationPolicySet_display_name int = iota
    propApplicationPolicySet_annotations int = iota
    propApplicationPolicySet_perms2 int = iota
    propApplicationPolicySet_uuid int = iota
    propApplicationPolicySet_parent_uuid int = iota
    propApplicationPolicySet_fq_name int = iota
)

// ApplicationPolicySet 
type ApplicationPolicySet struct {

    AllApplications bool `json:"all_applications"`
    ParentType string `json:"parent_type,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`

    FirewallPolicyRefs []*ApplicationPolicySetFirewallPolicyRef `json:"firewall_policy_refs,omitempty"`
    GlobalVrouterConfigRefs []*ApplicationPolicySetGlobalVrouterConfigRef `json:"global_vrouter_config_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// ApplicationPolicySetFirewallPolicyRef references each other
type ApplicationPolicySetFirewallPolicyRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *FirewallSequence
    
}

// ApplicationPolicySetGlobalVrouterConfigRef references each other
type ApplicationPolicySetGlobalVrouterConfigRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *ApplicationPolicySet) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeApplicationPolicySet makes ApplicationPolicySet
func MakeApplicationPolicySet() *ApplicationPolicySet{
    return &ApplicationPolicySet{
    //TODO(nati): Apply default
    AllApplications: false,
        ParentType: "",
        ParentUUID: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeApplicationPolicySetSlice makes a slice of ApplicationPolicySet
func MakeApplicationPolicySetSlice() []*ApplicationPolicySet {
    return []*ApplicationPolicySet{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ApplicationPolicySet) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc420183180 policy_management:0xc420183220])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *ApplicationPolicySet) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ApplicationPolicySet) GetDefaultName() string {
    return strings.Replace("default-application_policy_set", "_", "-", -1)
}

func (model *ApplicationPolicySet) GetType() string {
    return strings.Replace("application_policy_set", "_", "-", -1)
}

func (model *ApplicationPolicySet) GetFQName() []string {
    return model.FQName
}

func (model *ApplicationPolicySet) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ApplicationPolicySet) GetParentType() string {
    return model.ParentType
}

func (model *ApplicationPolicySet) GetUuid() string {
    return model.UUID
}

func (model *ApplicationPolicySet) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ApplicationPolicySet) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ApplicationPolicySet) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ApplicationPolicySet) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ApplicationPolicySet) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propApplicationPolicySet_all_applications) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AllApplications); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AllApplications as all_applications")
        }
        msg["all_applications"] = &val
    }
    
    if model.modified.Bit(propApplicationPolicySet_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propApplicationPolicySet_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propApplicationPolicySet_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propApplicationPolicySet_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propApplicationPolicySet_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propApplicationPolicySet_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propApplicationPolicySet_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propApplicationPolicySet_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ApplicationPolicySet) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ApplicationPolicySet) UpdateReferences() error {
    return nil
}


