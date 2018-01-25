package models


import ("fmt"
        "testing")
func TestRoutingPolicy(t *testing.T) {
    model := MakeRoutingPolicy()
    fmt.Println(model)
}
