package models

import (
	"fmt"
	"testing"
)

func TestServiceConnectionModule(t *testing.T) {
	model := MakeServiceConnectionModule()
	fmt.Println(model)
}
