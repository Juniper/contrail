package neutron

import (
	"github.com/Juniper/contrail/pkg/openstack/models"
)

type Request struct {
	Data    Data
	Context models.RequestContext
}

type Data struct {
	Filters  Filters
	ID       string
	Fields   []string
	Resource models.Resource
}

type Filters map[string][]string
type Fields []string

type CreateRequest struct {
	ctx models.RequestContext
	r   models.Resource
}

type UpdateRequest struct {
	ctx models.RequestContext
	r   models.Resource
}

type DeleteRequest struct {
	ctx models.RequestContext
	r   models.Resource
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
	ctx models.RequestContext
}

type AddInterfaceRequest struct {
	ctx models.RequestContext
	r   models.Resource
}

type RemoveInterfaceRequest struct {
	ctx models.RequestContext
	r   models.Resource
}
