package models

import (
	"fmt"
	"testing"
)

func TestNetworkIpam(t *testing.T) {
	model := MakeNetworkIpam()
	fmt.Println(model)
}
