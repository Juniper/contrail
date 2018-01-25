package models


import ("fmt"
        "testing")
func TestAlarm(t *testing.T) {
    model := MakeAlarm()
    fmt.Println(model)
}
