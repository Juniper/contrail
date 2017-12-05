package models

import (
	"fmt"
	"testing"
)

func TestNetworkPolicy(t *testing.T) {
	model := MakeNetworkPolicy()
	fmt.Println(model)
}
