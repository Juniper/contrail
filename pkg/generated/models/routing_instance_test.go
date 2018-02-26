package models

import (
	"fmt"
	"testing"
)

func TestRoutingInstance(t *testing.T) {
	model := MakeRoutingInstance()
	fmt.Println(model)
}
