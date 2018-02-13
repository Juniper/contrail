package models

import (
	"fmt"
	"testing"
)

func TestAllocationPoolType(t *testing.T) {
	model := MakeAllocationPoolType()
	fmt.Println(model)
}
