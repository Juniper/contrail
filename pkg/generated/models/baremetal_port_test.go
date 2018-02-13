package models

import (
	"fmt"
	"testing"
)

func TestBaremetalPort(t *testing.T) {
	model := MakeBaremetalPort()
	fmt.Println(model)
}
