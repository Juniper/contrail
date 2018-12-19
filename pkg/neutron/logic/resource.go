package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/services"
)

// Resource interface defines Neutron API operations
type Resource interface {
	Create(ctx context.Context, rp RequestParameters) (Response, error)
	Update(ctx context.Context, rp RequestParameters) (Response, error)
	Delete(ctx context.Context, rp RequestParameters, id string) (Response, error)
	Read(ctx context.Context, rp RequestParameters, id string) (Response, error)
	ReadAll(ctx context.Context, rp RequestParameters, filters Filters, fields Fields) (Response, error)
	ReadCount(ctx context.Context, rp RequestParameters, filters Filters) (Response, error)
	AddInterface(ctx context.Context, rp RequestParameters) (Response, error)
	DeleteInterface(ctx context.Context, rp RequestParameters) (Response, error)
}

// Response interface returned from Neutron API operations
type Response interface{}

type baseResource struct{}

// Create default implementation
func (b *baseResource) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// Update default implementation
func (b *baseResource) Update(ctx context.Context, rp RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// Delete default implementation
func (b *baseResource) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	return nil, errors.New("not implemented")
}

// Read default implementation
func (b *baseResource) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	return nil, errors.New("not implemented")
}

// ReadAll default implementation
func (b *baseResource) ReadAll(
	ctx context.Context, rp RequestParameters, f Filters, fields Fields,
) (Response, error) {
	return nil, errors.New("not implemented")
}

// ReadCount default implementation
func (b *baseResource) ReadCount(ctx context.Context, rp RequestParameters, f Filters) (Response, error) {
	return nil, errors.New("not implemented")
}

// AddInterface default implementation
func (b *baseResource) AddInterface(ctx context.Context, rp RequestParameters) (Response, error) {
	return nil, errors.New("not implemented")
}

// DeleteInterface default implementation
func (b *baseResource) DeleteInterface(ctx context.Context, rp RequestParameters) (Response, error) {
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

// UnmarshalJSON Filters.
func (f *Filters) UnmarshalJSON(data []byte) error {
	if *f == nil {
		*f = Filters{}
	}
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		var ss []string
		switch s := v.(type) {
		case []interface{}:
			for _, i := range s {
				switch c := i.(type) {
				case bool:
					ss = append(ss, fmt.Sprintf("%t", c))
				case string:
					ss = append(ss, fmt.Sprintf("%s", c))
				default:
					return errors.Errorf("%T filter not supported", v)
				}
			}
		default:
			return errors.Errorf("%T filter not supported", v)
		}

		(*f)[k] = ss
	}
	return nil
}
