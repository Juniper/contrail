package models

import (
	"fmt"
	"testing"
)

func TestOpenstackCluster(t *testing.T) {
	model := MakeOpenstackCluster()
	fmt.Println(model)
}
