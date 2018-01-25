package models


import ("fmt"
        "testing")
func TestKeyValuePair(t *testing.T) {
    model := MakeKeyValuePair()
    fmt.Println(model)
}
