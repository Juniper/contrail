package models

import (
	"fmt"
	"testing"
)

func TestE2ServiceProvider(t *testing.T) {
	model := MakeE2ServiceProvider()
	fmt.Println(model)
}
