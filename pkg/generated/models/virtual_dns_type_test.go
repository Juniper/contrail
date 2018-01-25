package models


import ("fmt"
        "testing")
func TestVirtualDnsType(t *testing.T) {
    model := MakeVirtualDnsType()
    fmt.Println(model)
}
