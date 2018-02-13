package models

import (
	"fmt"
	"testing"
)

func TestVirtualDNS(t *testing.T) {
	model := MakeVirtualDNS()
	fmt.Println(model)
}
