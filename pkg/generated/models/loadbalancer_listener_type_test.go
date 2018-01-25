package models


import ("fmt"
        "testing")
func TestLoadbalancerListenerType(t *testing.T) {
    model := MakeLoadbalancerListenerType()
    fmt.Println(model)
}
