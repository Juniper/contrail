package models

import (
	"fmt"
	"testing"
)

func TestAlarmAndList(t *testing.T) {
	model := MakeAlarmAndList()
	fmt.Println(model)
}
