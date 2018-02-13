package models

import (
	"fmt"
	"testing"
)

func TestGlobalQosConfig(t *testing.T) {
	model := MakeGlobalQosConfig()
	fmt.Println(model)
}
