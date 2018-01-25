package models


import ("fmt"
        "testing")
func TestAPIAccessList(t *testing.T) {
    model := MakeAPIAccessList()
    fmt.Println(model)
}
