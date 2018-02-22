package models


// MakePortMappings makes PortMappings
func MakePortMappings() *PortMappings{
    return &PortMappings{
    //TODO(nati): Apply default
    
            
                PortMappings:  MakePortMapSlice(),
            
        
    }
}

// MakePortMappingsSlice() makes a slice of PortMappings
func MakePortMappingsSlice() []*PortMappings {
    return []*PortMappings{}
}


