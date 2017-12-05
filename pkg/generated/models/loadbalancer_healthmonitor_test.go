package models

import (
	"fmt"
	"testing"
)

func TestLoadbalancerHealthmonitor(t *testing.T) {
	model := MakeLoadbalancerHealthmonitor()
	fmt.Println(model)
}
