package serviceif

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	services := []Service{
		&BaseService{},
		&BaseService{},
		&BaseService{},
	}

	Chain(services...)

	assert.Equal(t, services[0].Next(), services[1])
	assert.Equal(t, services[1].Next(), services[2])
	assert.Equal(t, services[2].Next(), nil)
}
