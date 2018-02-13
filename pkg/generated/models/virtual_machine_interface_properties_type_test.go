package models

import (
	"fmt"
	"testing"
)

func TestVirtualMachineInterfacePropertiesType(t *testing.T) {
	model := MakeVirtualMachineInterfacePropertiesType()
	fmt.Println(model)
}
