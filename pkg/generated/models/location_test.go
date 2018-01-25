package models


import ("fmt"
        "testing")
func TestLocation(t *testing.T) {
    model := MakeLocation()
    fmt.Println(model)
}
