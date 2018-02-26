package models

import (
	"fmt"
	"testing"
)

func TestVirtualRouter(t *testing.T) {
	model := MakeVirtualRouter()
	fmt.Println(model)
}
