package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuth(t *testing.T) {
	var nullAuth *AuthContext
	assert.Equal(t, nullAuth.IsAdmin(), true)
	assert.Equal(t, nullAuth.ProjectID(), "admin")
	assert.Equal(t, nullAuth.DomainID(), "admin")

	auth := NewAuthContext("default", "admin", "admin", []string{
		"admin",
	})

	assert.Equal(t, auth.IsAdmin(), true)
	assert.Equal(t, auth.ProjectID(), "admin")
	assert.Equal(t, auth.DomainID(), "default")

	auth = NewAuthContext("default", "demo", "demo", []string{})

	assert.Equal(t, auth.IsAdmin(), false)
	assert.Equal(t, auth.ProjectID(), "demo")
	assert.Equal(t, auth.DomainID(), "default")
}
