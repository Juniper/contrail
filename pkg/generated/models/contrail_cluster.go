package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	ConfigAuditTTL                     string         `json:"config_audit_ttl,omitempty"`
	StatisticsTTL                      string         `json:"statistics_ttl,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ContrailWebui                      string         `json:"contrail_webui,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	DataTTL                            string         `json:"data_ttl,omitempty"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface,omitempty"`
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	DefaultGateway                     string         `json:"default_gateway,omitempty"`
	FlowTTL                            string         `json:"flow_ttl,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
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
		Annotations:                        MakeKeyValuePairs(),
		UUID:                               "",
		ConfigAuditTTL:                     "",
		StatisticsTTL:                      "",
		FQName:                             []string{},
		ContrailWebui:                      "",
		ParentType:                         "",
		DefaultVrouterBondInterfaceMembers: "",
		Perms2:                      MakePermType2(),
		DataTTL:                     "",
		DefaultVrouterBondInterface: "",
		ParentUUID:                  "",
		IDPerms:                     MakeIdPermsType(),
		DisplayName:                 "",
		DefaultGateway:              "",
		FlowTTL:                     "",
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
