package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	UUID                               string         `json:"uuid,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	DataTTL                            string         `json:"data_ttl,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members,omitempty"`
	FlowTTL                            string         `json:"flow_ttl,omitempty"`
	StatisticsTTL                      string         `json:"statistics_ttl,omitempty"`
	ConfigAuditTTL                     string         `json:"config_audit_ttl,omitempty"`
	ContrailWebui                      string         `json:"contrail_webui,omitempty"`
	DefaultGateway                     string         `json:"default_gateway,omitempty"`
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
		DefaultVrouterBondInterface: "",
		FQName:                             []string{},
		Perms2:                             MakePermType2(),
		ParentType:                         "",
		Annotations:                        MakeKeyValuePairs(),
		ConfigAuditTTL:                     "",
		ContrailWebui:                      "",
		DefaultGateway:                     "",
		DefaultVrouterBondInterfaceMembers: "",
		FlowTTL:       "",
		StatisticsTTL: "",
		DataTTL:       "",
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		UUID:          "",
		ParentUUID:    "",
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
