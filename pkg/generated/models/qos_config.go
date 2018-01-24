package models

// QosConfig

import "encoding/json"

// QosConfig
type QosConfig struct {
	FQName                   []string                   `json:"fq_name,omitempty"`
	DisplayName              string                     `json:"display_name,omitempty"`
	QosConfigType            QosConfigType              `json:"qos_config_type,omitempty"`
	MPLSExpEntries           *QosIdForwardingClassPairs `json:"mpls_exp_entries,omitempty"`
	DefaultForwardingClassID ForwardingClassId          `json:"default_forwarding_class_id,omitempty"`
	DSCPEntries              *QosIdForwardingClassPairs `json:"dscp_entries,omitempty"`
	UUID                     string                     `json:"uuid,omitempty"`
	ParentUUID               string                     `json:"parent_uuid,omitempty"`
	VlanPriorityEntries      *QosIdForwardingClassPairs `json:"vlan_priority_entries,omitempty"`
	Perms2                   *PermType2                 `json:"perms2,omitempty"`
	ParentType               string                     `json:"parent_type,omitempty"`
	IDPerms                  *IdPermsType               `json:"id_perms,omitempty"`
	Annotations              *KeyValuePairs             `json:"annotations,omitempty"`

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
		QosConfigType:            MakeQosConfigType(),
		MPLSExpEntries:           MakeQosIdForwardingClassPairs(),
		DefaultForwardingClassID: MakeForwardingClassId(),
		DSCPEntries:              MakeQosIdForwardingClassPairs(),
		UUID:                     "",
		ParentUUID:               "",
		FQName:                   []string{},
		DisplayName:              "",
		VlanPriorityEntries:      MakeQosIdForwardingClassPairs(),
		Perms2:                   MakePermType2(),
		ParentType:               "",
		IDPerms:                  MakeIdPermsType(),
		Annotations:              MakeKeyValuePairs(),
	}
}

// MakeQosConfigSlice() makes a slice of QosConfig
func MakeQosConfigSlice() []*QosConfig {
	return []*QosConfig{}
}
