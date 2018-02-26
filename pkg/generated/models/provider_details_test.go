package models

import (
	"fmt"
	"testing"
)

func TestProviderDetails(t *testing.T) {
	model := MakeProviderDetails()
	fmt.Println(model)
}
