package models


// MakeSubnetListType makes SubnetListType
func MakeSubnetListType() *SubnetListType{
    return &SubnetListType{
    //TODO(nati): Apply default
    
            
                Subnet:  MakeSubnetTypeSlice(),
            
        
    }
}

// MakeSubnetListTypeSlice() makes a slice of SubnetListType
func MakeSubnetListTypeSlice() []*SubnetListType {
    return []*SubnetListType{}
}


