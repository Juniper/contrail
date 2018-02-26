package models

import (
	"fmt"
	"testing"
)

func TestKubernetesNode(t *testing.T) {
	model := MakeKubernetesNode()
	fmt.Println(model)
}
