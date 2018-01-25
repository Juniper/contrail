package models


import ("fmt"
        "testing")
func TestPhysicalInterface(t *testing.T) {
    model := MakePhysicalInterface()
    fmt.Println(model)
}
