package models

// FirewallRuleDirectionType

type FirewallRuleDirectionType string

// MakeFirewallRuleDirectionType makes FirewallRuleDirectionType
func MakeFirewallRuleDirectionType() FirewallRuleDirectionType {
	var data FirewallRuleDirectionType
	return data
}

// InterfaceToFirewallRuleDirectionType makes FirewallRuleDirectionType from interface
func InterfaceToFirewallRuleDirectionType(data interface{}) FirewallRuleDirectionType {
	return data.(FirewallRuleDirectionType)
}

// InterfaceToFirewallRuleDirectionTypeSlice makes a slice of FirewallRuleDirectionType from interface
func InterfaceToFirewallRuleDirectionTypeSlice(data interface{}) []FirewallRuleDirectionType {
	list := data.([]interface{})
	result := MakeFirewallRuleDirectionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleDirectionType(item))
	}
	return result
}

// MakeFirewallRuleDirectionTypeSlice() makes a slice of FirewallRuleDirectionType
func MakeFirewallRuleDirectionTypeSlice() []FirewallRuleDirectionType {
	return []FirewallRuleDirectionType{}
}
