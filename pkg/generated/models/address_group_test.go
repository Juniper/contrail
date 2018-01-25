package models


import ("fmt"
        "testing")
func TestAddressGroup(t *testing.T) {
    model := MakeAddressGroup()
    fmt.Println(model)
}
