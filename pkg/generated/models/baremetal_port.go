
package models
// BaremetalPort



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propBaremetalPort_pxe_enabled int = iota
    propBaremetalPort_perms2 int = iota
    propBaremetalPort_parent_type int = iota
    propBaremetalPort_node int = iota
    propBaremetalPort_port_id int = iota
    propBaremetalPort_annotations int = iota
    propBaremetalPort_uuid int = iota
    propBaremetalPort_fq_name int = iota
    propBaremetalPort_switch_id int = iota
    propBaremetalPort_parent_uuid int = iota
    propBaremetalPort_mac_address int = iota
    propBaremetalPort_switch_info int = iota
    propBaremetalPort_id_perms int = iota
    propBaremetalPort_display_name int = iota
)

// BaremetalPort 
type BaremetalPort struct {

    SwitchID string `json:"switch_id,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    MacAddress string `json:"mac_address,omitempty"`
    SwitchInfo string `json:"switch_info,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    PxeEnabled bool `json:"pxe_enabled"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    Node string `json:"node,omitempty"`
    PortID string `json:"port_id,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *BaremetalPort) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeBaremetalPort makes BaremetalPort
func MakeBaremetalPort() *BaremetalPort{
    return &BaremetalPort{
    //TODO(nati): Apply default
    Node: "",
        PortID: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        FQName: []string{},
        SwitchID: "",
        ParentUUID: "",
        MacAddress: "",
        SwitchInfo: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        PxeEnabled: false,
        Perms2: MakePermType2(),
        ParentType: "",
        
        modified: big.NewInt(0),
    }
}



// MakeBaremetalPortSlice makes a slice of BaremetalPort
func MakeBaremetalPortSlice() []*BaremetalPort {
    return []*BaremetalPort{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *BaremetalPort) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *BaremetalPort) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *BaremetalPort) GetDefaultName() string {
    return strings.Replace("default-baremetal_port", "_", "-", -1)
}

func (model *BaremetalPort) GetType() string {
    return strings.Replace("baremetal_port", "_", "-", -1)
}

func (model *BaremetalPort) GetFQName() []string {
    return model.FQName
}

func (model *BaremetalPort) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *BaremetalPort) GetParentType() string {
    return model.ParentType
}

func (model *BaremetalPort) GetUuid() string {
    return model.UUID
}

func (model *BaremetalPort) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *BaremetalPort) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *BaremetalPort) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *BaremetalPort) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *BaremetalPort) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propBaremetalPort_pxe_enabled) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PxeEnabled); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PxeEnabled as pxe_enabled")
        }
        msg["pxe_enabled"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_node) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Node); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Node as node")
        }
        msg["node"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_port_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PortID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PortID as port_id")
        }
        msg["port_id"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_switch_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SwitchID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SwitchID as switch_id")
        }
        msg["switch_id"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_mac_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacAddress as mac_address")
        }
        msg["mac_address"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_switch_info) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SwitchInfo); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SwitchInfo as switch_info")
        }
        msg["switch_info"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propBaremetalPort_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *BaremetalPort) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *BaremetalPort) UpdateReferences() error {
    return nil
}


