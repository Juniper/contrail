package models

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	model := MakeNode()
	fmt.Println(model)
}
