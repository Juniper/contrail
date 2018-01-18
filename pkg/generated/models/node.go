package models

// Node

import "encoding/json"

// Node
type Node struct {
	GCPImage                       string         `json:"gcp_image,omitempty"`
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	IPAddress                      string         `json:"ip_address,omitempty"`
	Password                       string         `json:"password,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
	ParentType                     string         `json:"parent_type,omitempty"`
	Hostname                       string         `json:"hostname,omitempty"`
	Username                       string         `json:"username,omitempty"`
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	Type                           string         `json:"type,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
	UUID                           string         `json:"uuid,omitempty"`
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
		AwsAmi:                         "",
		GCPMachineType:                 "",
		PrivatePowerManagementPassword: "",
		PrivatePowerManagementUsername: "",
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		IPAddress:                "",
		Password:                 "",
		Perms2:                   MakePermType2(),
		ParentType:               "",
		PrivateMachineProperties: "",
		PrivatePowerManagementIP: "",
		ParentUUID:               "",
		Hostname:                 "",
		Username:                 "",
		SSHKey:                   "",
		AwsInstanceType:          "",
		PrivateMachineState:      "",
		DisplayName:              "",
		UUID:                     "",
		MacAddress:               "",
		Type:                     "",
		GCPImage:                 "",
		Annotations:              MakeKeyValuePairs(),
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
