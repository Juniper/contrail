package logic

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/services"
)

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

type Response interface{}

type BaseResource struct{}

func (b *BaseResource) Create(ctx Context) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) Update(ctx Context) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) Delete(ctx Context, id string) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) Read(ctx Context, id string) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) ReadAll(ctx Context, filters Filters, fields Fields) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) ReadCount(ctx Context, filters Filters) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) AddInterface(ctx Context) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BaseResource) DeleteInterface(ctx Context) (Response, error) {
	return nil, fmt.Errorf("not implemented")
}

type Filters map[string][]string
type Fields []string

type Context struct {
	RequestContext RequestContext
	ReadService    services.ReadService
	WriteService   services.WriteService
}
