package models

import (
	"fmt"
	"testing"
)

func TestFlowAgingTimeoutList(t *testing.T) {
	model := MakeFlowAgingTimeoutList()
	fmt.Println(model)
}
