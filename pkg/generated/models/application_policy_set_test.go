package models

import (
	"fmt"
	"testing"
)

func TestApplicationPolicySet(t *testing.T) {
	model := MakeApplicationPolicySet()
	fmt.Println(model)
}
