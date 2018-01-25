package models


import ("fmt"
        "testing")
func TestVirtualMachineInterface(t *testing.T) {
    model := MakeVirtualMachineInterface()
    fmt.Println(model)
}
