package models


import ("fmt"
        "testing")
func TestFirewallRule(t *testing.T) {
    model := MakeFirewallRule()
    fmt.Println(model)
}
