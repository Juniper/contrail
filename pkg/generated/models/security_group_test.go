package models

import (
	"fmt"
	"testing"
)

func TestSecurityGroup(t *testing.T) {
	model := MakeSecurityGroup()
	fmt.Println(model)
}
