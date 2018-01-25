package models


import ("fmt"
        "testing")
func TestQosQueue(t *testing.T) {
    model := MakeQosQueue()
    fmt.Println(model)
}
