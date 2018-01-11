package models

// Node

import "encoding/json"

// Node
type Node struct {
	Username                       string         `json:"username"`
	GCPMachineType                 string         `json:"gcp_machine_type"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username"`
	ParentUUID                     string         `json:"parent_uuid"`
	ParentType                     string         `json:"parent_type"`
	IPAddress                      string         `json:"ip_address"`
	Type                           string         `json:"type"`
	FQName                         []string       `json:"fq_name"`
	Perms2                         *PermType2     `json:"perms2"`
	AwsAmi                         string         `json:"aws_ami"`
	PrivateMachineState            string         `json:"private_machine_state"`
	Annotations                    *KeyValuePairs `json:"annotations"`
	Hostname                       string         `json:"hostname"`
	MacAddress                     string         `json:"mac_address"`
	PrivateMachineProperties       string         `json:"private_machine_properties"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip"`
	UUID                           string         `json:"uuid"`
	Password                       string         `json:"password"`
	SSHKey                         string         `json:"ssh_key"`
	IDPerms                        *IdPermsType   `json:"id_perms"`
	DisplayName                    string         `json:"display_name"`
	AwsInstanceType                string         `json:"aws_instance_type"`
	GCPImage                       string         `json:"gcp_image"`
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
		IDPerms:                        MakeIdPermsType(),
		DisplayName:                    "",
		AwsInstanceType:                "",
		Type:                           "",
		Username:                       "",
		GCPMachineType:                 "",
		PrivatePowerManagementPassword: "",
		PrivatePowerManagementUsername: "",
		ParentUUID:                     "",
		ParentType:                     "",
		IPAddress:                      "",
		Perms2:                         MakePermType2(),
		FQName:                         []string{},
		MacAddress:                     "",
		AwsAmi:                         "",
		PrivateMachineState:            "",
		Annotations:                    MakeKeyValuePairs(),
		Hostname:                       "",
		SSHKey:                         "",
		PrivateMachineProperties: "",
		PrivatePowerManagementIP: "",
		UUID:     "",
		Password: "",
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
