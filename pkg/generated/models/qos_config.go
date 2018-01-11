package models

// QosConfig

import "encoding/json"

// QosConfig
type QosConfig struct {
	DefaultForwardingClassID ForwardingClassId          `json:"default_forwarding_class_id"`
	DSCPEntries              *QosIdForwardingClassPairs `json:"dscp_entries"`
	ParentType               string                     `json:"parent_type"`
	Annotations              *KeyValuePairs             `json:"annotations"`
	Perms2                   *PermType2                 `json:"perms2"`
	UUID                     string                     `json:"uuid"`
	QosConfigType            QosConfigType              `json:"qos_config_type"`
	VlanPriorityEntries      *QosIdForwardingClassPairs `json:"vlan_priority_entries"`
	FQName                   []string                   `json:"fq_name"`
	IDPerms                  *IdPermsType               `json:"id_perms"`
	DisplayName              string                     `json:"display_name"`
	ParentUUID               string                     `json:"parent_uuid"`
	MPLSExpEntries           *QosIdForwardingClassPairs `json:"mpls_exp_entries"`

	GlobalSystemConfigRefs []*QosConfigGlobalSystemConfigRef `json:"global_system_config_refs"`
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
		DefaultForwardingClassID: MakeForwardingClassId(),
		DSCPEntries:              MakeQosIdForwardingClassPairs(),
		ParentType:               "",
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		UUID:                     "",
		MPLSExpEntries:           MakeQosIdForwardingClassPairs(),
		VlanPriorityEntries:      MakeQosIdForwardingClassPairs(),
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		ParentUUID:               "",
	}
}

// MakeQosConfigSlice() makes a slice of QosConfig
func MakeQosConfigSlice() []*QosConfig {
	return []*QosConfig{}
}
