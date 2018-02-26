package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeOsImage makes OsImage
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
func InterfaceToOsImage(i interface{}) *OsImage {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OsImage{
		//TODO(nati): Apply default
		UUID:            schema.InterfaceToString(m["uuid"]),
		ParentUUID:      schema.InterfaceToString(m["parent_uuid"]),
		ParentType:      schema.InterfaceToString(m["parent_type"]),
		FQName:          schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:         InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:     schema.InterfaceToString(m["display_name"]),
		Annotations:     InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:          InterfaceToPermType2(m["perms2"]),
		Name:            schema.InterfaceToString(m["name"]),
		Owner:           schema.InterfaceToString(m["owner"]),
		ID:              schema.InterfaceToString(m["id"]),
		Size_:           schema.InterfaceToInt64(m["size"]),
		Status:          schema.InterfaceToString(m["status"]),
		Location:        schema.InterfaceToString(m["location"]),
		File:            schema.InterfaceToString(m["file"]),
		Checksum:        schema.InterfaceToString(m["checksum"]),
		CreatedAt:       schema.InterfaceToString(m["created_at"]),
		UpdatedAt:       schema.InterfaceToString(m["updated_at"]),
		ContainerFormat: schema.InterfaceToString(m["container_format"]),
		DiskFormat:      schema.InterfaceToString(m["disk_format"]),
		Protected:       schema.InterfaceToBool(m["protected"]),
		Visibility:      schema.InterfaceToString(m["visibility"]),
		Property:        schema.InterfaceToString(m["property"]),
		MinDisk:         schema.InterfaceToInt64(m["min_disk"]),
		MinRAM:          schema.InterfaceToInt64(m["min_ram"]),
		Tags:            schema.InterfaceToString(m["tags"]),
	}
}

// MakeOsImageSlice() makes a slice of OsImage
func MakeOsImageSlice() []*OsImage {
	return []*OsImage{}
}

// InterfaceToOsImageSlice() makes a slice of OsImage
func InterfaceToOsImageSlice(i interface{}) []*OsImage {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OsImage{}
	for _, item := range list {
		result = append(result, InterfaceToOsImage(item))
	}
	return result
}
