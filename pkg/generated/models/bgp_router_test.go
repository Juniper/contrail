package models

import (
	"fmt"
	"testing"
)

func TestBGPRouter(t *testing.T) {
	model := MakeBGPRouter()
	fmt.Println(model)
}
