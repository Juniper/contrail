package models


import ("fmt"
        "testing")
func TestIpamSubnetType(t *testing.T) {
    model := MakeIpamSubnetType()
    fmt.Println(model)
}
