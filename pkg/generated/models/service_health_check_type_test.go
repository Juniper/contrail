package models

import (
	"fmt"
	"testing"
)

func TestServiceHealthCheckType(t *testing.T) {
	model := MakeServiceHealthCheckType()
	fmt.Println(model)
}
