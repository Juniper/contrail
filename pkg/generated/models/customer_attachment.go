
package models
// CustomerAttachment



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propCustomerAttachment_display_name int = iota
    propCustomerAttachment_annotations int = iota
    propCustomerAttachment_perms2 int = iota
    propCustomerAttachment_uuid int = iota
    propCustomerAttachment_parent_uuid int = iota
    propCustomerAttachment_parent_type int = iota
    propCustomerAttachment_fq_name int = iota
    propCustomerAttachment_id_perms int = iota
)

// CustomerAttachment 
type CustomerAttachment struct {

    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`

    VirtualMachineInterfaceRefs []*CustomerAttachmentVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
    FloatingIPRefs []*CustomerAttachmentFloatingIPRef `json:"floating_ip_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// CustomerAttachmentFloatingIPRef references each other
type CustomerAttachmentFloatingIPRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// CustomerAttachmentVirtualMachineInterfaceRef references each other
type CustomerAttachmentVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *CustomerAttachment) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeCustomerAttachment makes CustomerAttachment
func MakeCustomerAttachment() *CustomerAttachment{
    return &CustomerAttachment{
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



// MakeCustomerAttachmentSlice makes a slice of CustomerAttachment
func MakeCustomerAttachmentSlice() []*CustomerAttachment {
    return []*CustomerAttachment{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *CustomerAttachment) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *CustomerAttachment) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *CustomerAttachment) GetDefaultName() string {
    return strings.Replace("default-customer_attachment", "_", "-", -1)
}

func (model *CustomerAttachment) GetType() string {
    return strings.Replace("customer_attachment", "_", "-", -1)
}

func (model *CustomerAttachment) GetFQName() []string {
    return model.FQName
}

func (model *CustomerAttachment) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *CustomerAttachment) GetParentType() string {
    return model.ParentType
}

func (model *CustomerAttachment) GetUuid() string {
    return model.UUID
}

func (model *CustomerAttachment) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *CustomerAttachment) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *CustomerAttachment) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *CustomerAttachment) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *CustomerAttachment) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propCustomerAttachment_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propCustomerAttachment_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propCustomerAttachment_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propCustomerAttachment_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propCustomerAttachment_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propCustomerAttachment_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propCustomerAttachment_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propCustomerAttachment_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *CustomerAttachment) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *CustomerAttachment) UpdateReferences() error {
    return nil
}


