package models


import ("fmt"
        "testing")
func TestSubnetType(t *testing.T) {
    model := MakeSubnetType()
    fmt.Println(model)
}
