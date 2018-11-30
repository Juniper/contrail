package neutron

import (
	"github.com/Juniper/contrail/pkg/openstack/models"
)

type Request struct {
	Data    Data                  `json:"data" yaml:"data"`
	Context models.RequestContext `json:"context" yaml:"context"`
}

type Data struct {
	Filters  Filters         `json:"filters" yaml:"filters"`
	ID       string          `json:"id" yaml:"id"`
	Fields   Fields          `json:"fields" yaml:"fields"`
	Resource models.Resource `json:"resource" yaml:"resource"`
}

type Filters map[string][]string
type Fields []string

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
	id  string
}

type ReadRequest struct {
	ctx    models.RequestContext
	fields Fields
	id     string
}

type ReadAllRequest struct {
	ctx     models.RequestContext
	filters Filters
}

type ReadCountRequest struct {
	ctx     models.RequestContext
	filters Filters
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
			Filters: Filters{},
			Fields:  Fields{},
		},
		Context: models.RequestContext{},
	}
}

func (r *Request) GetType() string {
	return r.Context.Type
}
