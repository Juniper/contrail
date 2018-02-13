package models

import (
	"fmt"
	"testing"
)

func TestVirtualDNSRecord(t *testing.T) {
	model := MakeVirtualDNSRecord()
	fmt.Println(model)
}
