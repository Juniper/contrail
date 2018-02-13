package models

import (
	"fmt"
	"testing"
)

func TestDomain(t *testing.T) {
	model := MakeDomain()
	fmt.Println(model)
}
