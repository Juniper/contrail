package services

import "github.com/Juniper/contrail/pkg/apisrv/baseapisrv"

// TODO Move this into a template in asf, probably into serivices_common.tmpl

// TODO Generate this?
// TODO Add a doc comment
func (service *ContrailService) RegisterGRPCAPI(r baseapisrv.Router) {
	r.RegisterService(&_ContrailService_serviceDesc, service)
	r.RegisterService(&_IPAM_serviceDesc, service)
	r.RegisterService(&_Chown_serviceDesc, service)
	r.RegisterService(&_SetTag_serviceDesc, service)
	r.RegisterService(&_RefRelax_serviceDesc, service)
	r.RegisterService(&_PropCollectionUpdate_serviceDesc, service)

	r.RegisterService(&_FQNameToID_serviceDesc, service)
	r.RegisterService(&_IDToFQName_serviceDesc, service)
	r.RegisterService(&_UserAgentKV_serviceDesc, service)
}
