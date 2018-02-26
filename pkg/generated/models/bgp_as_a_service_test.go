package models

import (
	"fmt"
	"testing"
)

func TestBGPAsAService(t *testing.T) {
	model := MakeBGPAsAService()
	fmt.Println(model)
}
