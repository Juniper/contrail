package models


import ("fmt"
        "testing")
func TestConfigRoot(t *testing.T) {
    model := MakeConfigRoot()
    fmt.Println(model)
}
