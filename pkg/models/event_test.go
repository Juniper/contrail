package models

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventEncoding(t *testing.T) {
	e := &Event{
		Operation: Operation_create,
		Resource: &Event_VirtualNetwork{
			VirtualNetwork: &VirtualNetwork{
				UUID: "vn_uuid",
			},
		},
	}
	m, err := json.Marshal(e)
	assert.NoError(t, err, "marhsal event failed")
	fmt.Println(string(m), err)
}
