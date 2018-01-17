package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	FQName                             []string       `json:"fq_name,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ContrailWebui                      string         `json:"contrail_webui,omitempty"`
	FlowTTL                            string         `json:"flow_ttl,omitempty"`
	StatisticsTTL                      string         `json:"statistics_ttl,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ConfigAuditTTL                     string         `json:"config_audit_ttl,omitempty"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	DefaultGateway                     string         `json:"default_gateway,omitempty"`
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	DataTTL                            string         `json:"data_ttl,omitempty"`
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
		Annotations: MakeKeyValuePairs(),
		DataTTL:     "",
		DefaultVrouterBondInterfaceMembers: "",
		ParentType:                         "",
		IDPerms:                            MakeIdPermsType(),
		DisplayName:                        "",
		ContrailWebui:                      "",
		FQName:                             []string{},
		Perms2:                             MakePermType2(),
		ConfigAuditTTL:                     "",
		FlowTTL:                            "",
		StatisticsTTL:                      "",
		UUID:                               "",
		DefaultGateway:                     "",
		DefaultVrouterBondInterface:        "",
		ParentUUID:                         "",
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
