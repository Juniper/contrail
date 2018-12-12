package logic

import (
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/services"
)

// Resource interface defines Neutron API operations
type Resource interface {
	Create(ctx RequestParameters) (Response, error)
	Update(ctx RequestParameters) (Response, error)
	Delete(ctx RequestParameters, id string) (Response, error)
	Read(ctx RequestParameters, id string) (Response, error)
	ReadAll(ctx RequestParameters, filters Filters, fields Fields) (Response, error)
	ReadCount(ctx RequestParameters, filters Filters) (Response, error)
	AddInterface(ctx RequestParameters) (Response, error)
	DeleteInterface(ctx RequestParameters) (Response, error)
}

// Response interface returned from Neutron API operations
type Response interface{}

type baseResource struct{}

// Create default implementation
func (b *baseResource) Create(ctx RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// Update default implementation
func (b *baseResource) Update(ctx RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// Delete default implementation
func (b *baseResource) Delete(ctx RequestParameters, id string) (Response, error) {
	return nil, errors.New("not implemented")
}

// Read default implementation
func (b *baseResource) Read(ctx RequestParameters, id string) (Response, error) {
	return nil, errors.New("not implemented")
}

// ReadAll default implementation
func (b *baseResource) ReadAll(ctx RequestParameters, f Filters, fields Fields) (Response, error) {
	return nil, errors.New("not implemented")
}

// ReadCount default implementation
func (b *baseResource) ReadCount(ctx RequestParameters, f Filters) (Response, error) {
	return nil, errors.New("not implemented")
}

// AddInterface default implementation
func (b *baseResource) AddInterface(ctx RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// DeleteInterface default implementation
func (b *baseResource) DeleteInterface(ctx RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// Filters used in Neutron read API
type Filters map[string][]string

// Fields used in Neutron read API
type Fields []string

// RequestParameters structure
type RequestParameters struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	WriteService   services.WriteService
	DBService      *db.Service
}
