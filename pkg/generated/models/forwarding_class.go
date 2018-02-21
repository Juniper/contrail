
package models
// ForwardingClass



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propForwardingClass_uuid int = iota
    propForwardingClass_parent_uuid int = iota
    propForwardingClass_parent_type int = iota
    propForwardingClass_fq_name int = iota
    propForwardingClass_forwarding_class_dscp int = iota
    propForwardingClass_forwarding_class_mpls_exp int = iota
    propForwardingClass_forwarding_class_id int = iota
    propForwardingClass_perms2 int = iota
    propForwardingClass_forwarding_class_vlan_priority int = iota
    propForwardingClass_id_perms int = iota
    propForwardingClass_display_name int = iota
    propForwardingClass_annotations int = iota
)

// ForwardingClass 
type ForwardingClass struct {

    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ForwardingClassVlanPriority VlanPriorityType `json:"forwarding_class_vlan_priority,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ForwardingClassDSCP DscpValueType `json:"forwarding_class_dscp,omitempty"`
    ForwardingClassMPLSExp MplsExpType `json:"forwarding_class_mpls_exp,omitempty"`
    ForwardingClassID ForwardingClassId `json:"forwarding_class_id,omitempty"`

    QosQueueRefs []*ForwardingClassQosQueueRef `json:"qos_queue_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// ForwardingClassQosQueueRef references each other
type ForwardingClassQosQueueRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *ForwardingClass) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeForwardingClass makes ForwardingClass
func MakeForwardingClass() *ForwardingClass{
    return &ForwardingClass{
    //TODO(nati): Apply default
    ForwardingClassVlanPriority: MakeVlanPriorityType(),
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        ForwardingClassDSCP: MakeDscpValueType(),
        ForwardingClassMPLSExp: MakeMplsExpType(),
        ForwardingClassID: MakeForwardingClassId(),
        Perms2: MakePermType2(),
        UUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeForwardingClassSlice makes a slice of ForwardingClass
func MakeForwardingClassSlice() []*ForwardingClass {
    return []*ForwardingClass{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ForwardingClass) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_qos_config:0xc42024a5a0])
    fqn := []string{}
    
    fqn = GlobalQosConfig{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *ForwardingClass) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-global_qos_config", "_", "-", -1)
}

func (model *ForwardingClass) GetDefaultName() string {
    return strings.Replace("default-forwarding_class", "_", "-", -1)
}

func (model *ForwardingClass) GetType() string {
    return strings.Replace("forwarding_class", "_", "-", -1)
}

func (model *ForwardingClass) GetFQName() []string {
    return model.FQName
}

func (model *ForwardingClass) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ForwardingClass) GetParentType() string {
    return model.ParentType
}

func (model *ForwardingClass) GetUuid() string {
    return model.UUID
}

func (model *ForwardingClass) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ForwardingClass) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ForwardingClass) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ForwardingClass) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ForwardingClass) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propForwardingClass_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_forwarding_class_dscp) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ForwardingClassDSCP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ForwardingClassDSCP as forwarding_class_dscp")
        }
        msg["forwarding_class_dscp"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_forwarding_class_mpls_exp) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ForwardingClassMPLSExp); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ForwardingClassMPLSExp as forwarding_class_mpls_exp")
        }
        msg["forwarding_class_mpls_exp"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_forwarding_class_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ForwardingClassID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ForwardingClassID as forwarding_class_id")
        }
        msg["forwarding_class_id"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_forwarding_class_vlan_priority) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ForwardingClassVlanPriority); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ForwardingClassVlanPriority as forwarding_class_vlan_priority")
        }
        msg["forwarding_class_vlan_priority"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propForwardingClass_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ForwardingClass) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ForwardingClass) UpdateReferences() error {
    return nil
}


