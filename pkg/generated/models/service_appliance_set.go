
package models
// ServiceApplianceSet



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceApplianceSet_service_appliance_ha_mode int = iota
    propServiceApplianceSet_service_appliance_driver int = iota
    propServiceApplianceSet_id_perms int = iota
    propServiceApplianceSet_annotations int = iota
    propServiceApplianceSet_uuid int = iota
    propServiceApplianceSet_parent_uuid int = iota
    propServiceApplianceSet_service_appliance_set_properties int = iota
    propServiceApplianceSet_parent_type int = iota
    propServiceApplianceSet_fq_name int = iota
    propServiceApplianceSet_display_name int = iota
    propServiceApplianceSet_perms2 int = iota
)

// ServiceApplianceSet 
type ServiceApplianceSet struct {

    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ServiceApplianceHaMode string `json:"service_appliance_ha_mode,omitempty"`
    ServiceApplianceDriver string `json:"service_appliance_driver,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ServiceApplianceSetProperties *KeyValuePairs `json:"service_appliance_set_properties,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`


    ServiceAppliances []*ServiceAppliance `json:"service_appliances,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceApplianceSet) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceApplianceSet makes ServiceApplianceSet
func MakeServiceApplianceSet() *ServiceApplianceSet{
    return &ServiceApplianceSet{
    //TODO(nati): Apply default
    ServiceApplianceSetProperties: MakeKeyValuePairs(),
        ParentType: "",
        FQName: []string{},
        DisplayName: "",
        Perms2: MakePermType2(),
        ServiceApplianceHaMode: "",
        ServiceApplianceDriver: "",
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeServiceApplianceSetSlice makes a slice of ServiceApplianceSet
func MakeServiceApplianceSetSlice() []*ServiceApplianceSet {
    return []*ServiceApplianceSet{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceApplianceSet) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_system_config:0xc4202e5360])
    fqn := []string{}
    
    fqn = GlobalSystemConfig{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *ServiceApplianceSet) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-global_system_config", "_", "-", -1)
}

func (model *ServiceApplianceSet) GetDefaultName() string {
    return strings.Replace("default-service_appliance_set", "_", "-", -1)
}

func (model *ServiceApplianceSet) GetType() string {
    return strings.Replace("service_appliance_set", "_", "-", -1)
}

func (model *ServiceApplianceSet) GetFQName() []string {
    return model.FQName
}

func (model *ServiceApplianceSet) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceApplianceSet) GetParentType() string {
    return model.ParentType
}

func (model *ServiceApplianceSet) GetUuid() string {
    return model.UUID
}

func (model *ServiceApplianceSet) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceApplianceSet) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceApplianceSet) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceApplianceSet) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceApplianceSet) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceApplianceSet_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_service_appliance_set_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceApplianceSetProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceApplianceSetProperties as service_appliance_set_properties")
        }
        msg["service_appliance_set_properties"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_service_appliance_ha_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceApplianceHaMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceApplianceHaMode as service_appliance_ha_mode")
        }
        msg["service_appliance_ha_mode"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_service_appliance_driver) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceApplianceDriver); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceApplianceDriver as service_appliance_driver")
        }
        msg["service_appliance_driver"] = &val
    }
    
    if model.modified.Bit(propServiceApplianceSet_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceApplianceSet) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceApplianceSet) UpdateReferences() error {
    return nil
}


