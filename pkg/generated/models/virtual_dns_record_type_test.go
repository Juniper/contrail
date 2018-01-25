package models


import ("fmt"
        "testing")
func TestVirtualDnsRecordType(t *testing.T) {
    model := MakeVirtualDnsRecordType()
    fmt.Println(model)
}
