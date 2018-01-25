package models


import ("fmt"
        "testing")
func TestFirewallRuleEndpointType(t *testing.T) {
    model := MakeFirewallRuleEndpointType()
    fmt.Println(model)
}
