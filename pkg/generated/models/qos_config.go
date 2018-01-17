package models

// QosConfig

import "encoding/json"

// QosConfig
type QosConfig struct {
	DSCPEntries              *QosIdForwardingClassPairs `json:"dscp_entries,omitempty"`
	Annotations              *KeyValuePairs             `json:"annotations,omitempty"`
	Perms2                   *PermType2                 `json:"perms2,omitempty"`
	IDPerms                  *IdPermsType               `json:"id_perms,omitempty"`
	DefaultForwardingClassID ForwardingClassId          `json:"default_forwarding_class_id,omitempty"`
	MPLSExpEntries           *QosIdForwardingClassPairs `json:"mpls_exp_entries,omitempty"`
	VlanPriorityEntries      *QosIdForwardingClassPairs `json:"vlan_priority_entries,omitempty"`
	DisplayName              string                     `json:"display_name,omitempty"`
	UUID                     string                     `json:"uuid,omitempty"`
	ParentUUID               string                     `json:"parent_uuid,omitempty"`
	ParentType               string                     `json:"parent_type,omitempty"`
	FQName                   []string                   `json:"fq_name,omitempty"`
	QosConfigType            QosConfigType              `json:"qos_config_type,omitempty"`

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
		DefaultForwardingClassID: MakeForwardingClassId(),
		DSCPEntries:              MakeQosIdForwardingClassPairs(),
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		IDPerms:                  MakeIdPermsType(),
		ParentType:               "",
		FQName:                   []string{},
		QosConfigType:            MakeQosConfigType(),
		MPLSExpEntries:           MakeQosIdForwardingClassPairs(),
		VlanPriorityEntries:      MakeQosIdForwardingClassPairs(),
		DisplayName:              "",
		UUID:                     "",
		ParentUUID:               "",
	}
}

// MakeQosConfigSlice() makes a slice of QosConfig
func MakeQosConfigSlice() []*QosConfig {
	return []*QosConfig{}
}
