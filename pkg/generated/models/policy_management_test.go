package models


import ("fmt"
        "testing")
func TestPolicyManagement(t *testing.T) {
    model := MakePolicyManagement()
    fmt.Println(model)
}
