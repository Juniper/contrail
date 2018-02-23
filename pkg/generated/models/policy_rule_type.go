package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakePolicyRuleType makes PolicyRuleType
func MakePolicyRuleType() *PolicyRuleType{
    return &PolicyRuleType{
    //TODO(nati): Apply default
    Direction: "",
        Protocol: "",
        
            
                DSTAddresses:  MakeAddressTypeSlice(),
            
        ActionList: MakeActionListType(),
        Created: "",
        RuleUUID: "",
        
            
                DSTPorts:  MakePortTypeSlice(),
            
        Application: []string{},
        LastModified: "",
        Ethertype: "",
        
            
                SRCAddresses:  MakeAddressTypeSlice(),
            
        RuleSequence: MakeSequenceType(),
        
            
                SRCPorts:  MakePortTypeSlice(),
            
        
    }
}

// MakePolicyRuleType makes PolicyRuleType
func InterfaceToPolicyRuleType(i interface{}) *PolicyRuleType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &PolicyRuleType{
    //TODO(nati): Apply default
    Direction: schema.InterfaceToString(m["direction"]),
        Protocol: schema.InterfaceToString(m["protocol"]),
        
            
                DSTAddresses:  InterfaceToAddressTypeSlice(m["dst_addresses"]),
            
        ActionList: InterfaceToActionListType(m["action_list"]),
        Created: schema.InterfaceToString(m["created"]),
        RuleUUID: schema.InterfaceToString(m["rule_uuid"]),
        
            
                DSTPorts:  InterfaceToPortTypeSlice(m["dst_ports"]),
            
        Application: schema.InterfaceToStringList(m["application"]),
        LastModified: schema.InterfaceToString(m["last_modified"]),
        Ethertype: schema.InterfaceToString(m["ethertype"]),
        
            
                SRCAddresses:  InterfaceToAddressTypeSlice(m["src_addresses"]),
            
        RuleSequence: InterfaceToSequenceType(m["rule_sequence"]),
        
            
                SRCPorts:  InterfaceToPortTypeSlice(m["src_ports"]),
            
        
    }
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
    return []*PolicyRuleType{}
}

// InterfaceToPolicyRuleTypeSlice() makes a slice of PolicyRuleType
func InterfaceToPolicyRuleTypeSlice(i interface{}) []*PolicyRuleType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*PolicyRuleType{}
    for _, item := range list {
        result = append(result, InterfaceToPolicyRuleType(item) )
    }
    return result
}



