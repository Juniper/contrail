package models

// Node

import "encoding/json"

// Node
type Node struct {
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
	ParentType                     string         `json:"parent_type,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
	Type                           string         `json:"type,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	IPAddress                      string         `json:"ip_address,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	GCPImage                       string         `json:"gcp_image,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	UUID                           string         `json:"uuid,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	Hostname                       string         `json:"hostname,omitempty"`
	Password                       string         `json:"password,omitempty"`
	Username                       string         `json:"username,omitempty"`
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
		PrivatePowerManagementPassword: "",
		Annotations:                    MakeKeyValuePairs(),
		MacAddress:                     "",
		SSHKey:                         "",
		AwsAmi:                         "",
		PrivatePowerManagementIP: "",
		UUID:                           "",
		FQName:                         []string{},
		IDPerms:                        MakeIdPermsType(),
		Hostname:                       "",
		Password:                       "",
		Username:                       "",
		PrivatePowerManagementUsername: "",
		Perms2:                   MakePermType2(),
		ParentUUID:               "",
		ParentType:               "",
		DisplayName:              "",
		Type:                     "",
		GCPMachineType:           "",
		PrivateMachineState:      "",
		PrivateMachineProperties: "",
		IPAddress:                "",
		AwsInstanceType:          "",
		GCPImage:                 "",
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
