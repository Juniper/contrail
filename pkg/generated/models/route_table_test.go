package models

import (
	"fmt"
	"testing"
)

func TestRouteTable(t *testing.T) {
	model := MakeRouteTable()
	fmt.Println(model)
}
