package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyMemdbSchema(t *testing.T) {
	schema := MemDBSchema()
	assert.NoError(t, schema.Validate())
}
