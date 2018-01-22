package models

// Node

import "encoding/json"

// Node
type Node struct {
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	Type                           string         `json:"type,omitempty"`
	Username                       string         `json:"username,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	Password                       string         `json:"password,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
	ParentType                     string         `json:"parent_type,omitempty"`
	Hostname                       string         `json:"hostname,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	GCPImage                       string         `json:"gcp_image,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	UUID                           string         `json:"uuid,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	IPAddress                      string         `json:"ip_address,omitempty"`
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
		AwsInstanceType:          "",
		GCPImage:                 "",
		GCPMachineType:           "",
		PrivateMachineProperties: "",
		UUID:                     "",
		IDPerms:                  MakeIdPermsType(),
		IPAddress:                "",
		Annotations:              MakeKeyValuePairs(),
		PrivatePowerManagementIP: "",
		Type:     "",
		Username: "",
		AwsAmi:   "",
		PrivatePowerManagementPassword: "",
		ParentUUID:                     "",
		FQName:                         []string{},
		DisplayName:                    "",
		MacAddress:                     "",
		Password:                       "",
		SSHKey:                         "",
		PrivateMachineState:            "",
		PrivatePowerManagementUsername: "",
		Perms2:     MakePermType2(),
		ParentType: "",
		Hostname:   "",
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
