package models

import (
	"fmt"
	"testing"
)

func TestKuberunetesNode(t *testing.T) {
	model := MakeKuberunetesNode()
	fmt.Println(model)
}
