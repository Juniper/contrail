package models

import (
	"fmt"
	"testing"
)

func TestDiscoveryServiceAssignment(t *testing.T) {
	model := MakeDiscoveryServiceAssignment()
	fmt.Println(model)
}
