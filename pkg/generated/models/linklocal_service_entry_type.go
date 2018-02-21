package models


// MakeLinklocalServiceEntryType makes LinklocalServiceEntryType
func MakeLinklocalServiceEntryType() *LinklocalServiceEntryType{
    return &LinklocalServiceEntryType{
    //TODO(nati): Apply default
    IPFabricServiceIP: []string{},
        LinklocalServiceName: "",
        LinklocalServiceIP: "",
        IPFabricServicePort: 0,
        IPFabricDNSServiceName: "",
        LinklocalServicePort: 0,
        
    }
}

// MakeLinklocalServiceEntryTypeSlice() makes a slice of LinklocalServiceEntryType
func MakeLinklocalServiceEntryTypeSlice() []*LinklocalServiceEntryType {
    return []*LinklocalServiceEntryType{}
}


