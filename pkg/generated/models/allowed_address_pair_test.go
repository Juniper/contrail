package models


import ("fmt"
        "testing")
func TestAllowedAddressPair(t *testing.T) {
    model := MakeAllowedAddressPair()
    fmt.Println(model)
}
