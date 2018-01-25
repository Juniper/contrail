package models


import ("fmt"
        "testing")
func TestInterfaceRouteTable(t *testing.T) {
    model := MakeInterfaceRouteTable()
    fmt.Println(model)
}
