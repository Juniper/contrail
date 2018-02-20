package models

import (
	"fmt"
	"testing"
)

func TestOpenStackAddress(t *testing.T) {
	model := MakeOpenStackAddress()
	fmt.Println(model)
}
