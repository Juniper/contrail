package models

import (
	"fmt"
	"testing"
)

func TestCustomerAttachment(t *testing.T) {
	model := MakeCustomerAttachment()
	fmt.Println(model)
}
