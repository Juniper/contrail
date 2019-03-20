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
	}, "", nil)

	assert.Equal(t, auth.IsAdmin(), true)
	assert.Equal(t, auth.ProjectID(), "admin")
	assert.Equal(t, auth.DomainID(), "default")

	auth = NewContext(
		"default", "demo", "demo", []string{}, "", nil)

	assert.Equal(t, auth.IsAdmin(), false)
	assert.Equal(t, auth.ProjectID(), "demo")
	assert.Equal(t, auth.DomainID(), "default")

	auth = NewContext(
		"default", "demo", "demo", []string{}, "authtoken", nil)

	assert.Equal(t, auth.IsAdmin(), false)
	assert.Equal(t, auth.ProjectID(), "demo")
	assert.Equal(t, auth.DomainID(), "default")
	assert.Equal(t, auth.AuthToken(), "authtoken")
}
