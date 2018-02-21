
package models

import "testing"


import "fmt"
func TestUserDefinedLogStatList(t *testing.T) {
    model := MakeUserDefinedLogStatList()
    fmt.Println(model)
}


func TestUserDefinedLogStatListFQName(t *testing.T) {
    obj := MakeUserDefinedLogStatList()
    if fqname := obj.GetFQName(); len(fqname != 0) {
        t.Errorf("Initial FQName is not empty for UserDefinedLogStatList")
    }
    parent := "fake-parent"
    firstName := "first name"
    fakeFQName := []string{"some", "fake", firstName}
    newname := "fake-me"
    obj.SetFQName(parent, fakeFQName)
    fqname = obj.GetFQName()
    if len(fqname) != len(fakeFQName) {
        t.Errorf("Wrong FQName length, should be %v instead of %v", len(fakeFQName), len(fqname))
    }
    if fqname[len(fqname)-1] != newname {
        t.Errorf("Name should be %v after SetFQName", newname)
    }
}

