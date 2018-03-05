package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeNode makes Node
func MakeNode() *Node {
	return &Node{
		//TODO(nati): Apply default
		UUID:                     "",
		ParentUUID:               "",
		ParentType:               "",
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		Hostname:                 "",
		IPAddress:                "",
		MacAddress:               "",
		Type:                     "",
		Password:                 "",
		Username:                 "",
		AwsAmi:                   "",
		AwsInstanceType:          "",
		GCPImage:                 "",
		GCPMachineType:           "",
		PrivateMachineProperties: "",
		PrivateMachineState:      "",
		IpmiAddress:              "",
		IpmiPassword:             "",
		IpmiUsername:             "",
	}
}

// MakeNode makes Node
func InterfaceToNode(i interface{}) *Node {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Node{
		//TODO(nati): Apply default
		UUID:                     schema.InterfaceToString(m["uuid"]),
		ParentUUID:               schema.InterfaceToString(m["parent_uuid"]),
		ParentType:               schema.InterfaceToString(m["parent_type"]),
		FQName:                   schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                  InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:              schema.InterfaceToString(m["display_name"]),
		Annotations:              InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                   InterfaceToPermType2(m["perms2"]),
		Hostname:                 schema.InterfaceToString(m["hostname"]),
		IPAddress:                schema.InterfaceToString(m["ip_address"]),
		MacAddress:               schema.InterfaceToString(m["mac_address"]),
		Type:                     schema.InterfaceToString(m["type"]),
		Password:                 schema.InterfaceToString(m["password"]),
		Username:                 schema.InterfaceToString(m["username"]),
		AwsAmi:                   schema.InterfaceToString(m["aws_ami"]),
		AwsInstanceType:          schema.InterfaceToString(m["aws_instance_type"]),
		GCPImage:                 schema.InterfaceToString(m["gcp_image"]),
		GCPMachineType:           schema.InterfaceToString(m["gcp_machine_type"]),
		PrivateMachineProperties: schema.InterfaceToString(m["private_machine_properties"]),
		PrivateMachineState:      schema.InterfaceToString(m["private_machine_state"]),
		IpmiAddress:              schema.InterfaceToString(m["ipmi_address"]),
		IpmiPassword:             schema.InterfaceToString(m["ipmi_password"]),
		IpmiUsername:             schema.InterfaceToString(m["ipmi_username"]),
	}
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}

// InterfaceToNodeSlice() makes a slice of Node
func InterfaceToNodeSlice(i interface{}) []*Node {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Node{}
	for _, item := range list {
		result = append(result, InterfaceToNode(item))
	}
	return result
}
