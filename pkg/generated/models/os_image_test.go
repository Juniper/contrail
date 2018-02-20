package models

import (
	"fmt"
	"testing"
)

func TestOsImage(t *testing.T) {
	model := MakeOsImage()
	fmt.Println(model)
}
