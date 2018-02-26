package models

import (
	"fmt"
	"testing"
)

func TestAnalyticsNode(t *testing.T) {
	model := MakeAnalyticsNode()
	fmt.Println(model)
}
