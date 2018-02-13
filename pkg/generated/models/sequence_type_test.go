package models

import (
	"fmt"
	"testing"
)

func TestSequenceType(t *testing.T) {
	model := MakeSequenceType()
	fmt.Println(model)
}
