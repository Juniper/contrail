package models

import (
	"fmt"
	"testing"
)

func TestProject(t *testing.T) {
	model := MakeProject()
	fmt.Println(model)
}
