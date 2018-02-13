package models

import (
	"fmt"
	"testing"
)

func TestPortTuple(t *testing.T) {
	model := MakePortTuple()
	fmt.Println(model)
}
