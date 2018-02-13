package models

import (
	"fmt"
	"testing"
)

func TestJunosServicePorts(t *testing.T) {
	model := MakeJunosServicePorts()
	fmt.Println(model)
}
