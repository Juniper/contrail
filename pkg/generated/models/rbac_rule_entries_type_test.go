package models

import (
	"fmt"
	"testing"
)

func TestRbacRuleEntriesType(t *testing.T) {
	model := MakeRbacRuleEntriesType()
	fmt.Println(model)
}
