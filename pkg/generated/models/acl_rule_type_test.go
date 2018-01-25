package models


import ("fmt"
        "testing")
func TestAclRuleType(t *testing.T) {
    model := MakeAclRuleType()
    fmt.Println(model)
}
