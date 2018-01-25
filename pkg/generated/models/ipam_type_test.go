package models


import ("fmt"
        "testing")
func TestIpamType(t *testing.T) {
    model := MakeIpamType()
    fmt.Println(model)
}
