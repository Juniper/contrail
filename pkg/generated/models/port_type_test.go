package models

import (
	"fmt"
	"testing"
)

func TestPortType(t *testing.T) {
	model := MakePortType()
	fmt.Println(model)
}
