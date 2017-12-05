package models

import (
	"fmt"
	"testing"
)

func TestQuotaType(t *testing.T) {
	model := MakeQuotaType()
	fmt.Println(model)
}
