package models


import ("fmt"
        "testing")
func TestMACLimitControlType(t *testing.T) {
    model := MakeMACLimitControlType()
    fmt.Println(model)
}
