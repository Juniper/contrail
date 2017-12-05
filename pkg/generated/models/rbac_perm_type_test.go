package models

import (
	"fmt"
	"testing"
)

func TestRbacPermType(t *testing.T) {
	model := MakeRbacPermType()
	fmt.Println(model)
}
