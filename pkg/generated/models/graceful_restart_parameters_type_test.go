package models

import (
	"fmt"
	"testing"
)

func TestGracefulRestartParametersType(t *testing.T) {
	model := MakeGracefulRestartParametersType()
	fmt.Println(model)
}
