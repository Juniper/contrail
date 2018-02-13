package models

import (
	"fmt"
	"testing"
)

func TestRouteTableType(t *testing.T) {
	model := MakeRouteTableType()
	fmt.Println(model)
}
