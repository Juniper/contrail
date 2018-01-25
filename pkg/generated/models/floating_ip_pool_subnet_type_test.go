package models


import ("fmt"
        "testing")
func TestFloatingIpPoolSubnetType(t *testing.T) {
    model := MakeFloatingIpPoolSubnetType()
    fmt.Println(model)
}
