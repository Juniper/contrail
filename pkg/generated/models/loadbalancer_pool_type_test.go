package models


import ("fmt"
        "testing")
func TestLoadbalancerPoolType(t *testing.T) {
    model := MakeLoadbalancerPoolType()
    fmt.Println(model)
}
