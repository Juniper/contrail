package models

// Node

import "encoding/json"

// Node
type Node struct {
	GCPImage                       string         `json:"gcp_image,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	ParentType                     string         `json:"parent_type,omitempty"`
	Hostname                       string         `json:"hostname,omitempty"`
	IPAddress                      string         `json:"ip_address,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	Password                       string         `json:"password,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	UUID                           string         `json:"uuid,omitempty"`
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
	Username                       string         `json:"username,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
	Type                           string         `json:"type,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
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
		DisplayName:                    "",
		Username:                       "",
		AwsInstanceType:                "",
		PrivateMachineProperties:       "",
		PrivatePowerManagementUsername: "",
		Annotations:                    MakeKeyValuePairs(),
		Type:                           "",
		AwsAmi:                         "",
		PrivateMachineState:            "",
		Password:                       "",
		SSHKey:                         "",
		GCPImage:                       "",
		GCPMachineType:                 "",
		ParentType:                     "",
		Hostname:                       "",
		IPAddress:                      "",
		MacAddress:                     "",
		Perms2:                         MakePermType2(),
		UUID:                           "",
		ParentUUID:                     "",
		PrivatePowerManagementIP:       "",
		FQName:  []string{},
		IDPerms: MakeIdPermsType(),
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
