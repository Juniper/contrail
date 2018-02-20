package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeKeypair makes Keypair
func MakeKeypair() *Keypair{
    return &Keypair{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Name: "",
        PublicKey: "",
        
    }
}

// MakeKeypair makes Keypair
func InterfaceToKeypair(i interface{}) *Keypair{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &Keypair{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        Name: schema.InterfaceToString(m["name"]),
        PublicKey: schema.InterfaceToString(m["public_key"]),
        
    }
}

// MakeKeypairSlice() makes a slice of Keypair
func MakeKeypairSlice() []*Keypair {
    return []*Keypair{}
}

// InterfaceToKeypairSlice() makes a slice of Keypair
func InterfaceToKeypairSlice(i interface{}) []*Keypair {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*Keypair{}
    for _, item := range list {
        result = append(result, InterfaceToKeypair(item) )
    }
    return result
}



