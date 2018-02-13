package models

import (
	"fmt"
	"testing"
)

func TestKeyValuePairs(t *testing.T) {
	model := MakeKeyValuePairs()
	fmt.Println(model)
}
