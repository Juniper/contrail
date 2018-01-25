package models


import ("fmt"
        "testing")
func TestUser(t *testing.T) {
    model := MakeUser()
    fmt.Println(model)
}
