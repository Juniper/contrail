package models

// FirewallRuleDirectionType

type FirewallRuleDirectionType string

func MakeFirewallRuleDirectionType() FirewallRuleDirectionType {
	var data FirewallRuleDirectionType
	return data
}

func InterfaceToFirewallRuleDirectionType(data interface{}) FirewallRuleDirectionType {
	return data.(FirewallRuleDirectionType)
}

func InterfaceToFirewallRuleDirectionTypeSlice(data interface{}) []FirewallRuleDirectionType {
	list := data.([]interface{})
	result := MakeFirewallRuleDirectionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleDirectionType(item))
	}
	return result
}

func MakeFirewallRuleDirectionTypeSlice() []FirewallRuleDirectionType {
	return []FirewallRuleDirectionType{}
}
