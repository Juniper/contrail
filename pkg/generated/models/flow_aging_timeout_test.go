package models


import ("fmt"
        "testing")
func TestFlowAgingTimeout(t *testing.T) {
    model := MakeFlowAgingTimeout()
    fmt.Println(model)
}
