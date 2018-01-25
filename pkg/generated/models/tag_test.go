package models


import ("fmt"
        "testing")
func TestTag(t *testing.T) {
    model := MakeTag()
    fmt.Println(model)
}
