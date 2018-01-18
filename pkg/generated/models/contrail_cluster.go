package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	UUID                               string         `json:"uuid,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	StatisticsTTL                      string         `json:"statistics_ttl,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	DataTTL                            string         `json:"data_ttl,omitempty"`
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members,omitempty"`
	DefaultGateway                     string         `json:"default_gateway,omitempty"`
	FlowTTL                            string         `json:"flow_ttl,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	ConfigAuditTTL                     string         `json:"config_audit_ttl,omitempty"`
	ContrailWebui                      string         `json:"contrail_webui,omitempty"`
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
		FlowTTL:        "",
		ParentType:     "",
		DisplayName:    "",
		ConfigAuditTTL: "",
		ContrailWebui:  "",
		DefaultGateway: "",
		ParentUUID:     "",
		FQName:         []string{},
		StatisticsTTL:  "",
		Perms2:         MakePermType2(),
		UUID:           "",
		DefaultVrouterBondInterface: "",
		IDPerms:                     MakeIdPermsType(),
		DataTTL:                     "",
		DefaultVrouterBondInterfaceMembers: "",
		Annotations:                        MakeKeyValuePairs(),
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
