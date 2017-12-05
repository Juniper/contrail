package models

import (
	"fmt"
	"testing"
)

func TestOpenstackComputeNodeRole(t *testing.T) {
	model := MakeOpenstackComputeNodeRole()
	fmt.Println(model)
}
