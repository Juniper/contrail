package models


import ("fmt"
        "testing")
func TestRbacRuleType(t *testing.T) {
    model := MakeRbacRuleType()
    fmt.Println(model)
}
