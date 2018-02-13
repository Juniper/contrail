package models

import (
	"fmt"
	"testing"
)

func TestPolicyEntriesType(t *testing.T) {
	model := MakePolicyEntriesType()
	fmt.Println(model)
}
