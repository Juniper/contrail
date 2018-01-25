package models


import ("fmt"
        "testing")
func TestLoadbalancer(t *testing.T) {
    model := MakeLoadbalancer()
    fmt.Println(model)
}
