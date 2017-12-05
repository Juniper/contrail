package models

import (
	"fmt"
	"testing"
)

func TestAclEntriesType(t *testing.T) {
	model := MakeAclEntriesType()
	fmt.Println(model)
}
