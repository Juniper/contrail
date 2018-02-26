package models

import (
	"fmt"
	"testing"
)

func TestFlavor(t *testing.T) {
	model := MakeFlavor()
	fmt.Println(model)
}
