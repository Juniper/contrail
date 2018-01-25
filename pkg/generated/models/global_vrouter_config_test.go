package models


import ("fmt"
        "testing")
func TestGlobalVrouterConfig(t *testing.T) {
    model := MakeGlobalVrouterConfig()
    fmt.Println(model)
}
