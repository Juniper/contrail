package models

import (
	"fmt"
	"testing"
)

func TestContrailCluster(t *testing.T) {
	model := MakeContrailCluster()
	fmt.Println(model)
}
