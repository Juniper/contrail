package models

import (
	"fmt"
	"testing"
)

func TestVPNGroup(t *testing.T) {
	model := MakeVPNGroup()
	fmt.Println(model)
}
