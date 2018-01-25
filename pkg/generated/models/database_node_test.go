package models


import ("fmt"
        "testing")
func TestDatabaseNode(t *testing.T) {
    model := MakeDatabaseNode()
    fmt.Println(model)
}
