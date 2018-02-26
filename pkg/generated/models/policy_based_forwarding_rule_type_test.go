package models

import (
	"fmt"
	"testing"
)

func TestPolicyBasedForwardingRuleType(t *testing.T) {
	model := MakePolicyBasedForwardingRuleType()
	fmt.Println(model)
}
