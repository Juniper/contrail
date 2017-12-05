package models

import (
	"fmt"
	"testing"
)

func TestSecurityLoggingObject(t *testing.T) {
	model := MakeSecurityLoggingObject()
	fmt.Println(model)
}
