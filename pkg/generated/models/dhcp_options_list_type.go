package models


// MakeDhcpOptionsListType makes DhcpOptionsListType
func MakeDhcpOptionsListType() *DhcpOptionsListType{
    return &DhcpOptionsListType{
    //TODO(nati): Apply default
    
            
                DHCPOption:  MakeDhcpOptionTypeSlice(),
            
        
    }
}

// MakeDhcpOptionsListTypeSlice() makes a slice of DhcpOptionsListType
func MakeDhcpOptionsListTypeSlice() []*DhcpOptionsListType {
    return []*DhcpOptionsListType{}
}


