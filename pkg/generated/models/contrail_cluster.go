package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	ConfigAuditTTL                     string         `json:"config_audit_ttl,omitempty"`
	StatisticsTTL                      string         `json:"statistics_ttl,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	DataTTL                            string         `json:"data_ttl,omitempty"`
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	DefaultGateway                     string         `json:"default_gateway,omitempty"`
	FlowTTL                            string         `json:"flow_ttl,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	ContrailWebui                      string         `json:"contrail_webui,omitempty"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface,omitempty"`
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
		DataTTL: "",
		DefaultVrouterBondInterfaceMembers: "",
		FQName:                      []string{},
		DefaultGateway:              "",
		FlowTTL:                     "",
		Annotations:                 MakeKeyValuePairs(),
		UUID:                        "",
		ParentType:                  "",
		IDPerms:                     MakeIdPermsType(),
		ContrailWebui:               "",
		DefaultVrouterBondInterface: "",
		Perms2:         MakePermType2(),
		ConfigAuditTTL: "",
		StatisticsTTL:  "",
		DisplayName:    "",
		ParentUUID:     "",
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
