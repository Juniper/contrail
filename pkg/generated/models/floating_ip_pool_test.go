package models

import (
	"fmt"
	"testing"
)

func TestFloatingIPPool(t *testing.T) {
	model := MakeFloatingIPPool()
	fmt.Println(model)
}
