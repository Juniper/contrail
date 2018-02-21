
package models
// VPNGroup



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVPNGroup_id_perms int = iota
    propVPNGroup_annotations int = iota
    propVPNGroup_perms2 int = iota
    propVPNGroup_provisioning_progress int = iota
    propVPNGroup_display_name int = iota
    propVPNGroup_uuid int = iota
    propVPNGroup_parent_uuid int = iota
    propVPNGroup_parent_type int = iota
    propVPNGroup_provisioning_progress_stage int = iota
    propVPNGroup_provisioning_state int = iota
    propVPNGroup_type int = iota
    propVPNGroup_fq_name int = iota
    propVPNGroup_provisioning_log int = iota
    propVPNGroup_provisioning_start_time int = iota
)

// VPNGroup 
type VPNGroup struct {

    Type string `json:"type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`

    LocationRefs []*VPNGroupLocationRef `json:"location_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// VPNGroupLocationRef references each other
type VPNGroupLocationRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *VPNGroup) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVPNGroup makes VPNGroup
func MakeVPNGroup() *VPNGroup{
    return &VPNGroup{
    //TODO(nati): Apply default
    IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ProvisioningProgress: 0,
        DisplayName: "",
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        ProvisioningProgressStage: "",
        ProvisioningState: "",
        Type: "",
        FQName: []string{},
        ProvisioningLog: "",
        ProvisioningStartTime: "",
        
        modified: big.NewInt(0),
    }
}



// MakeVPNGroupSlice makes a slice of VPNGroup
func MakeVPNGroupSlice() []*VPNGroup {
    return []*VPNGroup{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VPNGroup) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *VPNGroup) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VPNGroup) GetDefaultName() string {
    return strings.Replace("default-vpn_group", "_", "-", -1)
}

func (model *VPNGroup) GetType() string {
    return strings.Replace("vpn_group", "_", "-", -1)
}

func (model *VPNGroup) GetFQName() []string {
    return model.FQName
}

func (model *VPNGroup) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VPNGroup) GetParentType() string {
    return model.ParentType
}

func (model *VPNGroup) GetUuid() string {
    return model.UUID
}

func (model *VPNGroup) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VPNGroup) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VPNGroup) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VPNGroup) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VPNGroup) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVPNGroup_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Type); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Type as type")
        }
        msg["type"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propVPNGroup_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VPNGroup) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VPNGroup) UpdateReferences() error {
    return nil
}


