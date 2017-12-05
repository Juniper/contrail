package models

import (
	"fmt"
	"testing"
)

func TestForwardingClass(t *testing.T) {
	model := MakeForwardingClass()
	fmt.Println(model)
}
