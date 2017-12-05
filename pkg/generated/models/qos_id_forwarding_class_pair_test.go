package models

import (
	"fmt"
	"testing"
)

func TestQosIdForwardingClassPair(t *testing.T) {
	model := MakeQosIdForwardingClassPair()
	fmt.Println(model)
}
