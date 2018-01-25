package models


import ("fmt"
        "testing")
func TestIpamSubnets(t *testing.T) {
    model := MakeIpamSubnets()
    fmt.Println(model)
}
