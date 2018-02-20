package models

import (
	"fmt"
	"testing"
)

func TestBaremetalProperties(t *testing.T) {
	model := MakeBaremetalProperties()
	fmt.Println(model)
}
