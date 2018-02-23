package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeQosIdForwardingClassPairs makes QosIdForwardingClassPairs
func MakeQosIdForwardingClassPairs() *QosIdForwardingClassPairs{
    return &QosIdForwardingClassPairs{
    //TODO(nati): Apply default
    
            
                QosIDForwardingClassPair:  MakeQosIdForwardingClassPairSlice(),
            
        
    }
}

// MakeQosIdForwardingClassPairs makes QosIdForwardingClassPairs
func InterfaceToQosIdForwardingClassPairs(i interface{}) *QosIdForwardingClassPairs{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &QosIdForwardingClassPairs{
    //TODO(nati): Apply default
    
            
                QosIDForwardingClassPair:  InterfaceToQosIdForwardingClassPairSlice(m["qos_id_forwarding_class_pair"]),
            
        
    }
}

// MakeQosIdForwardingClassPairsSlice() makes a slice of QosIdForwardingClassPairs
func MakeQosIdForwardingClassPairsSlice() []*QosIdForwardingClassPairs {
    return []*QosIdForwardingClassPairs{}
}

// InterfaceToQosIdForwardingClassPairsSlice() makes a slice of QosIdForwardingClassPairs
func InterfaceToQosIdForwardingClassPairsSlice(i interface{}) []*QosIdForwardingClassPairs {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*QosIdForwardingClassPairs{}
    for _, item := range list {
        result = append(result, InterfaceToQosIdForwardingClassPairs(item) )
    }
    return result
}



