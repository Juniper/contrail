package models

import (
	"fmt"
	"testing"
)

func TestControllerNodeRole(t *testing.T) {
	model := MakeControllerNodeRole()
	fmt.Println(model)
}
