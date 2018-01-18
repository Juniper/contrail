package models

// Node

import "encoding/json"

// Node
type Node struct {
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
	UUID                           string         `json:"uuid,omitempty"`
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
	Hostname                       string         `json:"hostname,omitempty"`
	IPAddress                      string         `json:"ip_address,omitempty"`
	ParentType                     string         `json:"parent_type,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	Type                           string         `json:"type,omitempty"`
	Username                       string         `json:"username,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	GCPImage                       string         `json:"gcp_image,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	Password                       string         `json:"password,omitempty"`
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
		GCPImage:                       "",
		PrivatePowerManagementPassword: "",
		MacAddress:                     "",
		Password:                       "",
		SSHKey:                         "",
		PrivateMachineState:            "",
		PrivatePowerManagementUsername: "",
		DisplayName:                    "",
		UUID:                           "",
		ParentUUID:                     "",
		Hostname:                       "",
		IPAddress:                      "",
		PrivateMachineProperties:       "",
		ParentType:                     "",
		Annotations:                    MakeKeyValuePairs(),
		AwsAmi:                         "",
		GCPMachineType:                 "",
		IDPerms:                        MakeIdPermsType(),
		PrivatePowerManagementIP:       "",
		Perms2:          MakePermType2(),
		FQName:          []string{},
		Type:            "",
		Username:        "",
		AwsInstanceType: "",
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
