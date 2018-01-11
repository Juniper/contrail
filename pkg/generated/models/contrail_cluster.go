package models

// ContrailCluster

import "encoding/json"

// ContrailCluster
type ContrailCluster struct {
	ContrailWebui                      string         `json:"contrail_webui"`
	StatisticsTTL                      string         `json:"statistics_ttl"`
	Perms2                             *PermType2     `json:"perms2"`
	FQName                             []string       `json:"fq_name"`
	DisplayName                        string         `json:"display_name"`
	DataTTL                            string         `json:"data_ttl"`
	DefaultVrouterBondInterfaceMembers string         `json:"default_vrouter_bond_interface_members"`
	UUID                               string         `json:"uuid"`
	IDPerms                            *IdPermsType   `json:"id_perms"`
	ConfigAuditTTL                     string         `json:"config_audit_ttl"`
	ParentUUID                         string         `json:"parent_uuid"`
	Annotations                        *KeyValuePairs `json:"annotations"`
	DefaultGateway                     string         `json:"default_gateway"`
	DefaultVrouterBondInterface        string         `json:"default_vrouter_bond_interface"`
	FlowTTL                            string         `json:"flow_ttl"`
	ParentType                         string         `json:"parent_type"`
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
		ConfigAuditTTL:              "",
		ParentUUID:                  "",
		Annotations:                 MakeKeyValuePairs(),
		ParentType:                  "",
		DefaultGateway:              "",
		DefaultVrouterBondInterface: "",
		FlowTTL:                     "",
		FQName:                      []string{},
		DisplayName:                 "",
		ContrailWebui:               "",
		StatisticsTTL:               "",
		Perms2:                      MakePermType2(),
		IDPerms:                     MakeIdPermsType(),
		DataTTL:                     "",
		DefaultVrouterBondInterfaceMembers: "",
		UUID: "",
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}
