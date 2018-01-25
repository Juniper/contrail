package models


import ("fmt"
        "testing")
func TestIpamDnsAddressType(t *testing.T) {
    model := MakeIpamDnsAddressType()
    fmt.Println(model)
}
