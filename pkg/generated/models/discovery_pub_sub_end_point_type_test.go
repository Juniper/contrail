package models

import (
	"fmt"
	"testing"
)

func TestDiscoveryPubSubEndPointType(t *testing.T) {
	model := MakeDiscoveryPubSubEndPointType()
	fmt.Println(model)
}
