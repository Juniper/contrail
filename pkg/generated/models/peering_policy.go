
package models
// PeeringPolicy



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPeeringPolicy_parent_uuid int = iota
    propPeeringPolicy_parent_type int = iota
    propPeeringPolicy_peering_service int = iota
    propPeeringPolicy_id_perms int = iota
    propPeeringPolicy_display_name int = iota
    propPeeringPolicy_annotations int = iota
    propPeeringPolicy_perms2 int = iota
    propPeeringPolicy_uuid int = iota
    propPeeringPolicy_fq_name int = iota
)

// PeeringPolicy 
type PeeringPolicy struct {

    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    PeeringService PeeringServiceType `json:"peering_service,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PeeringPolicy) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePeeringPolicy makes PeeringPolicy
func MakePeeringPolicy() *PeeringPolicy{
    return &PeeringPolicy{
    //TODO(nati): Apply default
    ParentUUID: "",
        ParentType: "",
        PeeringService: MakePeeringServiceType(),
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakePeeringPolicySlice makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
    return []*PeeringPolicy{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PeeringPolicy) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *PeeringPolicy) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PeeringPolicy) GetDefaultName() string {
    return strings.Replace("default-peering_policy", "_", "-", -1)
}

func (model *PeeringPolicy) GetType() string {
    return strings.Replace("peering_policy", "_", "-", -1)
}

func (model *PeeringPolicy) GetFQName() []string {
    return model.FQName
}

func (model *PeeringPolicy) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PeeringPolicy) GetParentType() string {
    return model.ParentType
}

func (model *PeeringPolicy) GetUuid() string {
    return model.UUID
}

func (model *PeeringPolicy) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PeeringPolicy) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PeeringPolicy) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PeeringPolicy) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PeeringPolicy) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPeeringPolicy_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propPeeringPolicy_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propPeeringPolicy_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propPeeringPolicy_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propPeeringPolicy_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propPeeringPolicy_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propPeeringPolicy_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propPeeringPolicy_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propPeeringPolicy_peering_service) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PeeringService); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PeeringService as peering_service")
        }
        msg["peering_service"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PeeringPolicy) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PeeringPolicy) UpdateReferences() error {
    return nil
}


