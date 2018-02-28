package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeOsImage makes OsImage
// nolint
func MakeOsImage() *OsImage {
	return &OsImage{
		//TODO(nati): Apply default
		UUID:            "",
		ParentUUID:      "",
		ParentType:      "",
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		Name:            "",
		Owner:           "",
		ID:              "",
		Size_:           0,
		Status:          "",
		Location:        "",
		File:            "",
		Checksum:        "",
		CreatedAt:       "",
		UpdatedAt:       "",
		ContainerFormat: "",
		DiskFormat:      "",
		Protected:       false,
		Visibility:      "",
		Property:        "",
		MinDisk:         0,
		MinRAM:          0,
		Tags:            "",
	}
}

// MakeOsImage makes OsImage
// nolint
func InterfaceToOsImage(i interface{}) *OsImage {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OsImage{
		//TODO(nati): Apply default
		UUID:            common.InterfaceToString(m["uuid"]),
		ParentUUID:      common.InterfaceToString(m["parent_uuid"]),
		ParentType:      common.InterfaceToString(m["parent_type"]),
		FQName:          common.InterfaceToStringList(m["fq_name"]),
		IDPerms:         InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:     common.InterfaceToString(m["display_name"]),
		Annotations:     InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:          InterfaceToPermType2(m["perms2"]),
		Name:            common.InterfaceToString(m["name"]),
		Owner:           common.InterfaceToString(m["owner"]),
		ID:              common.InterfaceToString(m["id"]),
		Size_:           common.InterfaceToInt64(m["size"]),
		Status:          common.InterfaceToString(m["status"]),
		Location:        common.InterfaceToString(m["location"]),
		File:            common.InterfaceToString(m["file"]),
		Checksum:        common.InterfaceToString(m["checksum"]),
		CreatedAt:       common.InterfaceToString(m["created_at"]),
		UpdatedAt:       common.InterfaceToString(m["updated_at"]),
		ContainerFormat: common.InterfaceToString(m["container_format"]),
		DiskFormat:      common.InterfaceToString(m["disk_format"]),
		Protected:       common.InterfaceToBool(m["protected"]),
		Visibility:      common.InterfaceToString(m["visibility"]),
		Property:        common.InterfaceToString(m["property"]),
		MinDisk:         common.InterfaceToInt64(m["min_disk"]),
		MinRAM:          common.InterfaceToInt64(m["min_ram"]),
		Tags:            common.InterfaceToString(m["tags"]),
	}
}

// MakeOsImageSlice() makes a slice of OsImage
// nolint
func MakeOsImageSlice() []*OsImage {
	return []*OsImage{}
}

// InterfaceToOsImageSlice() makes a slice of OsImage
// nolint
func InterfaceToOsImageSlice(i interface{}) []*OsImage {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OsImage{}
	for _, item := range list {
		result = append(result, InterfaceToOsImage(item))
	}
	return result
}
