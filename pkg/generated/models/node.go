package models

// Node

import "encoding/json"

// Node
type Node struct {
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	IPAddress                      string         `json:"ip_address,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	UUID                           string         `json:"uuid,omitempty"`
	ParentType                     string         `json:"parent_type,omitempty"`
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
	Type                           string         `json:"type,omitempty"`
	Username                       string         `json:"username,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	GCPImage                       string         `json:"gcp_image,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	Hostname                       string         `json:"hostname,omitempty"`
	Password                       string         `json:"password,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
}

// String returns json representation of the object
func (model *Node) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeNode makes Node
func MakeNode() *Node {
	return &Node{
		//TODO(nati): Apply default
		MacAddress:               "",
		GCPMachineType:           "",
		PrivatePowerManagementIP: "",
		UUID:            "",
		ParentType:      "",
		Perms2:          MakePermType2(),
		Type:            "",
		Username:        "",
		AwsInstanceType: "",
		GCPImage:        "",
		IDPerms:         MakeIdPermsType(),
		Annotations:     MakeKeyValuePairs(),
		Hostname:        "",
		Password:        "",
		SSHKey:          "",
		PrivatePowerManagementPassword: "",
		DisplayName:                    "",
		FQName:                         []string{},
		IPAddress:                      "",
		AwsAmi:                         "",
		PrivateMachineProperties:       "",
		PrivateMachineState:            "",
		PrivatePowerManagementUsername: "",
		ParentUUID:                     "",
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
