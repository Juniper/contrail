package models

// CustomerAttachment

import "encoding/json"

// CustomerAttachment
type CustomerAttachment struct {
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`

	VirtualMachineInterfaceRefs []*CustomerAttachmentVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	FloatingIPRefs              []*CustomerAttachmentFloatingIPRef              `json:"floating_ip_refs,omitempty"`
}

// CustomerAttachmentVirtualMachineInterfaceRef references each other
type CustomerAttachmentVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// CustomerAttachmentFloatingIPRef references each other
type CustomerAttachmentFloatingIPRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *CustomerAttachment) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeCustomerAttachment makes CustomerAttachment
func MakeCustomerAttachment() *CustomerAttachment {
	return &CustomerAttachment{
		//TODO(nati): Apply default
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
	}
}

// MakeCustomerAttachmentSlice() makes a slice of CustomerAttachment
func MakeCustomerAttachmentSlice() []*CustomerAttachment {
	return []*CustomerAttachment{}
}
