package models

import (
	"fmt"
	"testing"
)

func TestContrailAnalyticsNode(t *testing.T) {
	model := MakeContrailAnalyticsNode()
	fmt.Println(model)
}
