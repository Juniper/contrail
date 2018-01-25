package models


import ("fmt"
        "testing")
func TestLoadbalancerListener(t *testing.T) {
    model := MakeLoadbalancerListener()
    fmt.Println(model)
}
