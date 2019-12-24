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

// TODO Add a doc comment
func RegisterRESTAPIHomepage(r baseapisrv.Router) {
	RegisterSingularPaths(func(path string, name string) {
		r.Register(path, "", name, "resource-base")
	})
	RegisterPluralPaths(func(path string, name string) {
		r.Register(path, "", name, "collection")
	})

	r.Register(RefUpdatePath, "POST", RefUpdatePath, "action")
	r.Register(RefRelaxForDeletePath, "POST", RefRelaxForDeletePath, "action")
	r.Register(PropCollectionUpdatePath, "POST", PropCollectionUpdatePath, "action")
	r.Register(SetTagPath, "POST", SetTagPath, "action")
	r.Register(ChownPath, "POST", ChownPath, "action")
	r.Register(IntPoolPath, "GET", IntPoolPath, "action")
	r.Register(IntPoolPath, "POST", IntPoolPath, "action")
	r.Register(IntPoolPath, "DELETE", IntPoolPath, "action")
	r.Register(IntPoolsPath, "POST", IntPoolsPath, "action")
	r.Register(IntPoolsPath, "DELETE", IntPoolsPath, "action")
	r.Register(ObjPerms, "GET", ObjPerms, "action")

	// TODO: register sync?

	// TODO action resources
	// TODO documentation
	// TODO VN IP alloc
	// TODO VN IP free
	// TODO subnet IP count
	// TODO security policy draft
}
