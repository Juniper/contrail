package models

import (
	"fmt"
	"testing"
)

func TestLogicalRouter(t *testing.T) {
	model := MakeLogicalRouter()
	fmt.Println(model)
}
