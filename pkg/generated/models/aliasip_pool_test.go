package models

import (
	"fmt"
	"testing"
)

func TestAliasIPPool(t *testing.T) {
	model := MakeAliasIPPool()
	fmt.Println(model)
}
