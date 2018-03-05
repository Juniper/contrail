package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeNode makes Node
// nolint
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
		ConfigurationVersion:     0,
		Hostname:                 "",
		IPAddress:                "",
		MacAddress:               "",
		Type:                     "",
		Password:                 "",
		SSHKey:                   "",
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
// nolint
func InterfaceToNode(i interface{}) *Node {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Node{
		//TODO(nati): Apply default
		UUID:                     common.InterfaceToString(m["uuid"]),
		ParentUUID:               common.InterfaceToString(m["parent_uuid"]),
		ParentType:               common.InterfaceToString(m["parent_type"]),
		FQName:                   common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                  InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:              common.InterfaceToString(m["display_name"]),
		Annotations:              InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                   InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:     common.InterfaceToInt64(m["configuration_version"]),
		Hostname:                 common.InterfaceToString(m["hostname"]),
		IPAddress:                common.InterfaceToString(m["ip_address"]),
		MacAddress:               common.InterfaceToString(m["mac_address"]),
		Type:                     common.InterfaceToString(m["type"]),
		Password:                 common.InterfaceToString(m["password"]),
		SSHKey:                   common.InterfaceToString(m["ssh_key"]),
		Username:                 common.InterfaceToString(m["username"]),
		AwsAmi:                   common.InterfaceToString(m["aws_ami"]),
		AwsInstanceType:          common.InterfaceToString(m["aws_instance_type"]),
		GCPImage:                 common.InterfaceToString(m["gcp_image"]),
		GCPMachineType:           common.InterfaceToString(m["gcp_machine_type"]),
		PrivateMachineProperties: common.InterfaceToString(m["private_machine_properties"]),
		PrivateMachineState:      common.InterfaceToString(m["private_machine_state"]),
		IpmiAddress:              common.InterfaceToString(m["ipmi_address"]),
		IpmiPassword:             common.InterfaceToString(m["ipmi_password"]),
		IpmiUsername:             common.InterfaceToString(m["ipmi_username"]),
	}
}

// MakeNodeSlice() makes a slice of Node
// nolint
func MakeNodeSlice() []*Node {
	return []*Node{}
}

// InterfaceToNodeSlice() makes a slice of Node
// nolint
func InterfaceToNodeSlice(i interface{}) []*Node {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Node{}
	for _, item := range list {
		result = append(result, InterfaceToNode(item))
	}
	return result
}
