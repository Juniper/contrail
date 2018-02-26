package models

import (
	"fmt"
	"testing"
)

func TestOpenStackFlavorProperty(t *testing.T) {
	model := MakeOpenStackFlavorProperty()
	fmt.Println(model)
}
