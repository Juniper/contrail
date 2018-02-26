package models

import (
	"fmt"
	"testing"
)

func TestPluginProperties(t *testing.T) {
	model := MakePluginProperties()
	fmt.Println(model)
}
