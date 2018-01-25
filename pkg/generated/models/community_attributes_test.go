package models


import ("fmt"
        "testing")
func TestCommunityAttributes(t *testing.T) {
    model := MakeCommunityAttributes()
    fmt.Println(model)
}
