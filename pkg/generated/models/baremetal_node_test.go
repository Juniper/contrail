package models


import ("fmt"
        "testing")
func TestBaremetalNode(t *testing.T) {
    model := MakeBaremetalNode()
    fmt.Println(model)
}
