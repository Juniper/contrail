package models

// QosConfig

import "encoding/json"

// QosConfig
type QosConfig struct {
	MPLSExpEntries           *QosIdForwardingClassPairs `json:"mpls_exp_entries,omitempty"`
	DSCPEntries              *QosIdForwardingClassPairs `json:"dscp_entries,omitempty"`
	UUID                     string                     `json:"uuid,omitempty"`
	ParentUUID               string                     `json:"parent_uuid,omitempty"`
	ParentType               string                     `json:"parent_type,omitempty"`
	FQName                   []string                   `json:"fq_name,omitempty"`
	IDPerms                  *IdPermsType               `json:"id_perms,omitempty"`
	QosConfigType            QosConfigType              `json:"qos_config_type,omitempty"`
	VlanPriorityEntries      *QosIdForwardingClassPairs `json:"vlan_priority_entries,omitempty"`
	DefaultForwardingClassID ForwardingClassId          `json:"default_forwarding_class_id,omitempty"`
	Annotations              *KeyValuePairs             `json:"annotations,omitempty"`
	Perms2                   *PermType2                 `json:"perms2,omitempty"`
	DisplayName              string                     `json:"display_name,omitempty"`

	GlobalSystemConfigRefs []*QosConfigGlobalSystemConfigRef `json:"global_system_config_refs,omitempty"`
}

// QosConfigGlobalSystemConfigRef references each other
type QosConfigGlobalSystemConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *QosConfig) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeQosConfig makes QosConfig
func MakeQosConfig() *QosConfig {
	return &QosConfig{
		//TODO(nati): Apply default
		MPLSExpEntries:           MakeQosIdForwardingClassPairs(),
		DSCPEntries:              MakeQosIdForwardingClassPairs(),
		UUID:                     "",
		ParentUUID:               "",
		QosConfigType:            MakeQosConfigType(),
		VlanPriorityEntries:      MakeQosIdForwardingClassPairs(),
		DefaultForwardingClassID: MakeForwardingClassId(),
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		ParentType:               "",
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
	}
}

// MakeQosConfigSlice() makes a slice of QosConfig
func MakeQosConfigSlice() []*QosConfig {
	return []*QosConfig{}
}
