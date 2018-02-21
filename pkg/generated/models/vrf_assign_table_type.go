package models


// MakeVrfAssignTableType makes VrfAssignTableType
func MakeVrfAssignTableType() *VrfAssignTableType{
    return &VrfAssignTableType{
    //TODO(nati): Apply default
    
            
                VRFAssignRule:  MakeVrfAssignRuleTypeSlice(),
            
        
    }
}

// MakeVrfAssignTableTypeSlice() makes a slice of VrfAssignTableType
func MakeVrfAssignTableTypeSlice() []*VrfAssignTableType {
    return []*VrfAssignTableType{}
}


