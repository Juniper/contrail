package models


// MakeSubnetType makes SubnetType
func MakeSubnetType() *SubnetType{
    return &SubnetType{
    //TODO(nati): Apply default
    IPPrefix: "",
        IPPrefixLen: 0,
        
    }
}

// MakeSubnetTypeSlice() makes a slice of SubnetType
func MakeSubnetTypeSlice() []*SubnetType {
    return []*SubnetType{}
}


