package models


import ("fmt"
        "testing")
func TestServiceEndpoint(t *testing.T) {
    model := MakeServiceEndpoint()
    fmt.Println(model)
}
