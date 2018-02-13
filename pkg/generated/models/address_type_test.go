package models

import (
	"fmt"
	"testing"
)

func TestAddressType(t *testing.T) {
	model := MakeAddressType()
	fmt.Println(model)
}
