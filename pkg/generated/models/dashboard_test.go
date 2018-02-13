package models

import (
	"fmt"
	"testing"
)

func TestDashboard(t *testing.T) {
	model := MakeDashboard()
	fmt.Println(model)
}
