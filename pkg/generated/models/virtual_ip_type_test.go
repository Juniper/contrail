package models


import ("fmt"
        "testing")
func TestVirtualIpType(t *testing.T) {
    model := MakeVirtualIpType()
    fmt.Println(model)
}
