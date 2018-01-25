package models


import ("fmt"
        "testing")
func TestQosConfig(t *testing.T) {
    model := MakeQosConfig()
    fmt.Println(model)
}
