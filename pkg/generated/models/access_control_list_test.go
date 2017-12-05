package models

import (
	"fmt"
	"testing"
)

func TestAccessControlList(t *testing.T) {
	model := MakeAccessControlList()
	fmt.Println(model)
}
