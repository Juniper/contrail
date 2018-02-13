package models

import (
	"fmt"
	"testing"
)

func TestKubernetesCluster(t *testing.T) {
	model := MakeKubernetesCluster()
	fmt.Println(model)
}
