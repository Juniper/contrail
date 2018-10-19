package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	var nullAuth *Context
	assert.Equal(t, nullAuth.IsAdmin(), true)
	assert.Equal(t, nullAuth.ProjectID(), "admin")
	assert.Equal(t, nullAuth.DomainID(), "admin")

	auth := NewContext("default", "admin", "admin", []string{
		"admin",
	})

	assert.Equal(t, auth.IsAdmin(), true)
	assert.Equal(t, auth.ProjectID(), "admin")
	assert.Equal(t, auth.DomainID(), "default")

	auth = NewContext("default", "demo", "demo", []string{})

	assert.Equal(t, auth.IsAdmin(), false)
	assert.Equal(t, auth.ProjectID(), "demo")
	assert.Equal(t, auth.DomainID(), "default")
}
