
package models
// E2ServiceProvider



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propE2ServiceProvider_display_name int = iota
    propE2ServiceProvider_annotations int = iota
    propE2ServiceProvider_parent_uuid int = iota
    propE2ServiceProvider_id_perms int = iota
    propE2ServiceProvider_e2_service_provider_promiscuous int = iota
    propE2ServiceProvider_perms2 int = iota
    propE2ServiceProvider_uuid int = iota
    propE2ServiceProvider_parent_type int = iota
    propE2ServiceProvider_fq_name int = iota
)

// E2ServiceProvider 
type E2ServiceProvider struct {

    E2ServiceProviderPromiscuous bool `json:"e2_service_provider_promiscuous"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`

    PhysicalRouterRefs []*E2ServiceProviderPhysicalRouterRef `json:"physical_router_refs,omitempty"`
    PeeringPolicyRefs []*E2ServiceProviderPeeringPolicyRef `json:"peering_policy_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// E2ServiceProviderPhysicalRouterRef references each other
type E2ServiceProviderPhysicalRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// E2ServiceProviderPeeringPolicyRef references each other
type E2ServiceProviderPeeringPolicyRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *E2ServiceProvider) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeE2ServiceProvider makes E2ServiceProvider
func MakeE2ServiceProvider() *E2ServiceProvider{
    return &E2ServiceProvider{
    //TODO(nati): Apply default
    ParentUUID: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        ParentType: "",
        FQName: []string{},
        E2ServiceProviderPromiscuous: false,
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeE2ServiceProviderSlice makes a slice of E2ServiceProvider
func MakeE2ServiceProviderSlice() []*E2ServiceProvider {
    return []*E2ServiceProvider{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *E2ServiceProvider) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *E2ServiceProvider) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *E2ServiceProvider) GetDefaultName() string {
    return strings.Replace("default-e2_service_provider", "_", "-", -1)
}

func (model *E2ServiceProvider) GetType() string {
    return strings.Replace("e2_service_provider", "_", "-", -1)
}

func (model *E2ServiceProvider) GetFQName() []string {
    return model.FQName
}

func (model *E2ServiceProvider) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *E2ServiceProvider) GetParentType() string {
    return model.ParentType
}

func (model *E2ServiceProvider) GetUuid() string {
    return model.UUID
}

func (model *E2ServiceProvider) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *E2ServiceProvider) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *E2ServiceProvider) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *E2ServiceProvider) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *E2ServiceProvider) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propE2ServiceProvider_e2_service_provider_promiscuous) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.E2ServiceProviderPromiscuous); err != nil {
            return nil, errors.Wrap(err, "Marshal of: E2ServiceProviderPromiscuous as e2_service_provider_promiscuous")
        }
        msg["e2_service_provider_promiscuous"] = &val
    }
    
    if model.modified.Bit(propE2ServiceProvider_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propE2ServiceProvider_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propE2ServiceProvider_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propE2ServiceProvider_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propE2ServiceProvider_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propE2ServiceProvider_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propE2ServiceProvider_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propE2ServiceProvider_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *E2ServiceProvider) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *E2ServiceProvider) UpdateReferences() error {
    return nil
}


