package models


// MakeDomainLimitsType makes DomainLimitsType
func MakeDomainLimitsType() *DomainLimitsType{
    return &DomainLimitsType{
    //TODO(nati): Apply default
    ProjectLimit: 0,
        VirtualNetworkLimit: 0,
        SecurityGroupLimit: 0,
        
    }
}

// MakeDomainLimitsTypeSlice() makes a slice of DomainLimitsType
func MakeDomainLimitsTypeSlice() []*DomainLimitsType {
    return []*DomainLimitsType{}
}


