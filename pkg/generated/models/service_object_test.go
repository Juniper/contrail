package models


import ("fmt"
        "testing")
func TestServiceObject(t *testing.T) {
    model := MakeServiceObject()
    fmt.Println(model)
}
