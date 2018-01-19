package models

// Node

import "encoding/json"

// Node
type Node struct {
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
	ParentType                     string         `json:"parent_type,omitempty"`
	Hostname                       string         `json:"hostname,omitempty"`
	IPAddress                      string         `json:"ip_address,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
	Password                       string         `json:"password,omitempty"`
	Username                       string         `json:"username,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
	Type                           string         `json:"type,omitempty"`
	GCPImage                       string         `json:"gcp_image,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
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
		PrivatePowerManagementIP:       "",
		PrivatePowerManagementPassword: "",
		IDPerms:    MakeIdPermsType(),
		ParentUUID: "",
		Password:   "",
		Username:   "",
		AwsAmi:     "",
		PrivateMachineProperties: "",
		FQName:                         []string{},
		MacAddress:                     "",
		GCPMachineType:                 "",
		PrivatePowerManagementUsername: "",
		DisplayName:                    "",
		Type:                           "",
		GCPImage:                       "",
		Perms2:                         MakePermType2(),
		UUID:                           "",
		PrivateMachineState:            "",
		Annotations:                    MakeKeyValuePairs(),
		ParentType:                     "",
		Hostname:                       "",
		IPAddress:                      "",
		SSHKey:                         "",
		AwsInstanceType:                "",
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
