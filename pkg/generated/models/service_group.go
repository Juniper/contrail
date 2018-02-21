
package models
// ServiceGroup



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceGroup_uuid int = iota
    propServiceGroup_parent_uuid int = iota
    propServiceGroup_parent_type int = iota
    propServiceGroup_annotations int = iota
    propServiceGroup_id_perms int = iota
    propServiceGroup_display_name int = iota
    propServiceGroup_perms2 int = iota
    propServiceGroup_service_group_firewall_service_list int = iota
    propServiceGroup_fq_name int = iota
)

// ServiceGroup 
type ServiceGroup struct {

    ServiceGroupFirewallServiceList *FirewallServiceGroupType `json:"service_group_firewall_service_list,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceGroup) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceGroup makes ServiceGroup
func MakeServiceGroup() *ServiceGroup{
    return &ServiceGroup{
    //TODO(nati): Apply default
    ParentType: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        ParentUUID: "",
        ServiceGroupFirewallServiceList: MakeFirewallServiceGroupType(),
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeServiceGroupSlice makes a slice of ServiceGroup
func MakeServiceGroupSlice() []*ServiceGroup {
    return []*ServiceGroup{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceGroup) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4202e5680 policy_management:0xc4202e5720])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *ServiceGroup) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceGroup) GetDefaultName() string {
    return strings.Replace("default-service_group", "_", "-", -1)
}

func (model *ServiceGroup) GetType() string {
    return strings.Replace("service_group", "_", "-", -1)
}

func (model *ServiceGroup) GetFQName() []string {
    return model.FQName
}

func (model *ServiceGroup) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceGroup) GetParentType() string {
    return model.ParentType
}

func (model *ServiceGroup) GetUuid() string {
    return model.UUID
}

func (model *ServiceGroup) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceGroup) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceGroup) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceGroup) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceGroup) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceGroup_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propServiceGroup_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propServiceGroup_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propServiceGroup_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propServiceGroup_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propServiceGroup_service_group_firewall_service_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceGroupFirewallServiceList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceGroupFirewallServiceList as service_group_firewall_service_list")
        }
        msg["service_group_firewall_service_list"] = &val
    }
    
    if model.modified.Bit(propServiceGroup_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propServiceGroup_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propServiceGroup_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceGroup) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceGroup) UpdateReferences() error {
    return nil
}


