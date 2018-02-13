package models

import (
	"fmt"
	"testing"
)

func TestFatFlowProtocols(t *testing.T) {
	model := MakeFatFlowProtocols()
	fmt.Println(model)
}
