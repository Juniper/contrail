package models


import ("fmt"
        "testing")
func TestFloatingIP(t *testing.T) {
    model := MakeFloatingIP()
    fmt.Println(model)
}
