package models

// Node

import "encoding/json"

// Node
type Node struct {
	Hostname                       string         `json:"hostname,omitempty"`
	IPAddress                      string         `json:"ip_address,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
	Password                       string         `json:"password,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	UUID                           string         `json:"uuid,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	Type                           string         `json:"type,omitempty"`
	Username                       string         `json:"username,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	GCPImage                       string         `json:"gcp_image,omitempty"`
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	ParentType                     string         `json:"parent_type,omitempty"`
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
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
		IPAddress:                "",
		SSHKey:                   "",
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		Hostname:                 "",
		AwsInstanceType:          "",
		PrivateMachineProperties: "",
		UUID:                           "",
		Password:                       "",
		PrivatePowerManagementIP:       "",
		PrivatePowerManagementUsername: "",
		Perms2:                         MakePermType2(),
		GCPMachineType:                 "",
		Type:                           "",
		Username:                       "",
		AwsAmi:                         "",
		GCPImage:                       "",
		PrivateMachineState:            "",
		PrivatePowerManagementPassword: "",
		ParentType:                     "",
		MacAddress:                     "",
		ParentUUID:                     "",
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
