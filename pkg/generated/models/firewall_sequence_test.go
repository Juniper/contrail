package models


import ("fmt"
        "testing")
func TestFirewallSequence(t *testing.T) {
    model := MakeFirewallSequence()
    fmt.Println(model)
}
