package models

import (
	"fmt"
	"testing"
)

func TestOpenStackLink(t *testing.T) {
	model := MakeOpenStackLink()
	fmt.Println(model)
}
