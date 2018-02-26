package models

import (
	"fmt"
	"testing"
)

func TestServiceHealthCheck(t *testing.T) {
	model := MakeServiceHealthCheck()
	fmt.Println(model)
}
