package models


import ("fmt"
        "testing")
func TestSubnet(t *testing.T) {
    model := MakeSubnet()
    fmt.Println(model)
}
