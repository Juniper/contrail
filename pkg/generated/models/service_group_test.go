package models


import ("fmt"
        "testing")
func TestServiceGroup(t *testing.T) {
    model := MakeServiceGroup()
    fmt.Println(model)
}
