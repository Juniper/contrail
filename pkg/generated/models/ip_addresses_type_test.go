package models


import ("fmt"
        "testing")
func TestIpAddressesType(t *testing.T) {
    model := MakeIpAddressesType()
    fmt.Println(model)
}
