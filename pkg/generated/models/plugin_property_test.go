package models

import (
	"fmt"
	"testing"
)

func TestPluginProperty(t *testing.T) {
	model := MakePluginProperty()
	fmt.Println(model)
}
