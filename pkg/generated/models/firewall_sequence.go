package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeFirewallSequence makes FirewallSequence
func MakeFirewallSequence() *FirewallSequence{
    return &FirewallSequence{
    //TODO(nati): Apply default
    Sequence: "",
        
    }
}

// MakeFirewallSequence makes FirewallSequence
func InterfaceToFirewallSequence(i interface{}) *FirewallSequence{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &FirewallSequence{
    //TODO(nati): Apply default
    Sequence: schema.InterfaceToString(m["sequence"]),
        
    }
}

// MakeFirewallSequenceSlice() makes a slice of FirewallSequence
func MakeFirewallSequenceSlice() []*FirewallSequence {
    return []*FirewallSequence{}
}

// InterfaceToFirewallSequenceSlice() makes a slice of FirewallSequence
func InterfaceToFirewallSequenceSlice(i interface{}) []*FirewallSequence {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*FirewallSequence{}
    for _, item := range list {
        result = append(result, InterfaceToFirewallSequence(item) )
    }
    return result
}



