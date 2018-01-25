package models


import ("fmt"
        "testing")
func TestConfigNode(t *testing.T) {
    model := MakeConfigNode()
    fmt.Println(model)
}
