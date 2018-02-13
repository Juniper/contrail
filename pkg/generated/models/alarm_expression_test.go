package models

import (
	"fmt"
	"testing"
)

func TestAlarmExpression(t *testing.T) {
	model := MakeAlarmExpression()
	fmt.Println(model)
}
