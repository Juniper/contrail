package models


// MakeLinklocalServicesTypes makes LinklocalServicesTypes
func MakeLinklocalServicesTypes() *LinklocalServicesTypes{
    return &LinklocalServicesTypes{
    //TODO(nati): Apply default
    
            
                LinklocalServiceEntry:  MakeLinklocalServiceEntryTypeSlice(),
            
        
    }
}

// MakeLinklocalServicesTypesSlice() makes a slice of LinklocalServicesTypes
func MakeLinklocalServicesTypesSlice() []*LinklocalServicesTypes {
    return []*LinklocalServicesTypes{}
}


