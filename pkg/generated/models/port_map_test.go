package models

import (
	"fmt"
	"testing"
)

func TestPortMap(t *testing.T) {
	model := MakePortMap()
	fmt.Println(model)
}
