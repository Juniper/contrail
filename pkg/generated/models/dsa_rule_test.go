package models

import (
	"fmt"
	"testing"
)

func TestDsaRule(t *testing.T) {
	model := MakeDsaRule()
	fmt.Println(model)
}
