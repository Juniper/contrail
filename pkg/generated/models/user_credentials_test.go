package models

import (
	"fmt"
	"testing"
)

func TestUserCredentials(t *testing.T) {
	model := MakeUserCredentials()
	fmt.Println(model)
}
