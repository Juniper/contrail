package models


import ("fmt"
        "testing")
func TestDhcpOptionsListType(t *testing.T) {
    model := MakeDhcpOptionsListType()
    fmt.Println(model)
}
