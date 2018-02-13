package models

import (
	"fmt"
	"testing"
)

func TestAllowedAddressPairs(t *testing.T) {
	model := MakeAllowedAddressPairs()
	fmt.Println(model)
}
