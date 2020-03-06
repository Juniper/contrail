package services

import (
	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/services/baseservices"
)

// Endpoint paths/names
const (
	PropCollectionUpdatePath = "prop-collection-update"
	SetTagPath               = "set-tag"
	ChownPath                = "chown"
)

// ContrailEndpointPlugin provides VNC API-specific endpoints.
type ContrailEndpointPlugin struct {
	Service Service

	InTransactionDoer InTransactionDoer
	MetadataGetter    baseservices.MetadataGetter
	IDToFQNameService IDToFQNameService
}

// RegisterHTTPAPI registers endpoints.
func (p *ContrailEndpointPlugin) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.POST(PropCollectionUpdatePath, p.RESTPropCollectionUpdate)
	r.POST(SetTagPath, p.RESTSetTag)
	r.POST(ChownPath, p.RESTChown)
}

// RegisterGRPCAPI registers GRPC services.
func (p *ContrailEndpointPlugin) RegisterGRPCAPI(r apiserver.GRPCRouter) {
	r.RegisterService(&_PropCollectionUpdate_serviceDesc, p)
	r.RegisterService(&_SetTag_serviceDesc, p)
	r.RegisterService(&_Chown_serviceDesc, p)
}
