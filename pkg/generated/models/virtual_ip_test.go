package models


import ("fmt"
        "testing")
func TestVirtualIP(t *testing.T) {
    model := MakeVirtualIP()
    fmt.Println(model)
}
