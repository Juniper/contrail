package models


import ("fmt"
        "testing")
func TestPortMappings(t *testing.T) {
    model := MakePortMappings()
    fmt.Println(model)
}
