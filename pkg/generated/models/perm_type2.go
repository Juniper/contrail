package models


// MakePermType2 makes PermType2
func MakePermType2() *PermType2{
    return &PermType2{
    //TODO(nati): Apply default
    Owner: "",
        OwnerAccess: 0,
        GlobalAccess: 0,
        
            
                Share:  MakeShareTypeSlice(),
            
        
    }
}

// MakePermType2Slice() makes a slice of PermType2
func MakePermType2Slice() []*PermType2 {
    return []*PermType2{}
}


