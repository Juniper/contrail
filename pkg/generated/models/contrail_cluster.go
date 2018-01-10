package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	DefaultGateway                     string         `json:"default_gateway"`
	ParentType                         string         `json:"parent_type"`
	FQName                             []string       `json:"fq_name"`
	Perms2                             *PermType2     `json:"perms2"`
	UUID                               string         `json:"uuid"`
	IDPerms                            *IdPermsType   `json:"id_perms"`
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members"`
	StatisticsTTL                      string         `json:"statistics_ttl"`
	DisplayName                        string         `json:"display_name"`
	ConfigAuditTTL                     string         `json:"config_audit_ttl"`
	ContrailWebui                      string         `json:"contrail_webui"`
	DataTTL                            string         `json:"data_ttl"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface"`
	FlowTTL                            string         `json:"flow_ttl"`
	ParentUUID                         string         `json:"parent_uuid"`
	Annotations                        *KeyValuePairs `json:"annotations"`
}

// String returns json representation of the object
func (model *ContrailCluster) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeContrailCluster makes ContrailCluster
func MakeContrailCluster() *ContrailCluster {
	return &ContrailCluster{
		//TODO(nati): Apply default
		DefaultGateway:                     "",
		ParentType:                         "",
		FQName:                             []string{},
		Perms2:                             MakePermType2(),
		UUID:                               "",
		IDPerms:                            MakeIdPermsType(),
		DisplayName:                        "",
		ConfigAuditTTL:                     "",
		ContrailWebui:                      "",
		DataTTL:                            "",
		DefaultVrouterBondInterface:        "",
		DefaultVrouterBondInterfaceMembers: "",
		StatisticsTTL:                      "",
		FlowTTL:                            "",
		ParentUUID:                         "",
		Annotations:                        MakeKeyValuePairs(),
	}
}

// InterfaceToContrailCluster makes ContrailCluster from interface
func InterfaceToContrailCluster(iData interface{}) *ContrailCluster {
	data := iData.(map[string]interface{})
	return &ContrailCluster{
		FlowTTL: data["flow_ttl"].(string),

		//{"title":"Flow Data Retention Time","description":"Flow Data Retention Time in hours","default":"2160","type":"string","permission":["create","update"]}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		DefaultGateway: data["default_gateway"].(string),

		//{"title":"Default Gateway","description":"Default Gateway","default":"","type":"string","permission":["create","update"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DefaultVrouterBondInterfaceMembers: data["default_vrouter_bond_interface_members"].(string),

		//{"title":"Default vRouter Bond Interface Members","description":"vRouter Bond Interface Members","default":"ens7f0,ens7f1","type":"string","permission":["create","update"]}
		StatisticsTTL: data["statistics_ttl"].(string),

		//{"title":"Statistics Data Retention Time","description":"Statistics Data Retention Time in hours","default":"2160","type":"string","permission":["create","update"]}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ConfigAuditTTL: data["config_audit_ttl"].(string),

		//{"title":"Configuration Audit Retention Time","description":"Configuration Audit Retention Time in hours","default":"2160","type":"string","permission":["create","update"]}
		ContrailWebui: data["contrail_webui"].(string),

		//{"title":"Contrail WebUI","default":"","type":"string","permission":["create","update"]}
		DataTTL: data["data_ttl"].(string),

		//{"title":"Data Retention Time","description":"Data Retention Time in hours","default":"48","type":"string","permission":["create","update"]}
		DefaultVrouterBondInterface: data["default_vrouter_bond_interface"].(string),

		//{"title":"Default vRouter Bond Interface","description":"vRouter Bond Interface","default":"bond0","type":"string","permission":["create","update"]}

	}
}

// InterfaceToContrailClusterSlice makes a slice of ContrailCluster from interface
func InterfaceToContrailClusterSlice(data interface{}) []*ContrailCluster {
	list := data.([]interface{})
	result := MakeContrailClusterSlice()
	for _, item := range list {
		result = append(result, InterfaceToContrailCluster(item))
	}
	return result
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
