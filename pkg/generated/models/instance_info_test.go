package models

import (
	"fmt"
	"testing"
)

func TestInstanceInfo(t *testing.T) {
	model := MakeInstanceInfo()
	fmt.Println(model)
}
