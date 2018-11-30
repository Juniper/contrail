package models

import (
	"github.com/Juniper/contrail/pkg/services"
)

func (n *NetworkRequest) Create(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) (Response, error) {
	return &NetworkResponse{
		Name: n.Name,
	}, nil
}
