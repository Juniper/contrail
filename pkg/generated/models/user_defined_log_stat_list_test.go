package models

import (
	"fmt"
	"testing"
)

func TestUserDefinedLogStatList(t *testing.T) {
	model := MakeUserDefinedLogStatList()
	fmt.Println(model)
}
