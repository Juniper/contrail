package models

import (
	"fmt"
	"testing"
)

func TestTelemetryResourceInfo(t *testing.T) {
	model := MakeTelemetryResourceInfo()
	fmt.Println(model)
}
