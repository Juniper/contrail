package models

import (
	"fmt"
	"testing"
)

func TestNamespace(t *testing.T) {
	model := MakeNamespace()
	fmt.Println(model)
}
