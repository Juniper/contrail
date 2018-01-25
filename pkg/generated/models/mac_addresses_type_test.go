package models


import ("fmt"
        "testing")
func TestMacAddressesType(t *testing.T) {
    model := MakeMacAddressesType()
    fmt.Println(model)
}
