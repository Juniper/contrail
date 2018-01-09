package models

// CustomerAttachment

import "encoding/json"

// CustomerAttachment
type CustomerAttachment struct {
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`

	VirtualMachineInterfaceRefs []*CustomerAttachmentVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
	FloatingIPRefs              []*CustomerAttachmentFloatingIPRef              `json:"floating_ip_refs"`
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
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
	}
}

// InterfaceToCustomerAttachment makes CustomerAttachment from interface
func InterfaceToCustomerAttachment(iData interface{}) *CustomerAttachment {
	data := iData.(map[string]interface{})
	return &CustomerAttachment{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}

	}
}

// InterfaceToCustomerAttachmentSlice makes a slice of CustomerAttachment from interface
func InterfaceToCustomerAttachmentSlice(data interface{}) []*CustomerAttachment {
	list := data.([]interface{})
	result := MakeCustomerAttachmentSlice()
	for _, item := range list {
		result = append(result, InterfaceToCustomerAttachment(item))
	}
	return result
}

// MakeCustomerAttachmentSlice() makes a slice of CustomerAttachment
func MakeCustomerAttachmentSlice() []*CustomerAttachment {
	return []*CustomerAttachment{}
}
