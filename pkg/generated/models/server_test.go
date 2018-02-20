package models

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	model := MakeServer()
	fmt.Println(model)
}
