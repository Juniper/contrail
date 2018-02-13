package models

import (
	"fmt"
	"testing"
)

func TestLoadbalancerMember(t *testing.T) {
	model := MakeLoadbalancerMember()
	fmt.Println(model)
}
