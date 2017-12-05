package models

import (
	"fmt"
	"testing"
)

func TestSNMPCredentials(t *testing.T) {
	model := MakeSNMPCredentials()
	fmt.Println(model)
}
