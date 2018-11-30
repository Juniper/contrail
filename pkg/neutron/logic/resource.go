package logic

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/services"
)

// Resource interface
type Resource interface {
	Create(ctx Context) (Response, error)
	Update(ctx Context) (Response, error)
	Delete(ctx Context, id string) (Response, error)
	Read(ctx Context, id string) (Response, error)
	ReadAll(ctx Context, filters Filters, fields Fields) (Response, error)
	ReadCount(ctx Context, filters Filters) (Response, error)
	AddInterface(ctx Context) (Response, error)
	DeleteInterface(ctx Context) (Response, error)
}

// Resource structure
type Response interface{}

// BaseResource structure
type BaseResource struct{}

// Create base logic
func (b *BaseResource) Create(ctx Context) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

// Update base logic
func (b *BaseResource) Update(ctx Context) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

// Delete base logic
func (b *BaseResource) Delete(ctx Context, id string) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

// Read base logic
func (b *BaseResource) Read(ctx Context, id string) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

// ReadAll base logic
func (b *BaseResource) ReadAll(ctx Context, filters Filters, fields Fields) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

// ReadCount base logic
func (b *BaseResource) ReadCount(ctx Context, filters Filters) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

// AddInterface base logic
func (b *BaseResource) AddInterface(ctx Context) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

// DeleteInterface base logic
func (b *BaseResource) DeleteInterface(ctx Context) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

// Filters type
type Filters map[string][]string

// Fields type
type Fields []string

// Context structure
type Context struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	WriteService   services.WriteService
}
