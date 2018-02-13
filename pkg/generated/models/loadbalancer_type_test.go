package models

import (
	"fmt"
	"testing"
)

func TestLoadbalancerType(t *testing.T) {
	model := MakeLoadbalancerType()
	fmt.Println(model)
}
