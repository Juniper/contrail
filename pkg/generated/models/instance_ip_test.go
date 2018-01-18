package models

import (
	"fmt"
	"testing"
)

func TestInstanceIP(t *testing.T) {
	model := MakeInstanceIP()
	fmt.Println(model)
}
