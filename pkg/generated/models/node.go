package models

// Node

import "encoding/json"

// Node
type Node struct {
	ParentType                     string         `json:"parent_type,omitempty"`
	Hostname                       string         `json:"hostname,omitempty"`
	Type                           string         `json:"type,omitempty"`
	AwsAmi                         string         `json:"aws_ami,omitempty"`
	GCPMachineType                 string         `json:"gcp_machine_type,omitempty"`
	PrivateMachineState            string         `json:"private_machine_state,omitempty"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip,omitempty"`
	Password                       string         `json:"password,omitempty"`
	SSHKey                         string         `json:"ssh_key,omitempty"`
	Username                       string         `json:"username,omitempty"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password,omitempty"`
	FQName                         []string       `json:"fq_name,omitempty"`
	DisplayName                    string         `json:"display_name,omitempty"`
	Perms2                         *PermType2     `json:"perms2,omitempty"`
	MacAddress                     string         `json:"mac_address,omitempty"`
	AwsInstanceType                string         `json:"aws_instance_type,omitempty"`
	GCPImage                       string         `json:"gcp_image,omitempty"`
	IDPerms                        *IdPermsType   `json:"id_perms,omitempty"`
	Annotations                    *KeyValuePairs `json:"annotations,omitempty"`
	UUID                           string         `json:"uuid,omitempty"`
	ParentUUID                     string         `json:"parent_uuid,omitempty"`
	IPAddress                      string         `json:"ip_address,omitempty"`
	PrivateMachineProperties       string         `json:"private_machine_properties,omitempty"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username,omitempty"`
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
		Hostname:                       "",
		Type:                           "",
		AwsAmi:                         "",
		ParentType:                     "",
		Password:                       "",
		SSHKey:                         "",
		Username:                       "",
		GCPMachineType:                 "",
		PrivateMachineState:            "",
		PrivatePowerManagementIP:       "",
		MacAddress:                     "",
		AwsInstanceType:                "",
		GCPImage:                       "",
		PrivatePowerManagementPassword: "",
		FQName:                         []string{},
		DisplayName:                    "",
		Perms2:                         MakePermType2(),
		IPAddress:                      "",
		PrivateMachineProperties:       "",
		PrivatePowerManagementUsername: "",
		IDPerms:     MakeIdPermsType(),
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ParentUUID:  "",
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
