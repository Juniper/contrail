package models

import (
	"fmt"
	"testing"
)

func TestServiceAppliance(t *testing.T) {
	model := MakeServiceAppliance()
	fmt.Println(model)
}
