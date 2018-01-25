package models


import ("fmt"
        "testing")
func TestMemberType(t *testing.T) {
    model := MakeMemberType()
    fmt.Println(model)
}
