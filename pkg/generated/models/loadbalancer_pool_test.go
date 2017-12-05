package models

import (
	"fmt"
	"testing"
)

func TestLoadbalancerPool(t *testing.T) {
	model := MakeLoadbalancerPool()
	fmt.Println(model)
}
