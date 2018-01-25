package models


import ("fmt"
        "testing")
func TestPeeringPolicy(t *testing.T) {
    model := MakePeeringPolicy()
    fmt.Println(model)
}
