package models

import (
	"fmt"
	"testing"
)

func TestQosIdForwardingClassPairs(t *testing.T) {
	model := MakeQosIdForwardingClassPairs()
	fmt.Println(model)
}
