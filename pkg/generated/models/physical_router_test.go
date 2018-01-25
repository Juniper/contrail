package models


import ("fmt"
        "testing")
func TestPhysicalRouter(t *testing.T) {
    model := MakePhysicalRouter()
    fmt.Println(model)
}
