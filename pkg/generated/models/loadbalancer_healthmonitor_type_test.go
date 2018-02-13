package models

import (
	"fmt"
	"testing"
)

func TestLoadbalancerHealthmonitorType(t *testing.T) {
	model := MakeLoadbalancerHealthmonitorType()
	fmt.Println(model)
}
