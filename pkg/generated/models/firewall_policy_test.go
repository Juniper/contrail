package models

import (
	"fmt"
	"testing"
)

func TestFirewallPolicy(t *testing.T) {
	model := MakeFirewallPolicy()
	fmt.Println(model)
}
