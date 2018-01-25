package models


import ("fmt"
        "testing")
func TestKeypair(t *testing.T) {
    model := MakeKeypair()
    fmt.Println(model)
}
