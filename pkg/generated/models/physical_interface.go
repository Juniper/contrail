
package models
// PhysicalInterface



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPhysicalInterface_display_name int = iota
    propPhysicalInterface_perms2 int = iota
    propPhysicalInterface_uuid int = iota
    propPhysicalInterface_parent_uuid int = iota
    propPhysicalInterface_parent_type int = iota
    propPhysicalInterface_ethernet_segment_identifier int = iota
    propPhysicalInterface_id_perms int = iota
    propPhysicalInterface_annotations int = iota
    propPhysicalInterface_fq_name int = iota
)

// PhysicalInterface 
type PhysicalInterface struct {

    ParentType string `json:"parent_type,omitempty"`
    EthernetSegmentIdentifier string `json:"ethernet_segment_identifier,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`

    PhysicalInterfaceRefs []*PhysicalInterfacePhysicalInterfaceRef `json:"physical_interface_refs,omitempty"`

    LogicalInterfaces []*LogicalInterface `json:"logical_interfaces,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// PhysicalInterfacePhysicalInterfaceRef references each other
type PhysicalInterfacePhysicalInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *PhysicalInterface) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePhysicalInterface makes PhysicalInterface
func MakePhysicalInterface() *PhysicalInterface{
    return &PhysicalInterface{
    //TODO(nati): Apply default
    EthernetSegmentIdentifier: "",
        DisplayName: "",
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        
        modified: big.NewInt(0),
    }
}



// MakePhysicalInterfaceSlice makes a slice of PhysicalInterface
func MakePhysicalInterfaceSlice() []*PhysicalInterface {
    return []*PhysicalInterface{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PhysicalInterface) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[physical_router:0xc4202e40a0])
    fqn := []string{}
    
    fqn = PhysicalRouter{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *PhysicalInterface) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-physical_router", "_", "-", -1)
}

func (model *PhysicalInterface) GetDefaultName() string {
    return strings.Replace("default-physical_interface", "_", "-", -1)
}

func (model *PhysicalInterface) GetType() string {
    return strings.Replace("physical_interface", "_", "-", -1)
}

func (model *PhysicalInterface) GetFQName() []string {
    return model.FQName
}

func (model *PhysicalInterface) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PhysicalInterface) GetParentType() string {
    return model.ParentType
}

func (model *PhysicalInterface) GetUuid() string {
    return model.UUID
}

func (model *PhysicalInterface) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PhysicalInterface) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PhysicalInterface) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PhysicalInterface) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PhysicalInterface) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPhysicalInterface_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propPhysicalInterface_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propPhysicalInterface_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propPhysicalInterface_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propPhysicalInterface_ethernet_segment_identifier) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EthernetSegmentIdentifier); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EthernetSegmentIdentifier as ethernet_segment_identifier")
        }
        msg["ethernet_segment_identifier"] = &val
    }
    
    if model.modified.Bit(propPhysicalInterface_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propPhysicalInterface_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propPhysicalInterface_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propPhysicalInterface_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PhysicalInterface) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PhysicalInterface) UpdateReferences() error {
    return nil
}


