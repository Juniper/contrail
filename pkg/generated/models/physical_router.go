package models


// MakePhysicalRouter makes PhysicalRouter
func MakePhysicalRouter() *PhysicalRouter{
    return &PhysicalRouter{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        PhysicalRouterManagementIP: "",
        PhysicalRouterSNMPCredentials: MakeSNMPCredentials(),
        PhysicalRouterRole: "",
        PhysicalRouterUserCredentials: MakeUserCredentials(),
        PhysicalRouterVendorName: "",
        PhysicalRouterVNCManaged: false,
        PhysicalRouterProductName: "",
        PhysicalRouterLLDP: false,
        PhysicalRouterLoopbackIP: "",
        PhysicalRouterImageURI: "",
        TelemetryInfo: MakeTelemetryStateInfo(),
        PhysicalRouterSNMP: false,
        PhysicalRouterDataplaneIP: "",
        PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
        
    }
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
    return []*PhysicalRouter{}
}


