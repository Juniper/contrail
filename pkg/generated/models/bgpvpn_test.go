package models


import ("fmt"
        "testing")
func TestBGPVPN(t *testing.T) {
    model := MakeBGPVPN()
    fmt.Println(model)
}
