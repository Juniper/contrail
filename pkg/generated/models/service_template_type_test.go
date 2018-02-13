package models

import (
	"fmt"
	"testing"
)

func TestServiceTemplateType(t *testing.T) {
	model := MakeServiceTemplateType()
	fmt.Println(model)
}
