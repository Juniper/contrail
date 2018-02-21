
package models
// NetworkDeviceConfig



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propNetworkDeviceConfig_perms2 int = iota
    propNetworkDeviceConfig_uuid int = iota
    propNetworkDeviceConfig_parent_uuid int = iota
    propNetworkDeviceConfig_parent_type int = iota
    propNetworkDeviceConfig_fq_name int = iota
    propNetworkDeviceConfig_id_perms int = iota
    propNetworkDeviceConfig_display_name int = iota
    propNetworkDeviceConfig_annotations int = iota
)

// NetworkDeviceConfig 
type NetworkDeviceConfig struct {

    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`

    PhysicalRouterRefs []*NetworkDeviceConfigPhysicalRouterRef `json:"physical_router_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// NetworkDeviceConfigPhysicalRouterRef references each other
type NetworkDeviceConfigPhysicalRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *NetworkDeviceConfig) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeNetworkDeviceConfig makes NetworkDeviceConfig
func MakeNetworkDeviceConfig() *NetworkDeviceConfig{
    return &NetworkDeviceConfig{
    //TODO(nati): Apply default
    ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeNetworkDeviceConfigSlice makes a slice of NetworkDeviceConfig
func MakeNetworkDeviceConfigSlice() []*NetworkDeviceConfig {
    return []*NetworkDeviceConfig{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *NetworkDeviceConfig) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *NetworkDeviceConfig) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *NetworkDeviceConfig) GetDefaultName() string {
    return strings.Replace("default-network_device_config", "_", "-", -1)
}

func (model *NetworkDeviceConfig) GetType() string {
    return strings.Replace("network_device_config", "_", "-", -1)
}

func (model *NetworkDeviceConfig) GetFQName() []string {
    return model.FQName
}

func (model *NetworkDeviceConfig) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *NetworkDeviceConfig) GetParentType() string {
    return model.ParentType
}

func (model *NetworkDeviceConfig) GetUuid() string {
    return model.UUID
}

func (model *NetworkDeviceConfig) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *NetworkDeviceConfig) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *NetworkDeviceConfig) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *NetworkDeviceConfig) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *NetworkDeviceConfig) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propNetworkDeviceConfig_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propNetworkDeviceConfig_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propNetworkDeviceConfig_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propNetworkDeviceConfig_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propNetworkDeviceConfig_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propNetworkDeviceConfig_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propNetworkDeviceConfig_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propNetworkDeviceConfig_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *NetworkDeviceConfig) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *NetworkDeviceConfig) UpdateReferences() error {
    return nil
}


