package models

import (
	"fmt"
	"testing"
)

func TestDriverInfo(t *testing.T) {
	model := MakeDriverInfo()
	fmt.Println(model)
}
