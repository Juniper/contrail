package models

import (
	"fmt"
	"testing"
)

func TestSubnetListType(t *testing.T) {
	model := MakeSubnetListType()
	fmt.Println(model)
}
