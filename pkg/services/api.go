package services

import "github.com/Juniper/contrail/pkg/apisrv/baseapisrv"

// TODO Move this into a template in asf, probably into serivices_common.tmpl

// Non-REST API endpoint paths.
const (
	FQNameToIDPath  = "fqname-to-id"
	IDToFQNamePath  = "id-to-fqname"
	UserAgentKVPath = "useragent-kv"
)

// RegisterHTTPAPI registers REST API and action endpoints.
func (service *ContrailService) RegisterHTTPAPI(r baseapisrv.HTTPRouter) error {
	// TODO Rename RegisterRESTAPI to RegisterHTTPAPI
	service.RegisterRESTAPI(r)

	r.POST(FQNameToIDPath, service.RESTFQNameToUUID)
	r.POST(IDToFQNamePath, service.RESTIDToFQName)
	r.POST(UserAgentKVPath, service.RESTUserAgentKV)
	r.POST(UploadCloudKeysPath, service.RESTUploadCloudKeys)

	return nil
}

// RegisterGRPCAPI registers GRPC services.
func (service *ContrailService) RegisterGRPCAPI(r baseapisrv.GRPCRouter) error {
	r.RegisterService(&_ContrailService_serviceDesc, service)
	r.RegisterService(&_IPAM_serviceDesc, service)
	r.RegisterService(&_Chown_serviceDesc, service)
	r.RegisterService(&_SetTag_serviceDesc, service)
	r.RegisterService(&_RefRelax_serviceDesc, service)
	r.RegisterService(&_PropCollectionUpdate_serviceDesc, service)

	r.RegisterService(&_FQNameToID_serviceDesc, service)
	r.RegisterService(&_IDToFQName_serviceDesc, service)
	r.RegisterService(&_UserAgentKV_serviceDesc, service)

	return nil
}
