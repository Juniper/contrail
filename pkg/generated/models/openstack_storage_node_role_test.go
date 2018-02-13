package models

import (
	"fmt"
	"testing"
)

func TestOpenstackStorageNodeRole(t *testing.T) {
	model := MakeOpenstackStorageNodeRole()
	fmt.Println(model)
}
