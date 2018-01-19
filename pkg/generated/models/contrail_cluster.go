package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	ConfigAuditTTL                     string         `json:"config_audit_ttl,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	DefaultGateway                     string         `json:"default_gateway,omitempty"`
	StatisticsTTL                      string         `json:"statistics_ttl,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	ContrailWebui                      string         `json:"contrail_webui,omitempty"`
	DataTTL                            string         `json:"data_ttl,omitempty"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface,omitempty"`
	FlowTTL                            string         `json:"flow_ttl,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
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
		ContrailWebui:                      "",
		DataTTL:                            "",
		DefaultVrouterBondInterface:        "",
		FlowTTL:                            "",
		ParentType:                         "",
		DefaultVrouterBondInterfaceMembers: "",
		UUID:           "",
		FQName:         []string{},
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		ConfigAuditTTL: "",
		ParentUUID:     "",
		DisplayName:    "",
		DefaultGateway: "",
		StatisticsTTL:  "",
		IDPerms:        MakeIdPermsType(),
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
