
package models
// ServiceAppliance



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceAppliance_service_appliance_properties int = iota
    propServiceAppliance_id_perms int = iota
    propServiceAppliance_display_name int = iota
    propServiceAppliance_perms2 int = iota
    propServiceAppliance_fq_name int = iota
    propServiceAppliance_service_appliance_ip_address int = iota
    propServiceAppliance_annotations int = iota
    propServiceAppliance_uuid int = iota
    propServiceAppliance_parent_uuid int = iota
    propServiceAppliance_parent_type int = iota
    propServiceAppliance_service_appliance_user_credentials int = iota
)

// ServiceAppliance 
type ServiceAppliance struct {

    ServiceApplianceUserCredentials *UserCredentials `json:"service_appliance_user_credentials,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ServiceApplianceIPAddress IpAddressType `json:"service_appliance_ip_address,omitempty"`
    ServiceApplianceProperties *KeyValuePairs `json:"service_appliance_properties,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    FQName []string `json:"fq_name,omitempty"`

    PhysicalInterfaceRefs []*ServiceAppliancePhysicalInterfaceRef `json:"physical_interface_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// ServiceAppliancePhysicalInterfaceRef references each other
type ServiceAppliancePhysicalInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *ServiceApplianceInterfaceType
    
}


// String returns json representation of the object
func (model *ServiceAppliance) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceAppliance makes ServiceAppliance
func MakeServiceAppliance() *ServiceAppliance{
    return &ServiceAppliance{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        ServiceApplianceUserCredentials: MakeUserCredentials(),
        Annotations: MakeKeyValuePairs(),
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Perms2: MakePermType2(),
        FQName: []string{},
        ServiceApplianceIPAddress: MakeIpAddressType(),
        ServiceApplianceProperties: MakeKeyValuePairs(),
        
        modified: big.NewInt(0),
    }
}



// MakeServiceApplianceSlice makes a slice of ServiceAppliance
func MakeServiceApplianceSlice() []*ServiceAppliance {
    return []*ServiceAppliance{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceAppliance) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[service_appliance_set:0xc4202e52c0])
    fqn := []string{}
    
    fqn = ServiceApplianceSet{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *ServiceAppliance) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-service_appliance_set", "_", "-", -1)
}

func (model *ServiceAppliance) GetDefaultName() string {
    return strings.Replace("default-service_appliance", "_", "-", -1)
}

func (model *ServiceAppliance) GetType() string {
    return strings.Replace("service_appliance", "_", "-", -1)
}

func (model *ServiceAppliance) GetFQName() []string {
    return model.FQName
}

func (model *ServiceAppliance) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceAppliance) GetParentType() string {
    return model.ParentType
}

func (model *ServiceAppliance) GetUuid() string {
    return model.UUID
}

func (model *ServiceAppliance) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceAppliance) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceAppliance) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceAppliance) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceAppliance) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceAppliance_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_service_appliance_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceApplianceIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceApplianceIPAddress as service_appliance_ip_address")
        }
        msg["service_appliance_ip_address"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_service_appliance_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceApplianceProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceApplianceProperties as service_appliance_properties")
        }
        msg["service_appliance_properties"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_service_appliance_user_credentials) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceApplianceUserCredentials); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceApplianceUserCredentials as service_appliance_user_credentials")
        }
        msg["service_appliance_user_credentials"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propServiceAppliance_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceAppliance) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceAppliance) UpdateReferences() error {
    return nil
}


