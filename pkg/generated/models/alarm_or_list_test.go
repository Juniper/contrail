package models


import ("fmt"
        "testing")
func TestAlarmOrList(t *testing.T) {
    model := MakeAlarmOrList()
    fmt.Println(model)
}
