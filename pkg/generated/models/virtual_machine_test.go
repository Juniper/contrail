package models


import ("fmt"
        "testing")
func TestVirtualMachine(t *testing.T) {
    model := MakeVirtualMachine()
    fmt.Println(model)
}
