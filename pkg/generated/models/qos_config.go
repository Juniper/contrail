package models

// QosConfig

import "encoding/json"

type QosConfig struct {
	Perms2                   *PermType2                 `json:"perms2"`
	VlanPriorityEntries      *QosIdForwardingClassPairs `json:"vlan_priority_entries"`
	DefaultForwardingClassID ForwardingClassId          `json:"default_forwarding_class_id"`
	DSCPEntries              *QosIdForwardingClassPairs `json:"dscp_entries"`
	IDPerms                  *IdPermsType               `json:"id_perms"`
	Annotations              *KeyValuePairs             `json:"annotations"`
	QosConfigType            QosConfigType              `json:"qos_config_type"`
	MPLSExpEntries           *QosIdForwardingClassPairs `json:"mpls_exp_entries"`
	UUID                     string                     `json:"uuid"`
	FQName                   []string                   `json:"fq_name"`
	DisplayName              string                     `json:"display_name"`

	// global_system_config <utils.Reference Value>
	GlobalSystemConfigRefs []*QosConfigGlobalSystemConfigRef

	GlobalQosConfigs []*QosConfigGlobalQosConfig
	Projects         []*QosConfigProject
}

// <utils.Reference Value>
type QosConfigGlobalSystemConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

type QosConfigGlobalQosConfig struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

type QosConfigProject struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

func (model *QosConfig) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeQosConfig() *QosConfig {
	return &QosConfig{
		//TODO(nati): Apply default
		DisplayName:              "",
		QosConfigType:            MakeQosConfigType(),
		MPLSExpEntries:           MakeQosIdForwardingClassPairs(),
		UUID:                     "",
		FQName:                   []string{},
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		VlanPriorityEntries:      MakeQosIdForwardingClassPairs(),
		DefaultForwardingClassID: MakeForwardingClassId(),
		DSCPEntries:              MakeQosIdForwardingClassPairs(),
		IDPerms:                  MakeIdPermsType(),
	}
}

func InterfaceToQosConfig(iData interface{}) *QosConfig {
	data := iData.(map[string]interface{})
	return &QosConfig{
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string"},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType"},"GoName":"Share","GoType":"[]*ShareType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2"}
		VlanPriorityEntries: InterfaceToQosIdForwardingClassPairs(data["vlan_priority_entries"]),

		//{"Title":"","Description":"Map of .1p priority code to applicable forwarding class for packet.","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"qos_id_forwarding_class_pair":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"forwarding_class_id":{"Title":"","Description":"","SQL":"","Default":"0","Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingClassID","GoType":"ForwardingClassId"},"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPair","CollectionType":"","Column":"","Item":null,"GoName":"QosIDForwardingClassPair","GoType":"QosIdForwardingClassPair"},"GoName":"QosIDForwardingClassPair","GoType":"[]*QosIdForwardingClassPair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPairs","CollectionType":"map","Column":"vlan_priority_entries","Item":null,"GoName":"VlanPriorityEntries","GoType":"QosIdForwardingClassPairs"}
		DefaultForwardingClassID: InterfaceToForwardingClassId(data["default_forwarding_class_id"]),

		//{"Title":"","Description":"Default forwarding class used for all non-specified QOS bits","SQL":"int","Default":"0","Operation":"","Presence":"required","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"default_forwarding_class_id","Item":null,"GoName":"DefaultForwardingClassID","GoType":"ForwardingClassId"}
		DSCPEntries: InterfaceToQosIdForwardingClassPairs(data["dscp_entries"]),

		//{"Title":"","Description":"Map of DSCP match condition and applicable forwarding class for packet.","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"qos_id_forwarding_class_pair":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"forwarding_class_id":{"Title":"","Description":"","SQL":"","Default":"0","Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingClassID","GoType":"ForwardingClassId"},"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPair","CollectionType":"","Column":"","Item":null,"GoName":"QosIDForwardingClassPair","GoType":"QosIdForwardingClassPair"},"GoName":"QosIDForwardingClassPair","GoType":"[]*QosIdForwardingClassPair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPairs","CollectionType":"map","Column":"dscp_entries","Item":null,"GoName":"DSCPEntries","GoType":"QosIdForwardingClassPairs"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string"},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string"},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string"},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool"},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string"},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string"},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType"},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType"},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string"},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair"},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs"}
		QosConfigType: InterfaceToQosConfigType(data["qos_config_type"]),

		//{"Title":"","Description":"Specifies if qos-config is for vhost, fabric or for project.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":["vhost","fabric","project"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosConfigType","CollectionType":"","Column":"qos_config_type","Item":null,"GoName":"QosConfigType","GoType":"QosConfigType"}
		MPLSExpEntries: InterfaceToQosIdForwardingClassPairs(data["mpls_exp_entries"]),

		//{"Title":"","Description":"Map of MPLS EXP values to applicable forwarding class for packet.","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"qos_id_forwarding_class_pair":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"forwarding_class_id":{"Title":"","Description":"","SQL":"","Default":"0","Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingClassID","GoType":"ForwardingClassId"},"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPair","CollectionType":"","Column":"","Item":null,"GoName":"QosIDForwardingClassPair","GoType":"QosIdForwardingClassPair"},"GoName":"QosIDForwardingClassPair","GoType":"[]*QosIdForwardingClassPair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPairs","CollectionType":"map","Column":"mpls_exp_entries","Item":null,"GoName":"MPLSExpEntries","GoType":"QosIdForwardingClassPairs"}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string"}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string"},"GoName":"FQName","GoType":"[]string"}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string"}

	}
}

func InterfaceToQosConfigSlice(data interface{}) []*QosConfig {
	list := data.([]interface{})
	result := MakeQosConfigSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosConfig(item))
	}
	return result
}

func MakeQosConfigSlice() []*QosConfig {
	return []*QosConfig{}
}
