package models


import ("fmt"
        "testing")
func TestWidget(t *testing.T) {
    model := MakeWidget()
    fmt.Println(model)
}
