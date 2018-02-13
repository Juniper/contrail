package models

import (
	"fmt"
	"testing"
)

func TestServiceInterfaceTag(t *testing.T) {
	model := MakeServiceInterfaceTag()
	fmt.Println(model)
}
