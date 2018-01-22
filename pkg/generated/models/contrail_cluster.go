package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members,omitempty"`
	StatisticsTTL                      string         `json:"statistics_ttl,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	ContrailWebui                      string         `json:"contrail_webui,omitempty"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	DefaultGateway                     string         `json:"default_gateway,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	DataTTL                            string         `json:"data_ttl,omitempty"`
	FlowTTL                            string         `json:"flow_ttl,omitempty"`
	ConfigAuditTTL                     string         `json:"config_audit_ttl,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
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
		ParentUUID:                         "",
		ParentType:                         "",
		ContrailWebui:                      "",
		DefaultVrouterBondInterface:        "",
		DefaultVrouterBondInterfaceMembers: "",
		StatisticsTTL:                      "",
		UUID:                               "",
		DefaultGateway:                     "",
		DisplayName:                        "",
		Perms2:                             MakePermType2(),
		DataTTL:                            "",
		FlowTTL:                            "",
		FQName:                             []string{},
		Annotations:                        MakeKeyValuePairs(),
		ConfigAuditTTL:                     "",
		IDPerms:                            MakeIdPermsType(),
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
