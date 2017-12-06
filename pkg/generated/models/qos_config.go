package models

// QosConfig

import "encoding/json"

// QosConfig
type QosConfig struct {
	DisplayName              string                     `json:"display_name"`
	Annotations              *KeyValuePairs             `json:"annotations"`
	Perms2                   *PermType2                 `json:"perms2"`
	MPLSExpEntries           *QosIdForwardingClassPairs `json:"mpls_exp_entries"`
	FQName                   []string                   `json:"fq_name"`
	DefaultForwardingClassID ForwardingClassId          `json:"default_forwarding_class_id"`
	DSCPEntries              *QosIdForwardingClassPairs `json:"dscp_entries"`
	UUID                     string                     `json:"uuid"`
	IDPerms                  *IdPermsType               `json:"id_perms"`
	QosConfigType            QosConfigType              `json:"qos_config_type"`
	VlanPriorityEntries      *QosIdForwardingClassPairs `json:"vlan_priority_entries"`

	// global_system_config <utils.Reference Value>
	GlobalSystemConfigRefs []*QosConfigGlobalSystemConfigRef

	Projects         []*QosConfigProject
	GlobalQosConfigs []*QosConfigGlobalQosConfig
}

// QosConfigGlobalSystemConfigRef references each other
type QosConfigGlobalSystemConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// QosConfig parents relation object

type QosConfigProject struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

type QosConfigGlobalQosConfig struct {
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
		FQName:                   []string{},
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		MPLSExpEntries:           MakeQosIdForwardingClassPairs(),
		VlanPriorityEntries:      MakeQosIdForwardingClassPairs(),
		DefaultForwardingClassID: MakeForwardingClassId(),
		DSCPEntries:              MakeQosIdForwardingClassPairs(),
		UUID:                     "",
		IDPerms:                  MakeIdPermsType(),
		QosConfigType:            MakeQosConfigType(),
	}
}

// InterfaceToQosConfig makes QosConfig from interface
func InterfaceToQosConfig(iData interface{}) *QosConfig {
	data := iData.(map[string]interface{})
	return &QosConfig{
		QosConfigType: InterfaceToQosConfigType(data["qos_config_type"]),

		//{"Title":"","Description":"Specifies if qos-config is for vhost, fabric or for project.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":["vhost","fabric","project"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosConfigType","CollectionType":"","Column":"qos_config_type","Item":null,"GoName":"QosConfigType","GoType":"QosConfigType","GoPremitive":false}
		VlanPriorityEntries: InterfaceToQosIdForwardingClassPairs(data["vlan_priority_entries"]),

		//{"Title":"","Description":"Map of .1p priority code to applicable forwarding class for packet.","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"qos_id_forwarding_class_pair":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"forwarding_class_id":{"Title":"","Description":"","SQL":"","Default":"0","Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingClassID","GoType":"ForwardingClassId","GoPremitive":false},"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPair","CollectionType":"","Column":"","Item":null,"GoName":"QosIDForwardingClassPair","GoType":"QosIdForwardingClassPair","GoPremitive":false},"GoName":"QosIDForwardingClassPair","GoType":"[]*QosIdForwardingClassPair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPairs","CollectionType":"map","Column":"vlan_priority_entries","Item":null,"GoName":"VlanPriorityEntries","GoType":"QosIdForwardingClassPairs","GoPremitive":false}
		DefaultForwardingClassID: InterfaceToForwardingClassId(data["default_forwarding_class_id"]),

		//{"Title":"","Description":"Default forwarding class used for all non-specified QOS bits","SQL":"int","Default":"0","Operation":"","Presence":"required","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"default_forwarding_class_id","Item":null,"GoName":"DefaultForwardingClassID","GoType":"ForwardingClassId","GoPremitive":false}
		DSCPEntries: InterfaceToQosIdForwardingClassPairs(data["dscp_entries"]),

		//{"Title":"","Description":"Map of DSCP match condition and applicable forwarding class for packet.","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"qos_id_forwarding_class_pair":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"forwarding_class_id":{"Title":"","Description":"","SQL":"","Default":"0","Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingClassID","GoType":"ForwardingClassId","GoPremitive":false},"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPair","CollectionType":"","Column":"","Item":null,"GoName":"QosIDForwardingClassPair","GoType":"QosIdForwardingClassPair","GoPremitive":false},"GoName":"QosIDForwardingClassPair","GoType":"[]*QosIdForwardingClassPair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPairs","CollectionType":"map","Column":"dscp_entries","Item":null,"GoName":"DSCPEntries","GoType":"QosIdForwardingClassPairs","GoPremitive":false}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string","GoPremitive":true}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string","GoPremitive":true},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string","GoPremitive":true},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string","GoPremitive":true},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string","GoPremitive":true},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"permissions_owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"permissions_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType","GoPremitive":false},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType","GoPremitive":false}
		MPLSExpEntries: InterfaceToQosIdForwardingClassPairs(data["mpls_exp_entries"]),

		//{"Title":"","Description":"Map of MPLS EXP values to applicable forwarding class for packet.","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"qos_id_forwarding_class_pair":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"forwarding_class_id":{"Title":"","Description":"","SQL":"","Default":"0","Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingClassID","GoType":"ForwardingClassId","GoPremitive":false},"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPair","CollectionType":"","Column":"","Item":null,"GoName":"QosIDForwardingClassPair","GoType":"QosIdForwardingClassPair","GoPremitive":false},"GoName":"QosIDForwardingClassPair","GoType":"[]*QosIdForwardingClassPair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPairs","CollectionType":"map","Column":"mpls_exp_entries","Item":null,"GoName":"MPLSExpEntries","GoType":"QosIdForwardingClassPairs","GoPremitive":false}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string","GoPremitive":true},"GoName":"FQName","GoType":"[]string","GoPremitive":true}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string","GoPremitive":true}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair","GoPremitive":false},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs","GoPremitive":false}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType","GoPremitive":false},"GoName":"Share","GoType":"[]*ShareType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2","GoPremitive":false}

	}
}

// InterfaceToQosConfigSlice makes a slice of QosConfig from interface
func InterfaceToQosConfigSlice(data interface{}) []*QosConfig {
	list := data.([]interface{})
	result := MakeQosConfigSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosConfig(item))
	}
	return result
}

// MakeQosConfigSlice() makes a slice of QosConfig
func MakeQosConfigSlice() []*QosConfig {
	return []*QosConfig{}
}
