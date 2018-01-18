package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	ContrailWebui                      string         `json:"contrail_webui,omitempty"`
	DataTTL                            string         `json:"data_ttl,omitempty"`
	DefaultGateway                     string         `json:"default_gateway,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ConfigAuditTTL                     string         `json:"config_audit_ttl,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	StatisticsTTL                      string         `json:"statistics_ttl,omitempty"`
	FlowTTL                            string         `json:"flow_ttl,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members,omitempty"`
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
		ConfigAuditTTL: "",
		ContrailWebui:  "",
		DataTTL:        "",
		DefaultGateway: "",
		UUID:           "",
		DefaultVrouterBondInterface: "",
		Perms2:                             MakePermType2(),
		ParentType:                         "",
		StatisticsTTL:                      "",
		ParentUUID:                         "",
		IDPerms:                            MakeIdPermsType(),
		Annotations:                        MakeKeyValuePairs(),
		DefaultVrouterBondInterfaceMembers: "",
		FlowTTL:     "",
		FQName:      []string{},
		DisplayName: "",
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
