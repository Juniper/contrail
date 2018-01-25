package models


import ("fmt"
        "testing")
func TestRouteTarget(t *testing.T) {
    model := MakeRouteTarget()
    fmt.Println(model)
}
