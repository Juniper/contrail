package models


import ("fmt"
        "testing")
func TestGlobalSystemConfig(t *testing.T) {
    model := MakeGlobalSystemConfig()
    fmt.Println(model)
}
