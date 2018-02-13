package models

import (
	"fmt"
	"testing"
)

func TestBridgeDomainMembershipType(t *testing.T) {
	model := MakeBridgeDomainMembershipType()
	fmt.Println(model)
}
