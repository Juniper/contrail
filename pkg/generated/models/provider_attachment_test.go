package models

import (
	"fmt"
	"testing"
)

func TestProviderAttachment(t *testing.T) {
	model := MakeProviderAttachment()
	fmt.Println(model)
}
