package models

import (
	"fmt"
	"testing"
)

func TestRouteTargetList(t *testing.T) {
	model := MakeRouteTargetList()
	fmt.Println(model)
}
