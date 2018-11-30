package neutron

import (
	"github.com/Juniper/contrail/pkg/openstack/models"
)

type Request struct {
	Data    Data                   `json:"data" yaml:"data"`
	Context models.RequestContext `json:"context" yaml:"context"`
}

type Data struct {
	Filters  models.Filters         `json:"filters" yaml:"filters"`
	ID       string          `json:"id" yaml:"id"`
	Fields   models.Fields          `json:"fields" yaml:"fields"`
	Resource models.Resource `json:"resource" yaml:"resource"`
}

type CreateRequest struct {
	ctx      models.RequestContext
	resource models.Resource
}

type UpdateRequest struct {
	ctx      models.RequestContext
	resource models.Resource
	id       string
}

type DeleteRequest struct {
	ctx models.RequestContext
	resource models.Resource
	id  string
}

type ReadRequest struct {
	ctx    models.RequestContext
	resource models.Resource
	fields models.Fields
	id     string
}

type ReadAllRequest struct {
	ctx     models.RequestContext
	resource models.Resource
	filters models.Filters
}

type ReadCountRequest struct {
	ctx     models.RequestContext
	resource models.Resource
	filters models.Filters
}

type AddInterfaceRequest struct {
	ctx      models.RequestContext
	resource models.Resource
}

type DeleteInterfaceRequest struct {
	ctx      models.RequestContext
	resource models.Resource
}

func NewRequest() *Request {
	return &Request{
		Data: Data{
			Filters: models.Filters{},
			Fields:  models.Fields{},
		},
		Context: models.RequestContext{},
	}
}

func (r *Request) GetType() string {
	return r.Context.Type
}
