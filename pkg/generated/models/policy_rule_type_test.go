package models

import (
	"fmt"
	"testing"
)

func TestPolicyRuleType(t *testing.T) {
	model := MakePolicyRuleType()
	fmt.Println(model)
}
