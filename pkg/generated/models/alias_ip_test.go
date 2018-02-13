package models

import (
	"fmt"
	"testing"
)

func TestAliasIP(t *testing.T) {
	model := MakeAliasIP()
	fmt.Println(model)
}
