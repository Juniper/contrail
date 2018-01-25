package models


import ("fmt"
        "testing")
func TestUserDefinedLogStat(t *testing.T) {
    model := MakeUserDefinedLogStat()
    fmt.Println(model)
}
