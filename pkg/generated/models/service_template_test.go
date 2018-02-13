package models

import (
	"fmt"
	"testing"
)

func TestServiceTemplate(t *testing.T) {
	model := MakeServiceTemplate()
	fmt.Println(model)
}
