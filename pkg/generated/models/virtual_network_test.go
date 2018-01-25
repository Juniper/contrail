package models


import ("fmt"
        "testing")
func TestVirtualNetwork(t *testing.T) {
    model := MakeVirtualNetwork()
    fmt.Println(model)
}
