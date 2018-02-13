package models

import (
	"fmt"
	"testing"
)

func TestTelemetryStateInfo(t *testing.T) {
	model := MakeTelemetryStateInfo()
	fmt.Println(model)
}
