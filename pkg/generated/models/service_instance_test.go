package models


import ("fmt"
        "testing")
func TestServiceInstance(t *testing.T) {
    model := MakeServiceInstance()
    fmt.Println(model)
}
