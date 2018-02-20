package models

import (
	"fmt"
	"testing"
)

func TestLocalLinkConnection(t *testing.T) {
	model := MakeLocalLinkConnection()
	fmt.Println(model)
}
