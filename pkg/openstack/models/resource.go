package models

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/services"
)

type Resource interface {
	Create(CreateContext) (Response, error)
	Update(UpdateContext) (Response, error)
	Delete(DeleteContext) (Response, error)
	Read(ReadContext) (Response, error)
	ReadAll(ReadAllContext) (Response, error)
	ReadCount(ReadCountContext) (Response, error)
	AddInterface(AddInterfaceContext) (Response, error)
	DeleteInterface(DeleteInterfaceContext) (Response, error)
}

type Response interface{}

type BaseResource struct{}

func (b *BaseResource) Create(CreateContext) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) Update(UpdateContext) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) Delete(DeleteContext) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) Read(ReadContext) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) ReadAll(ReadAllContext) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) ReadCount(ReadCountContext) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) AddInterface(AddInterfaceContext) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) DeleteInterface(DeleteInterfaceContext) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

type Filters map[string][]string
type Fields []string

type BaseContext struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	WriteService   services.WriteService
}

type CreateContext struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	WriteService   services.WriteService
}

type UpdateContext struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	WriteService   services.WriteService
}

type DeleteContext struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	WriteService   services.WriteService
	ID             string
}

type ReadContext struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	ID             string
}

type ReadAllContext struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	Fields         Fields
	Filters        Filters
}

type ReadCountContext struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	Filters        Filters
}

type AddInterfaceContext struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	WriteService   services.WriteService
}

type DeleteInterfaceContext struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	WriteService   services.WriteService
}
