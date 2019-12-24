package services

import "github.com/Juniper/asf/pkg/apisrv/baseapisrv"

// TODO Move this into a template in asf, probably into serivices_common.tmpl

func registerAPIHomepage(r baseapisrv.HTTPRouter) {
	RegisterSingularPaths(func(path string, name string) {
		r.Register(path, "", name, "resource-base")
	})
	RegisterPluralPaths(func(path string, name string) {
		r.Register(path, "", name, "collection")
	})

	r.Register(FQNameToIDPath, "POST", "name-to-id", "action")
	r.Register(IDToFQNamePath, "POST", "id-to-name", "action")
	r.Register(UserAgentKVPath, "POST", UserAgentKVPath, "action")
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
