package models


import ("fmt"
        "testing")
func TestDhcpOptionType(t *testing.T) {
    model := MakeDhcpOptionType()
    fmt.Println(model)
}
