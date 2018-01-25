package models


import ("fmt"
        "testing")
func TestLogicalInterface(t *testing.T) {
    model := MakeLogicalInterface()
    fmt.Println(model)
}
