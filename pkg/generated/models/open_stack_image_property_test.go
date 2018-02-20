package models

import (
	"fmt"
	"testing"
)

func TestOpenStackImageProperty(t *testing.T) {
	model := MakeOpenStackImageProperty()
	fmt.Println(model)
}
