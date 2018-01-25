package models


import ("fmt"
        "testing")
func TestTimerType(t *testing.T) {
    model := MakeTimerType()
    fmt.Println(model)
}
