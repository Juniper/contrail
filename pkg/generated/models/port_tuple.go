
package models
// PortTuple



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPortTuple_fq_name int = iota
    propPortTuple_id_perms int = iota
    propPortTuple_display_name int = iota
    propPortTuple_annotations int = iota
    propPortTuple_perms2 int = iota
    propPortTuple_uuid int = iota
    propPortTuple_parent_uuid int = iota
    propPortTuple_parent_type int = iota
)

// PortTuple 
type PortTuple struct {

    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PortTuple) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePortTuple makes PortTuple
func MakePortTuple() *PortTuple{
    return &PortTuple{
    //TODO(nati): Apply default
    DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        
        modified: big.NewInt(0),
    }
}



// MakePortTupleSlice makes a slice of PortTuple
func MakePortTupleSlice() []*PortTuple {
    return []*PortTuple{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PortTuple) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[service_instance:0xc4202e4460])
    fqn := []string{}
    
    fqn = ServiceInstance{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *PortTuple) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-service_instance", "_", "-", -1)
}

func (model *PortTuple) GetDefaultName() string {
    return strings.Replace("default-port_tuple", "_", "-", -1)
}

func (model *PortTuple) GetType() string {
    return strings.Replace("port_tuple", "_", "-", -1)
}

func (model *PortTuple) GetFQName() []string {
    return model.FQName
}

func (model *PortTuple) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PortTuple) GetParentType() string {
    return model.ParentType
}

func (model *PortTuple) GetUuid() string {
    return model.UUID
}

func (model *PortTuple) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PortTuple) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PortTuple) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PortTuple) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PortTuple) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPortTuple_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propPortTuple_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propPortTuple_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propPortTuple_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propPortTuple_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propPortTuple_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propPortTuple_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propPortTuple_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PortTuple) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PortTuple) UpdateReferences() error {
    return nil
}


