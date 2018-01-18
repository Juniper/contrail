package models

// Node

import "encoding/json"

// Node
type Node struct {
	IPAddress                      string         `json:"ip_address,omitempty"`
	GCPImage                       string         `json:"gcp_image,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
	Password                       string         `json:"password,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	UUID                           string         `json:"uuid,omitempty"`
	Type                           string         `json:"type,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	Hostname                       string         `json:"hostname,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	Username                       string         `json:"username,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	ParentType                     string         `json:"parent_type,omitempty"`
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
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
		Type:                           "",
		AwsAmi:                         "",
		ParentUUID:                     "",
		PrivateMachineProperties:       "",
		PrivateMachineState:            "",
		Hostname:                       "",
		MacAddress:                     "",
		SSHKey:                         "",
		Username:                       "",
		AwsInstanceType:                "",
		GCPMachineType:                 "",
		ParentType:                     "",
		Annotations:                    MakeKeyValuePairs(),
		Perms2:                         MakePermType2(),
		IPAddress:                      "",
		GCPImage:                       "",
		PrivatePowerManagementPassword: "",
		PrivatePowerManagementUsername: "",
		DisplayName:                    "",
		Password:                       "",
		PrivatePowerManagementIP:       "",
		FQName:  []string{},
		IDPerms: MakeIdPermsType(),
		UUID:    "",
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
