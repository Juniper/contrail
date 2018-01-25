package models


import ("fmt"
        "testing")
func TestLoadbalancerMemberType(t *testing.T) {
    model := MakeLoadbalancerMemberType()
    fmt.Println(model)
}
