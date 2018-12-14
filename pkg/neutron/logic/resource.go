package logic

import (
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/services"
)

// Resource interface defines Neutron API operations
type Resource interface {
	Create(rp RequestParameters) (Response, error)
	Update(rp RequestParameters) (Response, error)
	Delete(rp RequestParameters, id string) (Response, error)
	Read(rp RequestParameters, id string) (Response, error)
	ReadAll(rp RequestParameters, filters Filters, fields Fields) (Response, error)
	ReadCount(rp RequestParameters, filters Filters) (Response, error)
	AddInterface(rp RequestParameters) (Response, error)
	DeleteInterface(rp RequestParameters) (Response, error)
}

// Response interface returned from Neutron API operations
type Response interface{}

type baseResource struct{}

// Create default implementation
func (b *baseResource) Create(rp RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// Update default implementation
func (b *baseResource) Update(rp RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// Delete default implementation
func (b *baseResource) Delete(rp RequestParameters, id string) (Response, error) {
	return nil, errors.New("not implemented")
}

// Read default implementation
func (b *baseResource) Read(rp RequestParameters, id string) (Response, error) {
	return nil, errors.New("not implemented")
}

// ReadAll default implementation
func (b *baseResource) ReadAll(rp RequestParameters, f Filters, fields Fields) (Response, error) {
	return nil, errors.New("not implemented")
}

// ReadCount default implementation
func (b *baseResource) ReadCount(rp RequestParameters, f Filters) (Response, error) {
	return nil, errors.New("not implemented")
}

// AddInterface default implementation
func (b *baseResource) AddInterface(rp RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// DeleteInterface default implementation
func (b *baseResource) DeleteInterface(rp RequestParameters) (Response, error) {
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
}
